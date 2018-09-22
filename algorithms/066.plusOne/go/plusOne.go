package main

func plusOne(digits []int) []int {
    result := []int{}
    carries := 1
    for i := len(digits)-1; i >=0; i-- {
        sum := digits[i] + carries
        digits[i] = sum % 10
        carries = sum / 10

        if carries == 0 {
            return digits
        }
    }
    result = append(result, carries)
    for _, v := range digits {
        result = append(result, v)
    }
    return result
}
