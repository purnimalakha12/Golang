package day5

import "fmt"

//defining struct
type College struct{
	Name string
Address string
Students []string
TotalNoOfStudent int
}

func DisplayCollegeInfo (c College){
	fmt.Println("Name:",c.Name)
	fmt.Println("Address:",c.Address)
	fmt.Println("Students:",c.Students)
	fmt.Println("Total no of students:",c.TotalNoOfStudent)
}

func (c College)EmbededDisplayCollegeInfo(){
	fmt.Println("E.Name:",c.Name)
	fmt.Println("E.Address:",c.Address)
	fmt.Println("E.Students:",c.Students)
	fmt.Println("E.Total no of students:",c.TotalNoOfStudent)
}