package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)
var locationAPI string = "https://pokeapi.co/api/v2/location/"

func commandMap() error {
	res, err := http.Get(locationAPI)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locate []string
	if err = json.Unmarshal(body, &locate); err != nil {
		return nil
	}
	fmt.Printf("%s", body)
	fmt.Printf("%s", locate)

	return nil
}