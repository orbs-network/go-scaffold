package testing

import (
	"net/http"
	"github.com/onsi/gomega/gbytes"
	"time"
	"io/ioutil"
)

func ResponseBodyAsString(resp *http.Response) string {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(gbytes.TimeoutReader(resp.Body, 10 * time.Second))
	if err == gbytes.ErrTimeout {
		return "<timeout>"
	}
	if err != nil {
		return "<error>"
	}
	return string(bodyBytes)
}