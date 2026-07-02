type Solution struct{}

// strings.join is where you can actually do concatenation to all elements in an array
// if you do like strings.join, result := strings.Join(strs, " "), the 2nd element is separator element
// notes: need to use the double quote than the single quote for referring to a separator character
// result: go is awesome
// another example
// result := strings.join(strs, ","), 
// result: go,is,awesome 

func (s *Solution) Encode(strs []string) string {
    if len(strs) == 0{
        return ""
    }

    var encoded_num []string

    for i := 0 ; i < len(strs) ; i++{
        encoded_num = append(encoded_num, strconv.Itoa(len(strs[i])))
    }

    return strings.Join(encoded_num, ",") + "#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
    if len(encoded) == 0{
        return []string{}
    }
    // first split with #, use splitN, for specific spliting, in here, 2 parts
    first_split := strings.SplitN(encoded, "#", 2)
    // second split with ',', we take the left part, and then parse per every length of the substring
    second_split := strings.Split(first_split[0], ",")

    // start index for slicing inside
    i := 0
    var res []string

    for _, substring := range second_split {
        if substring == ""{
            continue
        }
        // if using comma, the comma value will be ignored, that's why the second value kept unused, 
        // because we don't want to detect any errors
        length, _ := strconv.Atoi(substring)
        // take the right argument
        res = append(res, first_split[1][i:i+length])
        i += length
    }

    return res
}
