# `tempCalc` - Weather CLI Tool 🌦️

`tempCalc` is a simple command-line interface (CLI) tool written in Go that fetches real-time weather information for a specified city. It uses the [OpenWeatherMap API](https://openweathermap.org/api) to provide temperature data in Celsius. This lightweight utility is perfect for quick weather checks directly from your terminal.

---

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Prerequisites](#prerequisites)
- [Build Instructions](#build-instructions)
- [Making the CLI Global](#making-the-cli-global)
- [Example](#example)
- [License](#license)

---

## Features

- Fetch real-time weather data for any city.
- Output includes:
  - City name.
  - Current temperature (in Celsius).
- Easy-to-use CLI interface.
- Environment variable support for storing API keys securely.

---

## Installation

### Clone the Repository
```bash
git clone https://github.com/your-username/weather-cli.git
cd weather-cli
```

### Install Dependencies
Ensure you have [Go](https://go.dev/) installed. Run:
```bash
make deps
```

---

## Prerequisites

1. **Go**: Install Go by following the instructions on the [official Go website](https://go.dev/).
2. **OpenWeatherMap API Key**:
   - Sign up at [OpenWeatherMap](https://openweathermap.org/).
   - Generate an API key.
   - Store the key in a `.env` file at the root of the project:
     ```env
     OpenWeatherMapApiKey=your_api_key_here
     ```

---

## Build Instructions

To build the project:
```bash
make build
```

This will create a binary file named `tempCalc` in the `bin/` directory.

---

## Making the CLI Global

To make `tempCalc` available globally on your system:

1. Build the binary:
   ```bash
   make build
   ```

2. Move the binary to `/usr/local/bin`:
   ```bash
   sudo mv bin/tempCalc /usr/local/bin/
   ```

3. Verify the installation:
   ```bash
   tempCalc --help
   ```

---

## Usage

Run the CLI with the `--city` flag to specify the city name:
```bash
tempCalc --city="City Name"
```

### Example:
Fetching the weather for London:
```bash
tempCalc --city="London"
```
Output:
```
city: London
temperature: 15.00
```

---

## Example `.env` File

```env
OpenWeatherMapApiKey=your_api_key_here
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to contribute to the project by submitting issues or pull requests on GitHub!