package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digit := []int{}
	for {
		if x < 10 {
			digit = append(digit, x)
			break
		}
		d := x % 10
		x /= 10
		digit = append(digit, d)
	}

	for i, v := range digit {
		if v != digit[len(digit)-i-1] {
			return false
		}
	}
	return true
}

// 翻转方式
func isPalindrome2(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) { 
        return false
    }   

    // 已经翻转的数字
    revertedNumber := 0
    for {
        if x <= revertedNumber {  // 控制只翻转一半
            break
        }
        revertedNumber = revertedNumber*10 + x%10
        x /= 10
    }                                                                                                                                                                         
    return x == revertedNumber || x == revertedNumber/10 // 考虑奇偶
}

func main() {
	fmt.Println(isPalindrome(121))
}
