// data_map
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Map!")

	var countryMap map[string]string

	countryMap = make(map[string]string)

	countryMap["china"] = "中国"
	for country := range countryMap {

		fmt.Println(country, "名称:", countryMap[country])
	}

	capital, ok := countryMap["china2"]

	if ok {
		fmt.Println("名称是:", capital)
	} else {
		fmt.Println("值不存在")
	}

	//delete
}
