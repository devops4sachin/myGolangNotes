package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname string, lname string) (s string) { // We are returning "s" here, this is called return naming
	s = fmt.Sprint(fname, lname)
	return // need not to specify "s" here
}

/*
IMPORTANT
Avoid using named returns.


Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/
