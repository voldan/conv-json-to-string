package main

import (
    "os"
    "log"
    "path/filepath"
    "io/ioutil"
    "fmt"
    "strings"
    "encoding/json"
)

func isJSONFile(fileName string) bool {
    if filepath.Ext(fileName) == ".json" {
        return true
    }
    return false
}

func validateJSONObject(content []byte) error {
    var j map[string]interface{}
    err := json.Unmarshal(content, &j)
    if err != nil {
        return err
    }
    return nil
}

func main()  {
    fileName := os.Args[1]
    if !isJSONFile(fileName) {
        log.Fatal("Error: input file extension is not json")
    }

    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    if err = validateJSONObject(content); err != nil {
        log.Fatalf("Invalid json schema file '%s', error: %v", fileName, err)
    }

    contentStr := string(content)
    contentStr = strings.Replace(contentStr, "\n", "", -1)
    contentStr = strings.Replace(contentStr, " ", "", -1)
    fmt.Println(contentStr)
}
