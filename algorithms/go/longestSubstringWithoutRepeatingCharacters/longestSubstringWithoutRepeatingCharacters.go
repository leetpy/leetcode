package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    index := [128]int{}
    ans := 0
    i := 0
    for j := 0; j < len(s); j++ {
        fmt.Println(s[j])
        i = max(i, index[s[j]])
        ans = max(ans, j-i+1)
        index[s[j]] = j + 1
    }
    fmt.Println(index)
    return ans
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func main() {
    s := "abcabcbb"
    fmt.Println(lengthOfLongestSubstring(s))
}
