package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomWeather() (float64, int, string) {
	conditions := []string{"Sunny", "Cloudy", "Rainy", "Stormy", "Snowy"}
	rand.Seed(time.Now().UnixNano())
	temp := 10 + rand.Float64()*30
	humidity := rand.Intn(100)
	condition := conditions[rand.Intn(len(conditions))]
	return temp, humidity, condition
}

func main() {
	var city string
	fmt.Print("Enter city name: ")
	fmt.Scanln(&city)

	temp, humidity, condition := randomWeather()
	fmt.Printf("Weather in %s: %.2f°C, %d%% humidity, %s\n", city, temp, humidity, condition)
}
