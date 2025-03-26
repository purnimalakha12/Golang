package main

import (
	"fmt"
	"log"

)


func main() {
	res,length:=adds(10,5,9,13)
	log.Printf("result=%d",res)
	log.Printf("length=%d",length)
//Anonymous&Closure Function
	func() {
        fmt.Println("This is an anonymous function")
    }()
	display := func(name string){
		fmt.Printf("name=%s\n",name)
	}
	display("John")
	
	next:=sequenceGenerator()
	a:=next()
	fmt.Printf("%d\n",a)
	b:=next()
	fmt.Printf("%d",b)
	for i:=0;i<10;i++{
		res:=next()
		fmt.Printf("%d\n",res)
	}
	cres := Calculate("add",49,20,25)
	fmt.Println(cres)
}
func adds(numbers ...int)(int,int){
 fmt.Printf("array:%d\n",numbers)
 sum:=0
 for _, number:=range numbers {
	sum+=number
 }
 return sum, len(numbers)
}

/* Closure Function 
Capture & reatin*/
func sequenceGenerator() func () int{
 i:=0
 return func() int {
	i++ 
	return i
 }
}
