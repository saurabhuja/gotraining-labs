package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println(string(data))
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	fmt.Print("Enter Number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aFloat)
	}

	// content := "Hello from Go!"
	file, err := os.Create("./filefromGo.txt")
	checkError(err)
	length, err := io.WriteString(file, input)
	checkError(err)
	fmt.Printf("Wrote file with %v characters.\n", length)
	defer file.Close()
	defer readFile(file.Name())

	url := "http://date.jsontest.com/"
	fmt.Println("Network requests:")
	response, err := http.Get(url)
	checkError(err)
	// fmt.Println(response)
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(bytes))
	dates := dateFromJson(bytes)
	for date := range dates {
		fmt.Println(dates[date].Date)
		fmt.Println(dates[date].Epoch)
		fmt.Println(dates[date].Time)
	}

}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type Datestruct struct {
	Date  string `json:"date"`
	Epoch int64  `json:"milliseconds_since_epoch"`
	Time  string `json:"time"`
}

func dateFromJson(content []byte) []Datestruct {
	dates := make([]Datestruct, 0, 20)

	//decoder way
	decoder := json.NewDecoder(bytes.NewReader(content))
	for {
		var date Datestruct
		err := decoder.Decode(&date)
		if err == io.EOF {
			break
		}
		checkError(err)
		dates = append(dates, date)
	}
	// json unmarshal way
	var date Datestruct
	json.Unmarshal(content, &date)
	dates = append(dates, date)
	return dates
}
