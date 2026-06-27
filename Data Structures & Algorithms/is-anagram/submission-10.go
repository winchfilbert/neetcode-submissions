// also sort can make this easier, if being asked by follow ups
// best sort algorithms can at least do 0(n log n)

func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

       // it doesnt mean that although you know mapping the dictionaries of string to int,
       /*  you make it to be like this
        s_hmap := make(map[string]int)
        h_hmap := make(map[string]int)

        instead, you use bytes (characters) to their counts
       */

        s_hmap := make(map[byte]int)
        t_hmap := make(map[byte]int)


    for i := 0; i < len(s); i++ {
        // i tried to make hashmaps by using the make keyword, to allocate the dictionary
        // In Go, if a key doesn't exist, it defaults to 0. 
        // So we can just do map[key]++ safely!
        s_hmap[s[i]]++
        t_hmap[t[i]]++
    }

    for char, count := range s_hmap{
        if t_hmap[char] != count{
            return false
        }
    }

    return true
}
