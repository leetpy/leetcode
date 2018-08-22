package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for k, v := range nums {
        if _, ok := m[v]; ok {
            return []int{m[v], k}
        } else {
            m[target-v] = k
        }
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    data := twoSum(nums, target)
    fmt.Println(data)
}
