package main

import (
	"flag"
	"github.com/bjorngi/go-smartercar/gps"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getGPS(gpsChan chan *gps.Location) {
	time.Sleep(60 * time.Second)
	payload := gps.Get("$GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*77")

	gpsChan <- payload
}

func GpsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		gpsChan := make(chan *gps.Location)

		go getGPS(gpsChan)
		payload := <-gpsChan

		conn.WriteJSON(payload)
		log.Printf("Sent %v", payload)
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/gps", GpsHandler)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
