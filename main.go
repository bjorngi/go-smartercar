package main

import (
	"flag"
	"fmt"
	"github.com/bjorngi/go-smartercar/gps"
	"github.com/bjorngi/go-smartercar/serial"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var addr = flag.String("addr", ":8080", "http service address")

var gpsChan chan *gps.Location = make(chan *gps.Location)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SerialReader() {
	serialChan := make(chan string)
	go serial.ReadArduino(serialChan)

	for {
		input := <-serialChan
		SelectParser(input)
	}

}

func SelectParser(input string) {
	arr := strings.Split(input, ",")
	switch arr[0] {
	case "$GPRMC":
		if gps.Checksum(input) {
			gpsData, err := gps.Get(arr)
			if err != nil {
				log.Println(err)
			} else {
				gpsChan <- gpsData
			}
		} else {
			log.Println("Checksum failed")
		}
	}
}

func GpsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		gpsData := <-gpsChan
		log.Println("Data sent")
		conn.WriteJSON(gpsData)
	}
}

func main() {
	fmt.Printf("Server started\n")

	defer fmt.Printf("Server stopped\n")

	go SerialReader()

	flag.Parse()
	http.HandleFunc("/gps", GpsHandler)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
