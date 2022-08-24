package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	testURL    string
	testClient http.Client
)

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing.
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func getQR(data string) string {
	var qrPath string
	// TODO: Implement getting a QR code
	// TODO: Implement saving a qr code file
	return qrPath
}

func main() {
	server()
}
