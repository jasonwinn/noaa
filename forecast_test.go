package noaa

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestForecast(t *testing.T) {
	forecast := &Forecast{Length: 5, ForecastDays: make([]ForecastDay, 5)}

	file, _ := ioutil.ReadFile(filepath.Join("test_data", "forecast.xml"))
	forecast.assignForecastData(file)

	if forecast.Length != 5 {
		t.Error(fmt.Sprintf("Expected %d Day Forecast, ~ Received %d Day Forecast", 5, forecast.Length))
	}

	// Check Forecast Data
	day := forecast.ForecastDays[0]

	if day.NameDay != "Today" {
		t.Error(fmt.Sprintf("Expected Day Name %s, ~ Received Day Name %s", "Today", day.NameDay))
	}

	if day.NameNight != "Tonight" {
		t.Error(fmt.Sprintf("Expected Night Name %s, ~ Received Night Name %s", "Tonight", day.NameNight))
	}

	if day.StartTime.String() != "2013-08-12 06:00:00 -0800 -0800" {
		t.Error(fmt.Sprintf("Expected Start Time %s, ~ Received Start Time %s", "2013-08-12 06:00:00 -0800 -0800", day.StartTime))
	}

	if day.EndTime.String() != "2013-08-13 06:00:00 -0800 -0800" {
		t.Error(fmt.Sprintf("Expected End Time %s, ~ Received End Time %s", "2013-08-13 06:00:00 -0800 -0800", day.EndTime))
	}

	if day.MaxTemperature != 49 {
		t.Error(fmt.Sprintf("Expected Max Temp %d, ~ Received Max Temp %d", 49, day.MaxTemperature))
	}

	if day.MinTemperature != 40 {
		t.Error(fmt.Sprintf("Expected Min Temp %d, ~ Received Min Temp %d", 40, day.MaxTemperature))
	}

	if day.PrecipitationChanceDay != 75 {
		t.Error(fmt.Sprintf("Expected Precip Chance %d, ~ Received Precip Chance %f", 75, day.PrecipitationChanceDay))
	}

	if day.PrecipitationChanceNight != 0 {
		t.Error(fmt.Sprintf("Expected Precip Chance %d, ~ Received Precip Chance %f", 0, day.PrecipitationChanceNight))
	}

	if day.SummaryDay["summary"] != "Drizzle" {
		t.Error(fmt.Sprintf("Expected Day Summary %s, ~ Received Day Summary %s", "Drizzle", day.SummaryDay["summary"]))
	}

	if day.SummaryDay["icon"] != "http://www.nws.noaa.gov/weather/images/fcicons/ra.jpg" {
		t.Error(fmt.Sprintf("Expected Day Icon %s, ~ Received Day Icon %s",
			"http://www.nws.noaa.gov/weather/images/fcicons/ra.jpg", day.SummaryDay["icon"]))
	}

	if day.SummaryNight["summary"] != "Mostly Cloudy" {
		t.Error(fmt.Sprintf("Expected Night Summary %s, ~ Received Night Summary %s", "Mostly Cloudy", day.SummaryNight["summary"]))
	}

	if day.SummaryNight["icon"] != "http://www.nws.noaa.gov/weather/images/fcicons/nbkn.jpg" {
		t.Error(fmt.Sprintf("Expected Night Icon %s, ~ Received Night Icon %s",
			"http://www.nws.noaa.gov/weather/images/fcicons/nbkn.jpg", day.SummaryNight["icon"]))
	}

	// Day Conditions
	if day.ConditionsDay[0].Coverage != "areas" {
		t.Error(fmt.Sprintf("Expected Day Coverage %s, ~ Received Day Coverage %s", "areas", day.ConditionsDay[0].Coverage))
	}

	if day.ConditionsDay[0].Intensity != "light" {
		t.Error(fmt.Sprintf("Expected Day Intensity %s, ~ Received Day Intensity %s", "light", day.ConditionsDay[0].Intensity))
	}

	if day.ConditionsDay[0].WeatherType != "drizzle" {
		t.Error(fmt.Sprintf("Expected Day Weather Type %s, ~ Received Day Weather Type %s", "drizzle", day.ConditionsDay[0].WeatherType))
	}

	if day.ConditionsDay[0].Qualifier != "none" {
		t.Error(fmt.Sprintf("Expected Day Qualifier %s, ~ Received Day Qualifier %s", "none", day.ConditionsDay[0].Qualifier))
	}

	// Night Conditions
	if len(day.ConditionsNight) > 0 {
		t.Error(fmt.Sprintf("Exptected No Night Conditions, ~ Received Conditions: %s", day.ConditionsNight))
	}

}
