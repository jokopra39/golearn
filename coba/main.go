package main

import "fmt"

func main() {

	result := []map[string]interface{}{{
		"one": 22,
		"two": 444,
	}, {
		"one": map[string]interface{}{"dana": 9000, "danaku": "oop"},
		"two": 6666,
	}}

	fmt.Println(result[1]["one"].(map[string]interface{})["danaku"].(string))
}
