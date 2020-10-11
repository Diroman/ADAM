package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func readImage() []byte {
	imgFile, err := os.Open("data/bmw5.jpeg") // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	return buf
}

func main() {
	url := "http://192.168.31.44:8081/recognition"

	photo := readImage()

	encoded := base64.StdEncoding.EncodeToString(photo)
	pld := fmt.Sprintf("{\"content\":\"%s\"}", encoded)
	payload := bytes.NewBuffer([]byte(pld))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal()
	}
	defer res.Body.Close()

	fmt.Println(string(body))
}