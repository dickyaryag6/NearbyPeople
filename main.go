package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	// coordinate of someone who needs people nearby
	lat         float64 = 12.9611159
	lon         float64 = 77.6362214
	PI          float64 = 3.14159265358979323846
	earthRadius float64 = 6371.0
)

type Customers struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
}

type Answers struct {
	Answers []Answer `json:"answers"`
}

type Answer struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func degreetoradian(degree float64) float64 {
	return (degree * PI / 180)
}

func distanceearth(lat1 string, lon1 string) float64 {
	lat_other, _ := strconv.ParseFloat(lat1, 64)
	lon_other, _ := strconv.ParseFloat(lon1, 64)

	lat_other = degreetoradian(lat_other)
	lon_other = degreetoradian(lon_other)
	lat_searcher := degreetoradian(lat)
	lon_searcher := degreetoradian(lon)

	// menghitung jarak
	centralAng := math.Acos(math.Sin(lat_searcher)*
		math.Sin(lat_other) + math.Cos(lat_searcher)*
		math.Cos(lat_other)*math.Cos(lon_other-lon_searcher))

	return (earthRadius * centralAng)
}

func main() {
	// READ JSON FILE
	jsonfile, err := os.Open("customers.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonfile.Close()

	//read our json as byte array
	bytevalue, _ := ioutil.ReadAll(jsonfile)

	var customers Customers

	json.Unmarshal(bytevalue, &customers)

	var answers []Answer

	for i := 0; i < len(customers.Customers); i++ {
		if distanceearth(customers.Customers[i].Latitude, customers.Customers[i].Longitude) <= 50.000 {
			answers = append(answers, Answer{
				UserID: customers.Customers[i].UserID,
				Name:   customers.Customers[i].Name})
		}
	}

	jsondata := &Answers{Answers: answers}

	file, _ := json.MarshalIndent(jsondata, "", "    ")

	_ = ioutil.WriteFile("answer.json", file, 0644)

}
