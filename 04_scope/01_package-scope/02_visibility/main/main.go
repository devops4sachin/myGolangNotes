package main

import (
	"fmt"

	"github.com/devops4sachin/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName) // We user the "vis.MyName" from package "vis". "MyName" is exported from the "vis" package.
	vis.PrintVar()
}
