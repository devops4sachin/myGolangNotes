package vis

import "fmt" // If you import fmt, it is not visible to other files in the same package.

// PrintVar is exported because it starts with a capital letter
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}
