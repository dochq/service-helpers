package network

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func SendRequest(method, link, data string, headers map[string]string) (resp *http.Response, body string, err error) {
	req, err := http.NewRequest(method, link, strings.NewReader(data))
	if err != nil {
		return resp, body, err
	}
	req.Close = true
	for key := range headers {
		req.Header.Set(key, headers[key])
	}
	http.DefaultClient = &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: time.Minute,
		},
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return resp, body, err
	}
	resp.Close = true
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, body, err
	}
	resp.Body.Close()
	body = string(respBody)
	return resp, body, err
}
