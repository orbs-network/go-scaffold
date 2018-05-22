package stress

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/orbs-network/go-scaffold/utils/testing"
	"github.com/orbs-network/go-scaffold/services"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"net/http"
	"fmt"
	"time"
	"math/rand"
	"io/ioutil"
	"log"
)

const NUM_TRANSACTIONS = 20000
const NUM_USERS = 100
const TRANSACTIONS_PER_BATCH = 200
const BATCHES_PER_SEC = 40

var _ = Describe("Transactions", func() {

	var (
		node services.Node
		stop chan error
	)

	BeforeEach(func() {
		stop = make(chan error, 10)
		node = services.NewNode(&logger.StubLogger{})
		node.Start(&stop)
	})

	AfterEach(func() {
		node.Stop()
		println("\x1b[95;1m")
		println("* If the test keeps failing on too many open files, see:")
		println("  https://stackoverflow.com/questions/7578594/how-to-increase-limits-on-sockets-on-osx-for-load-testing")
		println("* If the test fails on connection reset by peer, try again, it's flaky")
		println("\x1b[0m")
	})

	It("should handle lots and lots of transactions", func() {
		// customize HTTP client to handle many connections
		transport := http.Transport{
			IdleConnTimeout: time.Second*20,
			MaxIdleConns: TRANSACTIONS_PER_BATCH*10,
			MaxIdleConnsPerHost: TRANSACTIONS_PER_BATCH*10,
		}
		client := &http.Client{Transport: &transport}
		// create a local ledger for verification
		ledger := map[string]int32{}
		for i := 0; i < NUM_USERS; i++ {
			ledger[fmt.Sprintf("user%d", i+1)] = 0
		}
		// send all transactions over HTTP in batches
		rand.Seed(42)
		done := make(chan error, TRANSACTIONS_PER_BATCH)
		for i := 0; i < NUM_TRANSACTIONS / TRANSACTIONS_PER_BATCH; i++ {
			log.Printf("Sending %d transactions... (batch %d out of %d)", TRANSACTIONS_PER_BATCH, i+1, NUM_TRANSACTIONS / TRANSACTIONS_PER_BATCH)
			time.Sleep(time.Second / BATCHES_PER_SEC)
			for j := 0; j < TRANSACTIONS_PER_BATCH; j++ {
				from := randomizeUser()
				to := randomizeUser()
				amount := randomizeAmount()
				ledger[from] -= amount
				ledger[to] += amount
				go sendTransaction(client, from, to, amount, &done)
			}
			for j := 0; j < TRANSACTIONS_PER_BATCH; j++ {
				err := <- done
				Expect(err).ToNot(HaveOccurred())
			}
		}
		// verify the ledger
		for i := 0; i < NUM_USERS; i++ {
			user := fmt.Sprintf("user%d", i+1)
			resp, err := client.Get(fmt.Sprintf("http://localhost:8080/api/balance?from=%s", user))
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(ResponseBodyAsString(resp)).To(Equal(fmt.Sprintf("%d", ledger[user])))
		}
	})

})

func randomizeUser() string {
	return fmt.Sprintf("user%d", rand.Intn(NUM_USERS)+1)
}

func randomizeAmount() int32 {
	return rand.Int31n(1000)+1
}

func sendTransaction(client *http.Client, from string, to string, amount int32, done *chan error) {
	url := fmt.Sprintf("http://localhost:8080/api/transfer?from=%s&to=%s&amount=%d", from, to, amount)
	resp, err := client.Get(url)
	if err == nil {
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
	*done <- err
}