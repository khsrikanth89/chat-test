
package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"fmt"
	"time"
)
func main() {


origin := "http://localhost:4000/"
url := "ws://localhost:4000/sock"
ws := make([]*websocket.Conn,100)
var err error
start := time.Now()

for i := 0; i<100 ; i++ {
	ws[i], err = websocket.Dial(url, "", origin)
	if err != nil {
	    log.Fatal("Dial Failed: ",err)
	}
}
duration := time.Now().Sub(start)
fmt.Println("Time",duration)

if _, err := ws[0].Write([]byte("hello, world!\n")); err != nil {
    log.Fatal(err)
}
var msg = make([]byte, 512)
var n int
if n, err = ws[1].Read(msg); err != nil {
    log.Fatal(err)
}
fmt.Printf("Received: %s.\n", msg[:n])

}