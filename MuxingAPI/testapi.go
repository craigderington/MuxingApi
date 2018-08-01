package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"log"
)


// api consumer function
func main() {
	getUrl := string("http://localhost:8080/api/books")
	postUrl := string("http://localhost:8080/api/books")

	response, err := http.Get(getUrl)
	if err != nil {
		fmt.Printf("The API GET request returned error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	jsonData := map[string]interface{} {
		"isbn": "12344590753455",
		"title": "Murder at 1600",
		"author": map[string]string{
			"firstname": "Margaret",
			"lastname": "Truman",
		},
	}

	jsonToBytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalln(err)
	}

	response, err = http.Post(postUrl, "application/json", bytes.NewBuffer(jsonToBytes))
	if err != nil {
		fmt.Printf("The API POST call returned error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
