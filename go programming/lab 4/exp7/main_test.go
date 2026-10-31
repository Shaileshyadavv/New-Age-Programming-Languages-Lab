
package main

import "testing"

func TestDivide(t *testing.T) {
    result, err := divide(10, 2)

    if err != nil {
        t.Fatal(err)
    }

    if result != 5 {
        t.Errorf("Expected 5, got %f", result)
    }
}

func TestDivideByZero(t *testing.T) {
    _, err := divide(10, 0)

    if err == nil {
        t.Error("Expected division by zero error")
    }
}

func TestSquareRootNegative(t *testing.T) {
    _, err := squareRoot(-9)

    if err == nil {
        t.Error("Expected negative square root error")
    }
}

func TestSquareRoot(t *testing.T) {
    result, err := squareRoot(25)

    if err != nil {
        t.Fatal(err)
    }

    if result != 5 {
        t.Errorf("Expected 5, got %f", result)
    }
}

func TestEmptyInput(t *testing.T) {
    _, err := checkInput("")

    if err == nil {
        t.Error("Expected empty input error")
    }
}

func TestValidInput(t *testing.T) {
    result, err := checkInput("Hello")

    if err != nil {
        t.Fatal(err)
    }

    if result != "Hello" {
        t.Errorf("Expected Hello, got %s", result)
    }
}
