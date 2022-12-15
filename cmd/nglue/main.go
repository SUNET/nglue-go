package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Data struct {
	Hostname           string `json:"hostname"`
	Description        string `json:"description"`
	Servicestateid     int64  `json:"servicestateid"`
	Lastproblemid      string `json:"lastproblemid"`
	Problemid          string `json:"problemid"`
	Lastservicestateid int64  `json:"lastservicestateid"`
	Notification       string `json:"notification"`
	Servicestate       string `json:"servicestate"`
	Attempt_number     int    `json:"attempt_number"`
	Max_attempts       int    `json:"max_attempts"`
	Test_api           bool   `json:"test_api"`
	Debug              bool   `json:"debug"`
	Sync               bool   `json:"sync"`
	Validate           bool   `json:"validate"`
}

func main() {

	hostname := flag.String("hostname", "", "Hostname")
	description := flag.String("description", "", "Description")
	servicestateid := flag.Int64("servicestateid", 0, "Service state id")

	lastproblemid := flag.String("lastproblemid", "", "lastproblemid")
	problemid := flag.String("problemid", "", "problemid")
	lastservicestateid := flag.Int64("lastservicestateid", 0, "Service state id")

	notification := flag.String("notification", "", "notification")
	servicestate := flag.String("servicestate", "", "servicestate")

	attempt_number := flag.Int("attempt_number", 0, "attempt number")
	max_attempts := flag.Int("max_attempts", 0, "max attempts")
	test_api := flag.Bool("test_api", true, "test api")
	debug := flag.Bool("debug", true, "test api")
	sync := flag.Bool("sync", true, "test api")
	validate := flag.Bool("validate", true, "test api")

	flag.Parse()

	data := &Data{Hostname: *hostname, Description: *description, Servicestateid: *servicestateid, Lastproblemid: *lastproblemid, Problemid: *problemid, Lastservicestateid: *lastservicestateid, Notification: *notification, Servicestate: *servicestate, Attempt_number: *attempt_number,
		Max_attempts: *max_attempts, Test_api: *test_api, Debug: *debug, Sync: *sync, Validate: *validate}
	// Let us see if we get a server address from ENV
	url := os.Getenv("NGLUE_SERVER")
	if url == "" {
		// This is our default location
		url = "http://localhost:8000/"
	}
	// Convert the input into JSON
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	//request method
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	// set HTTP request header Content-Type
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		fmt.Println(err)
		return
	}
	// Doing the actual request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Reading back the output from the server
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// If there is any output, we are just printing it.
	// TODO: Here we can enable logic to verify if the previous calls were successful or not.
	fmt.Println(string(body))
}
