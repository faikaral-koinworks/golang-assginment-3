package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wind, water int
var isSent bool = true

var upgrader=websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin:func(r *http.Request)bool{
		return true
	},
}

func main() {
	go reroll()
	server := gin.Default()
	data := map[string]interface{}{
		"wind":  &wind,
		"water": &water,
	}

	server.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"status": data,
		})
	
	})
	server.GET("/ws",wsEndPoint)
	server.GET("/",func(c *gin.Context) {
		c.File("./index.html")
	})

	server.Run(":8080")
}

func reroll() {
	rand.Seed(time.Now().UnixNano())
	for {
		wind = rand.Intn(100)
		water = rand.Intn(100)
		isSent=false
		time.Sleep(15*time.Second)
	}
}


func wsEndPoint(c *gin.Context){
	ws,err:=upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil{
		panic(err)
	}

	data := map[string]interface{}{
		"wind":  &wind,
		"water": &water,
	}

	defer ws.Close()

	for {
		for isSent==false{
			err:=ws.WriteJSON(data)
			if err != nil{
				panic(err)
			}
			isSent=true
		}
		if err !=nil{
			log.Fatal(err)
		}
	}


}
