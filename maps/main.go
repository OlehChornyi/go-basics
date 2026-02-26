package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	//1st - length, 2nd - capacity
	userNames := make([]string, 2, 5)

	userNames[0] = "Serhiy"
	userNames[1] = "Yevhen"

	userNames = append(userNames, "Oleh")
	userNames = append(userNames, "Ihor")

	fmt.Println(userNames)

	cources := make(floatMap, 3)

	cources["go"] = 4.7
	cources["flutter"] = 4.8

	cources.output()

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range cources {
		fmt.Println(key)
		fmt.Println(value)
	}
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
