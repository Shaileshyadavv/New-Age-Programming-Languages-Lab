package main

import (
	"fmt"

	"custompkgdemo/mathutils"
	"custompkgdemo/stringutils"
)

func main() {
	fmt.Println("===== Custom Reusable Package Demo =====")

	word := "Golang Programming"
	fmt.Println("\n--- String Utilities ---")
	fmt.Printf("Original string : %s\n", word)
	fmt.Printf("Reversed string : %s\n", stringutils.Reverse(word))
	fmt.Printf("Vowel count     : %d\n", stringutils.CountVowels(word))

	num := 5
	base, exp := 2, 10
	fmt.Println("\n--- Math Utilities ---")
	fmt.Printf("Factorial of %d      : %d\n", num, mathutils.Factorial(num))
	fmt.Printf("%d raised to power %d : %d\n", base, exp, mathutils.Power(base, exp))

}
