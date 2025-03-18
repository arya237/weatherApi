package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weatherApi/data"
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

func getForcastNext15Days(location string) (data.Result, error) {
	
	// var income map[string]json.RawMessage
	var income data.Result

	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?key=%s", location, key)

	response, err := http.Get(url)

	if err != nil{
		log.Print(err.Error())
		return data.Result{}, err
	}

	datas, err := io.ReadAll(response.Body)

	if err != nil{
		log.Print(err.Error())
		return data.Result{} ,err
	}

	err = json.Unmarshal(datas, &income)

	if err != nil{
		log.Print(err.Error())
		return data.Result{}, err
	}

	return income, nil
}