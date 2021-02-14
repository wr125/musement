package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEverything(t *testing.T) {
	weatherAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockWeatherResponse)
	}))
	defer weatherAPI.Close()
	musementAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockMusementResponse)
	}))
	defer weatherAPI.Close()

	printedResults := []string{}
	printerRecorder := func(s string) { printedResults = append(printedResults, s) }

	run(musementAPI.URL, weatherAPI.URL, printerRecorder)

	if len(printedResults) != 1 {
		t.Errorf("Run should have printed 1 line but printed %v lines!", len(printedResults))
	}

	expected := "Processed city [De Wallen] | [Sunny] - [Partly cloudy]"
	if printedResults[0] != expected {
		t.Errorf("Run should have printed [%v] line but printed [%v]!", expected, printedResults[0])
	}
}

var mockMusementResponse = `[
	{
	  "id": 57,
	  "top": false,
	  "name": "Amsterdam",
	  "code": "amsterdam",
	  "content": "Amsterdam is Holland’s capital city, a cultural hub and one of the Europe’s favorite travel destinations. It may not be an enormous city but its unique features are enough to attract tourists for day trips and longer holidays.Built under sea level, it is defined a ‘Venice of the North’ due to its many canals – it was a city built in the year 1000, reclaiming the area from advancing waters. Some of the most popular attractions in Amsterdam are the Rijksmuseum with Rembrant's famous painting 'The Night Watch', the Van Gogh Museum, Anne Frank's House, Museum Willet, the Cromhout Houses and many more.",
	  "meta_description": "Book your tickets for museums, tours and attractions in Amsterdam. Discover the Rijksmuseum, sip a beer at the Heineken Experience or take a tour on the canals.",
	  "meta_title": "Things to do in Amsterdam: Museums, tours, and attractions",
	  "headline": "Things to do in Amsterdam",
	  "more": "Young people come to this city for its nightlife and possibly to see the world renown “Coffee Shops”, while art-lovers on the other hand come to enjoy some of the city’s museums and the beautiful architecture. Holland successfully made its very own Renaissance architecture in the 17th century, giving Amsterdam its very own unique atmosphere.",
	  "weight": 20,
	  "latitude": 52.374,
	  "longitude": 4.9,
	  "country": {
		"id": 124,
		"name": "Netherlands",
		"iso_code": "NL"
	  },
	  "cover_image_url": "https://images.musement.com/cover/0002/15/amsterdam_header-114429.jpeg",
	  "url": "https://www.musement.com/us/amsterdam/",
	  "activities_count": 219,
	  "time_zone": "Europe/Amsterdam",
	  "list_count": 1,
	  "venue_count": 22,
	  "show_in_popular": true
	}
  ]`

var mockWeatherResponse = `
{
	"location": {
	  "name": "De Wallen",
	  "region": "North Holland",
	  "country": "Netherlands",
	  "lat": 52.37,
	  "lon": 4.9,
	  "tz_id": "Europe/Amsterdam",
	  "localtime_epoch": 1613319333,
	  "localtime": "2021-02-14 17:15"
	},
	"current": {
	  "last_updated_epoch": 1613318458,
	  "last_updated": "2021-02-14 17:00",
	  "temp_c": 3.0,
	  "temp_f": 37.4,
	  "is_day": 1,
	  "condition": {
		"text": "Sunny",
		"icon": "//cdn.weatherapi.com/weather/64x64/day/113.png",
		"code": 1000
	  }
	},
	"forecast": {
	  "forecastday": [
		{
		  "date": "2021-02-14",
		  "date_epoch": 1613260800,
		  "day": {
			"condition": {
			  "text": "Partly cloudy",
			  "icon": "//cdn.weatherapi.com/weather/64x64/day/116.png",
			  "code": 1003
			},
			"uv": 1.0
		  }
		}
	  ]
	}
  }
`
