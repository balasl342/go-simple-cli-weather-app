# Weather CLI

A simple command-line tool written in Go to fetch and display real-time weather information from OpenWeatherMap API.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [License](#license)

## Introduction

CLI Weather App is a Go-based application that retrieves weather data for a specified city using the OpenWeatherMap API. The application displays the temperature, humidity, and weather description in a user-friendly format.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [OpenWeatherMap API Key](https://openweathermap.org/api) (sign up to get a free API key)

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/balasl342/go-simple-cli-weather-app.git

2. **Create configuration file:**
    ```
    Create a config.json file with the necessary configuration below.
    {
        "api_key": "xxxxxxxxx"
    }

3. **Build the application:**
    ```sh
    go build -o weathercli

## License
- This project is licensed under the MIT License - see the LICENSE file for details.