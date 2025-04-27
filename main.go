package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Fahim047/weatherman/weather"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("❌ Error loading .env file:", err)
		return
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("❌ API_KEY not set in .env")
		return
	}

	city := flag.String("city", "", "City name to fetch weather for")
	save := flag.Bool("save", false, "Save the city as default")
	flag.Parse()

	if *city == "" {
		defaultCity, err := weather.LoadDefaultCity()
		if err != nil || defaultCity == "" {
			fmt.Println("❌ No default city set. Please provide a city using -city flag")
			return
		}
		*city = defaultCity
	}

	if *save {
		err := weather.SaveDefaultCity(*city)
		if err != nil {
			fmt.Println("⚠️ Failed to save default city:", err)
		} else {
			fmt.Println("✅ Default city saved successfully!")
		}
	}

	w, err := weather.GetWeather(*city, apiKey)
	if err != nil {
		fmt.Println("❌ Error fetching weather:", err)
		return
	}

	fmt.Printf("🌤️  Weather in %s\n", w.Name)
	fmt.Printf("Temperature: %.1f°C\n", w.Main.Temp)
	fmt.Printf("Feels Like: %.1f°C\n", w.Main.FeelsLike)
	fmt.Printf("Humidity: %d%%\n", w.Main.Humidity)
	fmt.Printf("Conditions: %s\n", w.Weather[0].Description)
}
