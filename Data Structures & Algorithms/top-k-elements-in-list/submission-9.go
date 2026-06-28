func topKFrequent(nums []int, k int) []int {
    amap := make(map[int]int)
    for _, n := range nums{
        amap[n]++
    }

    // determine the key and value, then sort the value
    type kv struct {
        key int
        value int
    }

    // temp array for storing the dictionary
    var ss []kv

    // to put the dictionary inside an array
    for key, v := range amap{
        ss = append(ss, kv{key, v})
    }
    
    // arrow functions look alike
    sort.Slice(ss, func(i, j int) bool {
        return ss[i].value > ss[j].value
    })

    // make another temp array with the length of k input
    result := make([]int, k)
    for i := 0 ; i < k ; i++ {
        result[i] = ss[i].key
    }

    return result

}
