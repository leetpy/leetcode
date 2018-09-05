package main

import "fmt"

func isMonotonic(A []int) bool {
    if len(A) <= 1 {
        return true
    }
    
    isInc := true
    isDes := true
    for i := 1; i < len(A); i++ {
        if A[i] < A[i-1] {
            isInc = false
        }
        if A[i] > A[i-1] {
            isDes = false
        }
    }
    return isInc || isDes
}

func main() {
    a := []int{3, 2, 2, 1}
    fmt.Println(isMonotonic(a))
}
