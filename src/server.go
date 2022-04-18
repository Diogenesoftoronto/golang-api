package main

import (
	"fmt"
// "net/http"
"encoding/json"
"os"
// "strconv"
"io/ioutil"
)

type ServerConfig struct {
	Greetings map[language]message `json:"greetings"`
	Users []User `json:"users"` 
}

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age int `json:"age"`
	Social string `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter string `json:"twitter"`
}

type greetings struct {
	Greetings []greeting `json:"greetings"`
}

type greeting struct {
	Language string `json:"language"`
	Message string `json:"message"`
}
func main() {
	// open and destructure the opened file into the jsonFIle and error value
	jsonFile, err := os.Open("server_config.json")
	// if the there error is present  print an error message
	if err != nil {
		fmt.Printf("Error opening server_config.json")
	}
	// other wise 
	fmt.Println("server_config.json successfully opened")

	defer jsonFile.Close()

	byteJson, _ := ioutil.ReadAll(jsonFile)
	// need to initialize the value that will store the json file to be unmarshalled
	var serverConfig ServerConfig

	json.Unmarshal(byteJson, &serverConfig)
	// func greeting(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, greeting)
	// }

	fmt.Printf("hell is here: %v", serverConfig.Greetings.)
}