package weather

import (
	"encoding/json"
	"os"
)

func SaveDefaultCity(city string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := home + "/.weathermanrc"

	config := map[string]string{"default_city": city}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

func LoadDefaultCity() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := home + "/.weathermanrc"

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var config map[string]string
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return "", err
	}

	return config["default_city"], nil
}
