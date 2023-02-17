package main

import "fmt"

func main() {

	result := []map[string]interface{}{{
		"one": 22,
		"two": 444,
	}, {
		"one": map[string]interface{}{"dana": 9000, "danaku": []map[string]interface{}{
			{"disini": "okeeee"},
			{"disini": []map[string]interface{}{{"dimana": "not ok"},
				{"darimana": "asiaap"}, {"daribumi": []map[string]interface{}{{"akulah": "arjuna"}, {"yang": "mencari cinta"}}}}}},
			"two": 6666,
		}}}

	fmt.Println(result[1]["one"].(map[string]interface{})["danaku"].([]map[string]interface{})[1]["disini"].([]map[string]interface{})[2]["daribumi"].([]map[string]interface{})[1]["yang"])
}
