package main

import (
	"fmt"
	"net/http"
	"github.com/go-redis/redis/v8"
)

var client *redis.Client; 

func handler(w http.ResponseWriter, req *http.Request) {
	// Handle PUT requests 
	short, ok := req.URL.Query()["key"]
	long, ok2 := req.URL.Query()["value"]
	if ok && ok2 {
		handlePut(short[0], long[0])
	} else {
		short := req.URL.Path[1:]
		handleGet(w, req, short)
	}
}


func handleGet(w http.ResponseWriter, req *http.Request, key string) {
	value, err := client.Get(client.Context(), key).Result()
    if err != nil {
		http.NotFound(w, req)
	} else {
		w.Write([]byte(value))
	}
}

func handlePut(key string, value string) {
	client.Set(client.Context(), key, value, 0).Err()
}


func main() {

	client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	}) 

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Got error %s\n", err)
	}
}
