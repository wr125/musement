package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type condition struct {
	Text string `json:"text"`
}

type weatherSection struct {
	Condition condition `json:"condition"`
}

type forecastday struct {
	Day weatherSection `json:"day"`
}

type forecast struct {
	Forecastday []forecastday `json:"forecastday"`
}

type location struct {
	Name string `json:"name"`
}

type weather struct {
	Location location       `json:"location"`
	Current  weatherSection `json:"current"`
	Forecast forecast       `json:"forecast"`
}
type city struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func getCityMusement(musementAPIURL string, cities *[]city) error {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, musementAPIURL, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, cities)
}

func makeWeatherAPIQuery(weatherAPIURL string, lat, long float64, weather *weather) error {
	url := fmt.Sprintf("%v/v1/forecast.json?key=8c392dfae4eb40b0abd132405210902&q=%v,%v&days=1", weatherAPIURL, lat, long)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return json.Unmarshal(body, weather)
}

func run(musementAPIURL, weatherAPIURL string, print func(s string)) {
	cities := []city{}
	err := getCityMusement(musementAPIURL, &cities)
	if err != nil {
		log.Fatal(err)
	}

	for _, city := range cities {
		weather := weather{}
		makeWeatherAPIQuery(weatherAPIURL, city.Latitude, city.Longitude, &weather)
		output := fmt.Sprintf("Processed city [%v] | [%v] - [%v]",
			weather.Location.Name,
			weather.Current.Condition.Text,
			weather.Forecast.Forecastday[0].Day.Condition.Text,
		)
		print(output)
		// Don't get rate-limited!
		time.Sleep(1 * time.Second)

	}
}

func main() {
	musementAPIURL := "https://api.musement.com/api/v3/cities"
	weatherAPIURL := "http://api.weatherapi.com"
	print := func(s string) { fmt.Println(s) }

	run(musementAPIURL, weatherAPIURL, print)
}
