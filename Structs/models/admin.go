package models

type Person struct {
	Name string
	Age  int
}

type Admin struct {
	User
	Person
	CanDeleteStuff bool
}

func main() {
	adm := Admin{
		CanDeleteStuff: true,
	}

	adm.Person.Name = "John Doe"
}
