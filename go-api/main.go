package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// url := os.Args[1]
	url := "https://httpbin.org/ip"
	api(url)

	// jsonData := map[string]string{"firstName": "Paranie", "lastName": "Paranie", "email": "paranie@gamil.com"}
	// jsonValue, err := json.Marshal(jsonData)
	// if err != nil {
	// 	log.Fatalf("Error marshaling JSON: %v", err)
	// }
	// response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	// if err != nil {
	// 	log.Fatalf("Error sending POST request: %v", err)
	// }
	// defer response.Body.Close()
	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response body: %v", err)
	// }

	// fmt.Println("POST Response:")
	// fmt.Println(string(data))

	postUrl := "https://httpbin.org/post"
	data := `{"name": "Parnaietharan", "age": 23, "email":"paranietharan64@gmail.com"}`
	apiPost(postUrl, data)

	// direct
	// jsonData := map[string]string{"firstName": "Paranie", "lastName": "Paranie", "email": "paranie@gamil.com"}
	// jsonValue, _ := json.Marshal(jsonData)
	// response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	// if err != nil {
	// 	fmt.Printf("Error : %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// }
}

func api(url string) {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func apiPost(url string, jsonData string) {
	jsonBody := []byte(jsonData)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:")
	fmt.Println(string(body))
}
