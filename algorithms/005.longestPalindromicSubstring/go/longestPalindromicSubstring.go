package main

import "fmt"

func longestPalindrome(s string) string {
    if len(s) < 1 {
        return ""
    }
    start := 0
    end := 0
    for i, _ := range s {
        l1 := expandAroundCenter(s, i, i)
        l2 := expandAroundCenter(s, i, i+1)
        l := max(l1, l2)
        if l > (end - start) {
            start = i - (l-1)/2
            end = i + l/2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) int {
    for {
        if left >=0 && right < len(s) && s[left] == s[right] {
            left--
            right++
        } else {
            break
        }
    }
    return right - left - 1
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func main() {
    fmt.Println(longestPalindrome("aba"))
}
