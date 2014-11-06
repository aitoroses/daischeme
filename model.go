package main


type MyModel struct {
	FirstName string 			`json:"firstName"`
	Name string 			`json:"name"`
	Salary float64 			`json:"salary"`
	Nested Nested 			`json:"nested"`

}

type Nested struct {
	Val1 float64 			`json:"val1"`
	Val2 bool 			`json:"val2"`

}
