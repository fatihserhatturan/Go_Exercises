// main.go
package main
import (
    "fmt"
    "mathutils/mathutils.go"
)

func main() {
    fmt.Println("Sum:", mathutils.Add(5, 3))
    fmt.Println("Diff:", mathutils.Sub(10, 4))
    fmt.Println("Mul:", mathutils.Mul(6, 7))
    fmt.Println("Div:", mathutils.Div(10, 2))
}
