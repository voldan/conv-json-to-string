package main

import (
    "os"
    "log"
    "path/filepath"
    "io/ioutil"
    "fmt"
    "strings"
)

func isJsonFile(fileName string) bool {
    if filepath.Ext(fileName) == ".json" {
        return true
    }
    return false
}

func main()  {
    fileName := os.Args[1]
    if !isJsonFile(fileName) {
        log.Fatal("Error: input file extension is not a json")
    }

    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }
    contentStr := string(content)
    contentStr = strings.Replace(contentStr, "\n", "", -1)
    contentStr = strings.Replace(contentStr, " ", "", -1)
    fmt.Println(contentStr)
}
