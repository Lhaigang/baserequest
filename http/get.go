package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequest(url string, headerMap map[string]interface{})  string{
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
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
