func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    char_counter1 := make(map[byte]int)
    char_counter2 := make(map[byte]int)

    for i := range s{
        char_counter1[s[i]]++
        char_counter2[t[i]]++
    }

    // key, value
    for char, count := range char_counter1{
        if char_counter2[char] != count{
            return false
        }
    }

    return true
}
