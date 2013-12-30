package network

import (
	"github.com/liutong19890905/util/security"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HttpGetSimpleRequest(url string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func HttpGetAuthRequest(url string, username string, password string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)
	auth_str := username + ":" + password
	reqest.Header.Add("Authorization", "Basic "+security.EncodeBase64(auth_str))
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func HttpPostSimpleRequest(url string, data string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", url, strings.NewReader(data))
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}
