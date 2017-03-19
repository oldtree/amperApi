package amperApi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var AmperClient *http.Client

func AmperGetRequest(targetUrl string) (respData []byte, err error) {
	req, err := http.NewRequest("GET", targetUrl, nil)
	req.Header.Add("Authorization", AuthCode)
	if err != nil {
		return nil, err
	}
	resp, err := AmperClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request error : %d ", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AmperPostRequest(targetUrl string, postdata []byte) (respData []byte, err error) {
	postdataReader := bytes.NewReader(postdata)
	req, err := http.NewRequest("POST", targetUrl, postdataReader)
	req.Header.Add("Authorization", AuthCode)
	if err != nil {
		return nil, err
	}
	resp, err := AmperClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request error : %d ", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
