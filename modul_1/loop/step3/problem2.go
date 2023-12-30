package main

import (
    "fmt"
)

func main(){
    var a, b, summ int
    fmt.Scan(&a, &b)
    
    for i := a; i <= b; i++ {
        summ += i
    }
    fmt.Println(summ)
}