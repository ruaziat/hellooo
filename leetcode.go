package main
import (
    "fmt"
"sort"
)

func main() {
var a = []int{7, 2}
var b = []int{3, 1}
var c = append(a, b...)
    sort. Ints ( c ) 
    fmt . Println ( "Cортированный массив:" ,  c )
}

