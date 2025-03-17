package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var key any

func init(){

 	data, err := os.ReadFile("/home/arya/Desktop/weatherApi/dsn.json")
	
	if err != nil{
		log.Fatal(err.Error())
	}

	var getdata map[string]any

	err = json.Unmarshal(data, &getdata)

	key = getdata["key"]

	if err != nil{
		log.Fatal(err.Error())
	}
}

func getForcastNext15Days(location string) (map[string]json.RawMessage, error) {
	
	var income map[string]json.RawMessage

	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?key=%s", location, key)

	fmt.Print("*****", url)

	response, err := http.Get(url)

	if err != nil{
		log.Print(err.Error())
		return nil, err
	}

	data, err := io.ReadAll(response.Body)

	if err != nil{
		log.Print(err.Error())
		return nil, err
	}

	err = json.Unmarshal(data, &income)

	if err != nil{
		log.Print(err.Error())
		return nil, err
	}

	return income, nil
}