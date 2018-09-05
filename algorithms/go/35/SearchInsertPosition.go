package main

import "fmt"

func searchInsert(nums []int, target int) int {
    for k, v := range nums {
        if target <= v {
            return k
        }
    }
    return len(nums)
}

func main() {
  target := 5
  nums := []int{1, 3, 5, 6}
  fmt.Println(searchInsert(nums, target))
}
