package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Ocelani/medexp/pkg/utils"
	"github.com/joho/godotenv"
)

var after string

func runQuery(after string) (buffer bytes.Buffer) {
	q := utils.Query(after)

	// Prepares the json stub to send
	query, err := json.Marshal(q)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(bytes.NewBuffer(query))

	// Bearer token auth header
	var bearer = "Bearer " + os.Getenv("GITHUB_TOKEN")

	// Creates the request
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(query))
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}

	// Receives resp from the req
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	// Closes the response when it has ended
	defer resp.Body.Close()

	// Collects the data
	data, _ := ioutil.ReadAll(resp.Body)

	json.Indent(&buffer, data, "", "\t")
	writeResult(&buffer)
	return
}

func writeResult(buffer *bytes.Buffer) {
	// Writes into a json file
	err := ioutil.WriteFile("out/repos.json", buffer.Bytes(), 0777)
	if err != nil {
		log.Println("Error trying to write file.\n[ERRO] -", err)
	}
}

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	runQuery(after)
}
