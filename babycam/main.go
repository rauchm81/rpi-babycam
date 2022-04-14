package main

import (
	"log"
	"net/http"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pin rpio.Pin
)

func ledOn(w http.ResponseWriter, r *http.Request) {
	log.Println("led On")
	pin.High()
}

func ledOff(w http.ResponseWriter, r *http.Request) {
	log.Println("led Off")
	pin.Low()
}

func main() {

	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}

	pin = rpio.Pin(4)
	pin.Output()

	http.HandleFunc("/api/led/on", ledOn)
	http.HandleFunc("/api/led/off", ledOff)
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":5000", nil))
}
