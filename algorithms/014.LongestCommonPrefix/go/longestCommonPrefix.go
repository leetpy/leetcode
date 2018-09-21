package main

import "fmt"

func longestCommonPrefix(strs []string) string {
        if len(strs) == 0 {
                return ""
        }

        firstStr := strs[0]
        if firstStr == "" {
                return ""
        }

        var found bool

        for i := 0; i < len(firstStr); i++ {
                for j := 1; j < len(strs); j++ {
                        curStr := strs[j]
                        if i >= len(curStr) || curStr[i] != firstStr[i] {
                                found = true
                                break
                        }
                }
                if found {
                        return strs[0][:i]
                }
        }
        return firstStr

}

func main() {
        strs := [][]string{
                {"flower", "flow", "flight"},
                {"flower", "f", "flight"},
                {"b", ""},
                {"", "b"},
                {"a", "ac"},
                {"aaa", "aa", "aaa"},
                {"flower", "flow", "flight"},
                {"aa", "a"},
                {"aa", "aabc", "a"},
                {"aa", "aabc", "ac"},
        }
        for _, s := range strs {
                fmt.Println(longestCommonPrefix(s))
        }
}
