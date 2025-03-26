package main
import ("fmt")

func Calculate(condition string ,numbers ...int,) int{
  if condition == "multiply" {
	fmt.Printf("array:%d\n",numbers)
	mul:=1
	for _, number:=range numbers {
	   mul*=number
	}
	return mul
  }else if condition == "add"{
    add:=0
	for _, number:=range numbers {
	   add+=number
	}
	return add
  }else{
	sub:=0
	for _, number:=range numbers {
	   sub-=number
	}
	return sub
  }
}