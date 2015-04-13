package main

import (
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	privateKey []byte
)

func init() {
	privateKey, _ = ioutil.ReadFile("key.pem")
}

func main() {
	wg := new(sync.WaitGroup)
	start := time.Now()
	wg.Add(1)
	generateKey(wg, start, -1)
	log.Println("1 request took", time.Since(start))

	start = time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		generateKey(wg, start, i)
	}
	log.Println("10 request serially took", time.Since(start))

	start = time.Now()
	for i := 10; i < 20; i++ {
		wg.Add(1)
		go generateKey(wg, start, i)
	}
	wg.Wait()
	log.Println("10 request concurrently took", time.Since(start))
}

func generateKey(wg *sync.WaitGroup, start time.Time, client int) {

	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["key"] = "foo"
	token.Claims["exp"] = time.Now().Add(time.Hour)
	token.SignedString(privateKey)
	log.Printf("Client #%d took %s", client, time.Since(start))

	defer wg.Done()
}
