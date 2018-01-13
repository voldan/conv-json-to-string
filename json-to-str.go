package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func isJSONFile(fileName string) bool {
	if filepath.Ext(fileName) == ".json" {
		return true
	}
	return false
}

func validateJSON(content []byte, jsonSchema bool) error {
	var jObj map[string]interface{}
	var j interface{}

	if jsonSchema {
		err := json.Unmarshal(content, &jObj)
		if err != nil {
			return err
		}
	} else {
		err := json.Unmarshal(content, &j)
		if err != nil {
			return err
		}
	}
	return nil
}

func escapeStr(content string) string {
	content = strings.Replace(content, "\\", "\\\\", -1)
	content = strings.Replace(content, "\"", "\\\"", -1)
	return content
}

func main() {
	schemaPtr := flag.Bool("s", false, "json file is a json schema")
	escapedStringPtr := flag.Bool("e", false, "convert json to the escape string")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Error: no input file listed")
	}

	fileName := args[0]
	if !isJSONFile(fileName) {
		log.Fatalf("Error: extension of the input file '%s' is not json", fileName)
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if err = validateJSON(content, *schemaPtr); err != nil {
		log.Fatalf("Invalid json file '%s', error: %v", fileName, err)
	}

	contentStr := string(content)
	contentStr = strings.Replace(contentStr, "\n", "", -1)
	contentStr = strings.Replace(contentStr, " ", "", -1)
	if *escapedStringPtr {
		contentStr = escapeStr(contentStr)
	}
	fmt.Println(contentStr)
}
