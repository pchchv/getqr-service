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

// Load values from .env into the system
func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

// Getting a value. Outputs a panic if the value is missing
func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func getQR(data string) string {
	name := strings.Split(strings.Split(data, "/")[2], ".")[0]
	err := getqr.WriteFile(data, getqr.Medium, 256, name+".png")
	if err != nil {
		log.Panic(err)
	}
	qrPath, err := filepath.Abs(filepath.Base("qr.png"))
	if err != nil {
		log.Panic(err)
	}
	return qrPath
}

func main() {
	server()
}
