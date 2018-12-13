package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// post请求 json格式
func JsonPostRequest(url string, headerMap, params map[string]interface{}) string {
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error() )
	}
	reader := bytes.NewReader(bytesData)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, reader)
	header := make(http.Header)
	if len(headerMap) > 0 {
		for k, v := range headerMap {
			header.Set(k, v.(string))
		}
	}
	request.Header = header
	result, err := client.Do(request)
	defer result.Body.Close()
	if err != nil {
		log.Println(url + ".error")
	}
	if result.StatusCode != 200 {
		log.Println(url + ".error")
	}
	body, _ := ioutil.ReadAll(result.Body)

	return string(body)
}

// post请求 form格式
func FormPostRequest(url1 string, headerMap, params map[string]interface{}) string {
	postValue := url.Values{}
	for key, value := range params{
		postValue.Set(key, value.(string))
	}
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url1, strings.NewReader(postValue.Encode()))
	header := make(http.Header)
	if len(headerMap) > 0 {
		for k, v := range headerMap {
			header.Set(k, v.(string))
		}
	}
	request.Header = header
	result, err := client.Do(request)
	defer result.Body.Close()
	if err != nil {
		log.Println(url1 + ".error")
	}
	if result.StatusCode != 200 {
		log.Println(url1 + ".error")
	}
	body, _ := ioutil.ReadAll(result.Body)

	return string(body)
}
