package main

import "fmt"

func main() {

	printHeader()
	operation, err := readOperation()

	if err != nil {
		panic(err)
	}
	num1, num2, _ := readNumbers()

	result, err := calculate(operation, num1, num2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The result of %d %s %d is %d\n", num1, operation, num2, result)
}

func printHeader() {
	fmt.Println("")
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println("* Calculator in Go 1.0 *")
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println("")
}

func calculate(operation string, a, b int16) (result int16, err error) {
	switch operation {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		result = a / b
	}

	return
}

func readOperation() (operation string, err error) {
	fmt.Print("Enter the operation to perform (add, sub, mul, div): ")
	fmt.Scanf("%s", &operation)

	containsOperation := containsOperation(operation)

	if !containsOperation {
		err = fmt.Errorf("invalid operation")
		return
	}

	return
}

func readNumbers() (a, b int16, err error) {
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &b)

	return
}

var operations = []string{"add", "sub", "mul", "div"}

func containsOperation(operation string) bool {
	for _, o := range operations {
		if o == operation {
			return true
		}
	}

	return false
}
