package parse

import (
	"io/ioutil"
	"net/http"
)

func SendRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetRespBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func GetJson(url string) []byte {
	resp, err := SendRequest(url)
	if err != nil {
		panic(err)
	}
	body, err := GetRespBody(resp)
	if err != nil {
		panic(err)
	}
	return body
}
