package day4

import (
	"fmt"
)

func CheckMap() {
	person := map[string]interface{}{
		"name":    "ram",
		"age":     45,
		"address": "kathmandu",
	}
	fmt.Println("ram:", person, "length:", len(person))
	for key, value := range person {
		fmt.Println("key:", key, "value:", value)
	}
	student := make(map[string]any)
	student["name"] = "Ramesh"
	student["age"] = 56
	student["class"] = "10th"
	fmt.Println("student:", student)
	book := map[string]any{
		"book1": "book1",
		"book2": "book2",
		"book3": "book3",
	}
	mergedMap := make(map[string]any)
	maps.Copy(mergedMap, book)
	maps.Copy(mergedMap, student)
	fmt.Println("Merged Map:", mergedMap)
}
