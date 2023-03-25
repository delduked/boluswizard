package controllers

import (
	"fmt"
	"log"

	websocket "github.com/gofiber/websocket/v2"
)

var wsChan = make(chan string)
var clients = make(map[*websocket.Conn]string)

func Ws(c *websocket.Conn) {

	// c.Locals is added to the *websocket.Conn
	log.Println(c.Locals("allowed"))  // true
	log.Println(c.Params("id"))       // 123
	log.Println(c.Query("v"))         // 1.0
	log.Println(c.Cookies("session")) // ""

	clients[c] = ""

	//go ListenForWs(c)
	var (
		msg []byte
		err error
	)
	for {
		if _, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}

		wsChan <- string(msg)

		// if err = c.WriteMessage(mt, msg); err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
	}

}

func ListenToChannel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	for {

		test := <-wsChan
		log.Printf("recv: %s", test)
	}
}
