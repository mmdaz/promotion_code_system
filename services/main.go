package main

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"sync"
)

func main () {
	f := 98910000000
	var wg sync.WaitGroup

	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func(i int) {
			err := sendReq(f + i)
			if err != nil {
				log.Error(err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}


func sendReq(phoneNumber int) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:8000/applyCode"), nil)
	if err != nil {
		return err
	}
	req.Header.Set("PhoneNumber", fmt.Sprint(phoneNumber))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(string(data))
	}
	return nil
}