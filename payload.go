Copypackage main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	// Define command line arguments
	filepath := flag.String("o", "", "Output file path")
	timestamp := flag.Bool("t", false, "Include timestamp in file name")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("No payload data provided.")
	}

	payloadString := strings.Join(args, " ")
	payloadBytes := []byte(payloadString)

	if bytes.HasPrefix(payloadBytes, []byte("{")) || bytes.HasPrefix(payloadBytes, []byte("[")) {
		// JSON payload
		var payloadJSON interface{}
		err := json.Unmarshal(payloadBytes, &payloadJSON)
		if err != nil {
			log.Fatalf("Failed to decode JSON payload: %v", err)
		}

		// Output to console
		fmt.Printf("%+v\n", payloadJSON)

		// Output to file
		outputPayload(payloadBytes, *filepath, "json", *timestamp)

	} else if bytes.HasPrefix(payloadBytes, []byte("<")) {
		// XML payload
		var payloadXML interface{}
		err := xml.Unmarshal(payloadBytes, &payloadXML)
		if err != nil {
			log.Fatalf("Failed to decode XML payload: %v", err)
		}

		// Output to console
		fmt.Printf("%+v\n", payloadXML)

		// Output to file
		outputPayload(payloadBytes, *filepath, "xml", *timestamp)

	} else {
		// YAML payload
		var payloadYAML interface{}
		err := yaml.Unmarshal(payloadBytes, &payloadYAML)
		if err != nil {
			log.Fatalf("Failed to decode YAML payload: %v", err)
		}

		// Output to console
		fmt.Printf("%+v\n", payloadYAML)

		// Output to file
		outputPayload(payloadBytes, *filepath, "yaml", *timestamp)
	}
}

func outputPayload(payload []byte, filepath string, format string, timestamp bool) {
	if filepath == "" {
		return
	}

	// Add timestamp to filename if specified
	if timestamp {
		timestamp := fmt.Sprintf("%d", Now().Unix())
		filename := fmt.Sprintf("payload-%s.%s", timestamp, format)
		filepath = filepath + "/" + filename
	}

	// Write payload to file
	err := ioutil.WriteFile(filepath, payload, 0644)
	if err != nil {
		log.Fatalf("Failed to write payload to file: %v", err)
	} else {
		log.Printf("Payload written to file: %s", filepath)
	}
}
