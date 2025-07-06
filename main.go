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

	compare := flag.Bool("comp",false, "Enter two city name to compare")
	city := flag.String("city", "", "City name to fetch weather for")
	save := flag.Bool("save", false, "Save the city as default")
	flag.Parse()

	args := flag.Args()

	if *compare {
		if len(args) != 2 {
			fmt.Println("❌ In compare mode (-comp), provide exactly two city names like:")
			fmt.Println("weatherman -comp \"London\" \"California\"")
			return
		}

		city1 := args[0]
		city2 := args[1]

		w1, err1 := weather.GetWeather(city1, apiKey)
		w2, err2 := weather.GetWeather(city2, apiKey)

		if err1 != nil || err2 != nil {
			fmt.Println("❌ Error fetching weather:")
			if err1 != nil {
				fmt.Printf(" - %s: %v\n", city1, err1)
			}
			if err2 != nil {
				fmt.Printf(" - %s: %v\n", city2, err2)
			}
			return
		}

		fmt.Printf("🌤️  Weather Comparison: %s vs %s\n\n", city1, city2)
		printWeatherComparison(w1, w2)
		return
	}

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

	printWeatherCondition(w)
}

func printWeatherCondition(w *weather.WeatherResponse) {
	fmt.Printf("🌤️  Weather in %s\n", w.Name)
	fmt.Println("------------------------------")
	fmt.Printf("Temperature : %.1f°C\n", w.Main.Temp)
	fmt.Printf("Feels Like  : %.1f°C\n", w.Main.FeelsLike)
	fmt.Printf("Humidity    : %d%%\n", w.Main.Humidity)
	fmt.Printf("Conditions  : %s\n", w.Weather[0].Description)
}

func printWeatherComparison(w1, w2 *weather.WeatherResponse) {
	fmt.Printf("%-15s | %-20s | %-20s\n", "Attribute", w1.Name, w2.Name)
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-15s | %-20.1f | %-20.1f\n", "Temperature (°C)", w1.Main.Temp, w2.Main.Temp)
	fmt.Printf("%-15s | %-20.1f | %-20.1f\n", "Feels Like (°C)", w1.Main.FeelsLike, w2.Main.FeelsLike)
	fmt.Printf("%-15s | %-20d | %-20d\n", "Humidity (%)", w1.Main.Humidity, w2.Main.Humidity)
	fmt.Printf("%-15s | %-20s | %-20s\n", "Conditions", w1.Weather[0].Description, w2.Weather[0].Description)
}