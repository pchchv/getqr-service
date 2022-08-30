package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pchchv/getqr"
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

func getQR(data string) (string, string) {
	name := strings.Split(strings.Split(data, "/")[2], ".")[0]
	err := getqr.WriteFile(data, getqr.Medium, 256, name+".png")
	if err != nil {
		log.Panic(err)
	}
	qrPath, err := filepath.Abs(filepath.Base("qr.png"))
	if err != nil {
		log.Panic(err)
	}
	return qrPath, name
}

func main() {
	server()
}
