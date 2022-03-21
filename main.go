package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var wind, water int

func main() {
	go reroll()
	server := gin.Default()
	data := map[string]interface{}{
		"wind":  &wind,
		"water": &water,
	}

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"status": data,
		})
	})

	server.Run(":8080")
}

func reroll() {
	rand.Seed(time.Now().UnixNano())
	for {
		wind = rand.Intn(100)
		water = rand.Intn(100)
		fmt.Println(wind)
		fmt.Println(water)
		time.Sleep(5 * time.Second)
	}
}
