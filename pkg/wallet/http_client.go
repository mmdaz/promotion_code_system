package wallet

import (
	"errors"
	"fmt"
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpClient struct {
	config *pkg.Config
}

func NewHttpClient(config *pkg.Config) *HttpClient {
	return &HttpClient{config: config}
}

func (w *HttpClient) IncreaseAmount(phoneNumber int, amount int) error {
	client := &http.Client{}
	reqBody := strings.NewReader(fmt.Sprintf(`{ "amount": %v }`, amount))
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/increaseCash", w.config.EndPoints.Wallet), reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("PhoneNumber", fmt.Sprint(phoneNumber))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("increase in wallet failed")
	}
	return nil
}
