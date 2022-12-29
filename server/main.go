package main

import (
	"fmt"
	"net/http"
	"github.com/go-redis/redis/v8"
)

var client *redis.Client; 

func handler(w http.ResponseWriter, req *http.Request) {
	// Handle PUT requests 
	key, ok := req.URL.Query()["key"]
	value, ok2 := req.URL.Query()["value"]
	if ok && ok2 {
		client.Set(client.Context(), key[0], value[0], 0).Err()
	} else {
		key := req.URL.Path[1:]
		value, err := client.Get(client.Context(), key).Result()
		if err != nil {
			http.NotFound(w, req)
		} else {
			w.Write([]byte(value))
		}
	}
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
		fmt.Printf("Cannot start server. Error: %s\n", err)
	}
}
