package publicapi

import (
	"net/http"
	"github.com/orbs-network/go-experiment/types/services/publicapi"
	"github.com/orbs-network/go-experiment/types/protocol"
	"fmt"
	"strconv"
	"net/url"
)

func (s *server) transferHandler(w http.ResponseWriter, r *http.Request) {
	from , err := parseStringParameter(r.URL, "from")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	to, err := parseStringParameter(r.URL, "to")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	amount, err := parseInt32Parameter(r.URL, "amount")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	t := protocol.Transaction{
		From: &protocol.Address{Username: from},
		To: &protocol.Address{Username: to},
		Amount: amount,
	}
	balance, err := s.service.Transfer(&publicapi.TransferInput{Transaction: &t})
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	w.Write([]byte(fmt.Sprintf("%v", balance.Result)))
}

func (s *server) balanceHandler(w http.ResponseWriter, r *http.Request) {
	from , err := parseStringParameter(r.URL, "from")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	balance, err := s.service.GetBalance(&publicapi.GetBalanceInput{From: &protocol.Address{Username: from}})
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	w.Write([]byte(fmt.Sprintf("%v", balance.Result)))
}

func parseStringParameter(url *url.URL, param string) (string, error) {
	keys, ok := url.Query()[param]
	if !ok || len(keys) < 1 {
		return "", fmt.Errorf("missing parameter '%s'", param)
	}
	value := keys[0]
	if value == "" {
		return "", fmt.Errorf("parameter '%s' is empty", param)
	}
	return value, nil
}

func parseInt32Parameter(url *url.URL, param string) (int32, error) {
	keys, ok := url.Query()[param]
	if !ok || len(keys) < 1 {
		return 0, fmt.Errorf("missing parameter '%s'", param)
	}
	value, err := strconv.ParseInt(keys[0], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parameter '%s' is an invalid number", param)
	}
	return int32(value), nil
}