package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Function to generate a Base64 string to sign hash - run once and add string do .env
// func init() {
// 	key := make([]byte, 64)
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal((err))
// 	}
// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Load()

	r := router.Generate()
	fmt.Printf("Listening to port %d!\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
