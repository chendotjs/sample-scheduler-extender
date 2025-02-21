package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/cnych/sample-scheduler-extender/controller"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.POST("/filter", controller.Filter)
	router.POST("/prioritize", controller.Prioritize)

	log.Printf("start up sample-scheduler-extender!\n")
	log.Fatal(http.ListenAndServe(":8888", router))
}
