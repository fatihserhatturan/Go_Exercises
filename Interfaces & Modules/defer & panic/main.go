package main
import (
    "fmt"
    "os"
)

func saveFile(name string, content string) {
    file, err := os.Create(name)
    if err != nil {
        panic("failed to create file")
    }
    defer file.Close()
    _, err = file.WriteString(content)
    if err != nil {
        panic("failed to write to file")
    }
    fmt.Println("file saved successfully")
}

func riskyDivision(a, b int) int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

func demo() {
    defer fmt.Println("demo finished")
    fmt.Println("result 1:", riskyDivision(10, 2))
    fmt.Println("result 2:", riskyDivision(5, 0))
    fmt.Println("this line will run after recovery")
}

func main() {
    defer fmt.Println("program exited")
    fmt.Println("starting program")
    demo()
    saveFile("output.txt", "This is a test file")
    fmt.Println("end of main")
}
