package main

import (
	"fmt"
	"purr/day4"
	"purr/day5"
)
func main(){
	day4.CheckMap()
	samriddhi := day5.College{
		Name:"Samriddhi",
		Address:"Bhaktapur",
		Students:[]string{"ram","shyam","hari"},TotalNoOfStudent:30,
	}
	day5.DisplayCollegeInfo(samriddhi)
	samriddhi.EmbededDisplayCollegeInfo()

	circle := day5.Rectangle{
		Radius:30
	}
	rect := day5.Rectangle{
		Length: 500,
		Breadth: 34,
	}
	fmt.Println("Circle info....")
	day5.DisplayShapeInfo(circle)
	fmt.Println("Rectangle info....")
	day5.DisplayShapeInfo(rect)
}

