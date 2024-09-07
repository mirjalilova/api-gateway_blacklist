package main

import (
    "os"
    "fmt"
    "github.com/pdfcrowd/pdfcrowd-go"
    "io/ioutil"
)

func readFile(fileName string) []byte {
    content, err := ioutil.ReadFile(fileName)
    handleError(err)
    return content
}

func main() {
    client := pdfcrowd.NewPdfToTextClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

    outputStream, err := os.Create("invoice.txt")

    handleError(err)

    defer outputStream.Close()

    err = client.ConvertRawDataToStream(readFile("/path/to/hello_world.pdf"), outputStream)

    handleError(err)
}

func handleError(err error) {
    if err != nil {
        why, ok := err.(pdfcrowd.Error)
        if ok {
            os.Stderr.WriteString(fmt.Sprintf("Pdfcrowd Error: %s\n", why))
        } else {
            os.Stderr.WriteString(fmt.Sprintf("Generic Error: %s\n", err))
        }

        panic(err.Error())
    }
}