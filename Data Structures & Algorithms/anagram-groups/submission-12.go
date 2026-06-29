// n log n solution actually
func groupAnagrams(strs []string) [][]string {
    // this creates e.g "hi": ["what", "why", "how"]
    groups := make(map[string][]string)

    for _, value := range strs{
        // needs to be sorted so that the index key value can be passed to that array of same key
        sorted_value_for_key := sort_for_key(value)
        groups[sorted_value_for_key] = append(groups[sorted_value_for_key], value)
    }

    var result [][]string
    for _, group := range groups {
        result = append(result, group)
    }
    return result
    
}


func sort_for_key(str string) string{
    // the rune is used than bytes said it's for unicode replacement (?), need to look it up
    characters := []rune(str) // split to every character, i guess
    sort.Slice(characters, func(i, j int)bool{
        return characters[i] < characters[j]
    })
    return string(characters)
}
