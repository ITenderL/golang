package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	// 对象转json
	movie := Movie{"喜剧之王", 2000, 50, []string{"周星驰", "张柏芝"}}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json Marshal 转换失败")
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)
	// json转对象
	movie1 := Movie{}
	err = json.Unmarshal(jsonStr, &movie1)
	if err != nil {
		fmt.Println("json unmarshal err ", err)
		return
	}
	fmt.Printf("%v\n", movie1)
}
