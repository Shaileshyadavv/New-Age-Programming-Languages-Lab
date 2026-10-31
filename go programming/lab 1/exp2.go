package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func getInt(sc *bufio.Scanner, msg string) int {
    for {
        fmt.Print(msg)
        sc.Scan()

        num, err := strconv.Atoi(sc.Text())

        if err == nil && num >= -100000 && num <= 100000 {
            return num
        }

        fmt.Println("Invalid input! Enter an integer between -100000 and 100000.")
    }
}

func getFloat(sc *bufio.Scanner, msg string) float64 {
    for {
        fmt.Print(msg)
        sc.Scan()

        num, err := strconv.ParseFloat(sc.Text(), 64)

        if err == nil && num >= -100000 && num <= 100000 {
            return num
        }

        fmt.Println("Invalid input! Enter a valid number between -100000 and 100000.")
    }
}

func main() {
    sc := bufio.NewScanner(os.Stdin)

    fmt.Println("----- INTEGER CALCULATOR -----")

    a := getInt(sc, "Enter first integer: ")
    b := getInt(sc, "Enter second integer: ")

    fmt.Println("Addition:", a+b)
    fmt.Println("Subtraction:", a-b)
    fmt.Println("Multiplication:", a*b)

    fmt.Println("\n----- FLOAT CALCULATOR -----")

    x := getFloat(sc, "Enter first float number: ")
    y := getFloat(sc, "Enter second float number: ")

    fmt.Println("Addition:", x+y)
    fmt.Println("Subtraction:", x-y)
    fmt.Println("Multiplication:", x*y)

    fmt.Println("\nCalculator completed successfully!")
}
