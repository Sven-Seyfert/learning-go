package main

import (
	"03-scope/packageone"
	"fmt"
)

var packageLevelVar = "This is a package level variable, only for the \"main\" package."

func main() {
	blockLevelVar := "This is a block level variable for the \"main\" func."

	fmt.Println(packageLevelVar, blockLevelVar)

	anotherFunction()

	fmt.Println("")

	fmt.Println(packageone.PackageLevelVar)
	fmt.Println(packageone.PublicFuncOfPackageone())
}

func anotherFunction() {
	blockLevelVar := "This is a block level variable for the \"anotherFunction\" func."

	fmt.Println(blockLevelVar)
}
