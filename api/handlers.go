package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"slices"
	"weatherApi/data"
	"os"

	"github.com/gin-gonic/gin"
)

var key string

func init(){

	data, err := os.ReadFile("/home/arya/Desktop/weatherApi/dsn.json")
   
   if err != nil{
	   log.Fatal(err.Error())
   }

   var getdata map[string]string

   err = json.Unmarshal(data, &getdata)

   key = getdata["key"]

   if err != nil{
	   log.Fatal(err.Error())
   }
}


func forecast(c *gin.Context){
	
	legalParms := []string{"us", "uk", "metric", "base", "days", "hours", "minutes", "alerts",
	 "current", "evurlents", "obs","remote", "fcst", "stats", "statsfcst", "ar", "bg", "cs", "da",
	  "de", "el", "en", "es", "fa", "fi", "fr", "he", "it", "ja", "ko", "nl", "pl", "hu", "ru", 
	  "sk", "pt", "ru", "sk", "sr", "sv", "tr", "uk", "vi", "zh"}

	includeQuery := c.QueryArray("include")
	unitGroup := c.Query("unitGroup")
	lang := c.Query("lang")
	elemnts := c.QueryArray("elements")
	day1 := c.Query("day1")
	day2 := c.Query("day2")

	existError := slices.ContainsFunc(includeQuery, func(s string) bool {

		return !slices.Contains(legalParms, s)

	}) || slices.ContainsFunc(elemnts, func(s string) bool {

		return !slices.Contains(legalParms, s)}) ||

	!slices.Contains(legalParms, unitGroup) || !slices.Contains(legalParms, lang)
	

	if existError {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query is invalid"})
		return 
	}

	include := ""

	for _, data := range includeQuery{
		include += data 
	}

	elems := ""

	for _, data := range elemnts{
		elems += data 
	}

	days := ""

	if day1 != ""{days += "/" + day1}
	if day2 != ""{days += "/" + day2}


	baseUrl := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/tehran"
	queryParam := url.Values{}

	baseUrl += days

	queryParam.Add("include", include)
	queryParam.Add("unitGroup", unitGroup)
	queryParam.Add("lang", lang)
	queryParam.Add("elems", elems)

	baseUrl = baseUrl + "?" + queryParam.Encode() + fmt.Sprintf("&key=%s", key)

	response, err := http.Get(baseUrl)

	if err != nil{
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	datas, err := io.ReadAll(response.Body)

	if err != nil{
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	var income data.Result

	err = json.Unmarshal(datas, &income)

	if err != nil{
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, income)
}