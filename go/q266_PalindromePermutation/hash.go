package q266_PermutePalindrome

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)

    for _, c := range s {
        m[c]++
    }

    odd := 0
    for _, v := range m {
        if v%2 != 0 {
            odd++
        }
    }

    return odd == 0 || odd == 1
}