package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://api.weatherapi.com/v1/forecast.json?key=8c392dfae4eb40b0abd132405210902&q=London&days=2"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
