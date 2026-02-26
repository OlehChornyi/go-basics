package main

import "fmt"

func main() {
	fmt.Println("Maps things")
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
}
