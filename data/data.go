package data

import "encoding/json"

type Result struct{

	QueryCost json.Number	`json:"queryCost"`
	Latitude json.Number	`json:"latitude"`
    Longitude json.Number	`json:"longitude"`
    ResolvedAddress string	`json:"resolvedAddress"`
    Address string			`json:"address"`
    Timezone string			`json:"timezone"`
    Tzoffset json.Number	`json:"tzoffset"`
	Description string		`json:"description"`
	Days []Day				`json:"days"`
	CurrentConditions Condition 

}

// type Information struct{

// 	QueryCost json.Number	`json:"queryCost"`
// 	Latitude json.Number	`json:"latitude"`
//     Longitude json.Number	`json:"longitude"`
//     ResolvedAddress string	`json:"resolvedAddress"`
//     Address string			`json:"address"`
//     Timezone string			`json:"timezone"`
//     Tzoffset string			`json:"tzoffset"`
// }

type Day struct{

	Datetime string			
	DatetimeEpoch json.Number	
	Tempmax json.Number			
	Tempmin json.Number			
	Temp json.Number
	Feelslikemax json.Number
	Feelslikemin json.Number
	Feelslike json.Number
	Dew json.Number
	Humidity json.Number
	Precip json.Number
	Precipprob json.Number
	Precipcover json.Number
	Preciptype any
	Snow json.Number
	Snowdepth json.Number
	Windgust json.Number
	Windspeed json.Number
	Winddir json.Number
	Pressure json.Number
	Cloudcover json.Number
	Visibility json.Number
	Solarradiation json.Number
	Solarenergy json.Number
	Uvindex json.Number
	Severerisk json.Number
	Sunrise string
	SunriseEpoch json.Number
	Sunset string
	SunsetEpoch json.Number
	Moonphase json.Number
	Conditions string
	Description string
	Icon string
	Source string
	Hours []Hour		`json:"hours"`
}

type Hour struct{

	Datetime string
	DatetimeEpoch json.Number
	Temp json.Number
	Feelslike json.Number
	Humidity json.Number
	Dew json.Number
	Precip json.Number
	Precipprob json.Number
	Snow json.Number
	Snowdepth json.Number
	Preciptype any
	Windgust json.Number
	Windspeed json.Number
	Winddir json.Number
	Pressure json.Number
	Visibility json.Number
	Cloudcover json.Number
	Solarradiation json.Number
	Solarenergy json.Number
	Uvindex json.Number
	Severerisk json.Number
	Conditions string
	Icon string
	Source string
}

type Condition struct{

	Datetime string
	DatetimeEpoch json.Number
	Temp json.Number
	Feelslike json.Number
	Humidity json.Number
	Dew json.Number
	Precip json.Number
	Precipprob json.Number
	Snow json.Number
	Snowdepth json.Number
	Preciptype json.Number
	Windgust json.Number
	Windspeed json.Number
	Winddir json.Number
	Pressure json.Number
	Visibility json.Number
	Cloudcover json.Number
	Solarradiation json.Number
	Solarenergy json.Number
	Uvindex json.Number
	Severerisk json.Number
	Conditions string
	Icon string
	Source string
	Sunrise string
    SunriseEpoch json.Number
    Sunset string
    SunsetEpoch json.Number
    Moonphase json.Number
}