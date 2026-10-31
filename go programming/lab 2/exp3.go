package main
import "fmt"

func main() {

    fmt.Println("----- SLICE OPERATIONS -----")

    students := []string{"Rahul", "Aman", "Priya"}

    fmt.Println("Original Slice:", students)

    students = append(students, "Shailesh")
    fmt.Println("After Adding:", students)

    index := 1
    students = append(students[:index], students[index+1:]...)
    fmt.Println("After Removing:", students)

    students[1] = "Rohit"
    fmt.Println("After Updating:", students)


    fmt.Println("\n----- MAP OPERATIONS -----")

    marks := map[string]int{
        "Go": 85,
        "Java": 90,
        "Python": 80,
    }

    fmt.Println("Original Map:", marks)

    marks["C++"] = 95
    fmt.Println("After Inserting:", marks)

    delete(marks, "Java")
    fmt.Println("After Deleting:", marks)

    subject := "Go"
    if value, found := marks[subject]; found {
        fmt.Println("Marks in", subject, ":", value)
    } else {
        fmt.Println("Subject not found")
    }

    fmt.Println("\nAll operations completed!")
}
