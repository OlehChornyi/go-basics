package main

import "fmt"

func main() {
	//1st - length, 2nd - capacity
	userNames := make([]string, 2, 5)

	userNames[0] = "Serhiy"
	userNames[1] = "Yevhen"

	userNames = append(userNames, "Oleh")
	userNames = append(userNames, "Ihor")

	fmt.Println(userNames)

	cources := make(map[string]float64, 3)

	cources["go"] = 4.7
	cources["flutter"] = 4.8

	fmt.Println(cources)
}

// type Websites struct {
// 	google string
// 	aws    string
// }

// func main() {
// 	fmt.Println("Maps things")
// 	websites := map[string]string{
// 		"Google":              "https://google.com",
// 		"Amazon Web Services": "https://aws.com",
// 	}

// 	fmt.Println(websites)
// 	fmt.Println(websites["Amazon Web Services"])

// 	websites["LinkedIn"] = "https://linkedin.com"
// 	fmt.Println(websites)

// 	delete(websites, "Google")
// 	fmt.Println(websites)
// }
