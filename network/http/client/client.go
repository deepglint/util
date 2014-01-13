package client

import (
	"github.com/liutong19890905/util/security"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (result string, err error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	response, err := client.Do(reqest)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	result = string(body)
	return
}

func AuthGet(url string, username string, password string) (result string, err error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	auth_str := username + ":" + password
	reqest.Header.Add("Authorization", "Basic "+security.EncodeBase64(auth_str))
	response, err := client.Do(reqest)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	result = string(body)
	return
}

func Post(url string, data string) (result string, err error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return
	}
	response, err := client.Do(reqest)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	result = string(body)
	return
}
