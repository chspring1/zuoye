package main

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	person     Person
	EmployeeID string
}

func PrintInfo(e Employee) {
	println("Name:", e.person.Name)
	println("Age:", e.person.Age)
	println("Employee ID:", e.EmployeeID)

}

func main() {
	e := Employee{
		person: Person{
			Name: "Hou_wenjin",
			Age:  25,
		},
		EmployeeID: "9528",
	}
	PrintInfo(e)
}
