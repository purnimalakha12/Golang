package day5
type Shape interface{
	Area() int
	Perimeter() int
}

type Rectangle struct{
	Length int
	Breadth int
}

func (r Rectangle)Area() int{
	return r.Length*r.Breadth
}
func (r Rectangle)Perimeter() int{
	return (r.Length+r.Breadth)
}

type Circle struct{
	Radius float32
}

func (c Circle) Area() {
	return 3* c.Radius* c.Radius
}
func (c Circle) Perimeter() {
	return 2*3* c.Radius* 
}

func DisplayShapeInfo(s Shape){
	fmt.Println("Area:",s.Area())
	fmt.Println("Perimeter:",s.Perimeter())
}
