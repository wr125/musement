package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wr125/musement/handler"
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
type name struct {
	Lat  float64 `xml:"latitude"`
	Long float64 `xml:"longitude"`
}
type city struct {
	Name name `xml:"name"`
}

func getCityMusement(Name *name) error {

	url := fmt.Sprintf("https://api.musement.com/api/v3/cities")
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	return xml.Unmarshal(body, &Name)
}

func makeWeatherAPIQuery(lat, long float64, weather *weather) error {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=8c392dfae4eb40b0abd132405210902&q=%v,%v&days=1", lat, long)
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

func main() {

	// TODO: get all cities from Musement API; get lat/long for those cities
	// For example: London is lat = 51.52, long = -0.11
	weather := weather{}
	//getCityMusement(name.Lat,name.Long)

	makeWeatherAPIQuery(51.52, -0.11, &weather)

	fmt.Printf("Processed city [%v] | [%v] - [%v]\n",
		weather.Location.Name,
		weather.Current.Condition.Text,
		weather.Forecast.Forecastday[0].Day.Condition.Text,
	)
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/result", handler.Search)

	fmt.Println("Server is running..")
	http.ListenAndServe(":4000", nil)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

}
