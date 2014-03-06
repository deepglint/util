package client

import (
	"deepglint/util/security"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const Timeout = 5

func Get(url_addr string) (result string, err error) {
	urler := url.URL{}
	url_proxy, _ := urler.Parse(url_addr)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*Timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(10 * time.Second))
				return c, nil
			},
			Proxy: http.ProxyURL(url_proxy),
		},
	}
	reqest, err := http.NewRequest("GET", url_addr, nil)
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

func AuthGet(url_addr string, username string, password string) (result string, err error) {
	urler := url.URL{}
	url_proxy, _ := urler.Parse(url_addr)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*Timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(10 * time.Second))
				return c, nil
			},
			Proxy: http.ProxyURL(url_proxy),
		},
	}
	reqest, err := http.NewRequest("GET", url_addr, nil)
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

func AuthPost(url_addr string, data string, username string, password string) (result string, err error) {
	urler := url.URL{}
	url_proxy, _ := urler.Parse(url_addr)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*Timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(10 * time.Second))
				return c, nil
			},
			Proxy: http.ProxyURL(url_proxy),
		},
	}
	reqest, err := http.NewRequest("POST", url_addr, strings.NewReader(data))
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

func Post(url_addr string, data string) (result string, err error) {
	urler := url.URL{}
	url_proxy, _ := urler.Parse(url_addr)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*Timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(10 * time.Second))
				return c, nil
			},
			Proxy: http.ProxyURL(url_proxy),
		},
	}
	reqest, err := http.NewRequest("POST", url_addr, strings.NewReader(data))
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
