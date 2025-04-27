# ☀️ Weatherman CLI

A simple command-line weather app written in Go. Fetches current weather information for any city using the [OpenWeatherMap API](https://openweathermap.org/api).

---

## 🚀 Features

- Get current weather by city name
- Displays temperature, humidity, and weather description
- Loads API key from `.env` file
- Easy-to-use command-line interface

---

## 🛠️ Tech Stack

- **Go** (Golang)
- **OpenWeatherMap API**
- **joho/godotenv** for environment variable management

---

## 📦 Installation

### 1. Clone the repository

```bash
git clone https://github.com/Fahim047/weatherman.git
cd weatherman
```

### 2. Create a .env file in the root:

```bash
API_KEY=your_openweathermap_api_key_here
```

### 3. Install dependencies

```bash
go mod tidy
```

## 🧪 Usage

Fetch weather for a city:

```bash
go run main.go -city "London"
```

Example Output:

```makefile
🌤️  Weather in London
Temperature: 14.2°C
Feels Like: 15°C
Humidity: 67%
Conditions: broken clouds
```

### 🔧 Flags

| Flag    | Description                    |
| ------- | ------------------------------ |
| `-city` | (Required) City name to lookup |
| `-save` | (Optional) Save default city   |

### ⚙️ .env Setup

Make sure you create a .env file in the root directory:

```bash
API_KEY=your_openweathermap_api_key_here
```

And add it to your .gitignore:

```bash
.env
```

## 🙋‍♂️ Author

Made with ❤️ and Go by [Fahimul Islam](https://fifolio.vercel.app)

## 📄 License

MIT License — feel free to use and build on it.
