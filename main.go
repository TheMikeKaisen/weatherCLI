package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func query(city string) (weatherData, error) {
	// load env
	key := os.Getenv("OpenWeatherMapApiKey")

	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?APPID=" + key + "&q=" + city)
	if err != nil {
		log.Println("Error while getting the data")
		return weatherData{}, err
	}
	defer response.Body.Close()

	var data weatherData
	// parse the response
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println("Error while decoding the data!")
		return weatherData{}, err
	}

	return data, nil

}

func main() {
	// load the env
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error occurred while loading env.")
		os.Exit(1)
	}

	var cityName string

	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "A sample CLI",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := query(cityName)
			if err != nil {
				log.Println("Error while querying the data!")
				os.Exit(1)
			}
			name := result.Name
			tempInCalcius := result.Main.Kelvin - 273
			fmt.Printf("city: %s\n", name)
			fmt.Printf("temperature: %f %s\n", tempInCalcius, " Celsius")
		},
	}

	rootCmd.PersistentFlags().StringVar(&cityName, "city", "", "name of the city")
	rootCmd.Execute()

}
