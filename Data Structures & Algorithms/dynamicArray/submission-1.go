type DynamicArray struct {
    capacity int
    array []int
    size int
}

func NewDynamicArray(capacity int) *DynamicArray {
    if capacity <= 0 {
        return nil;
    }

    da := &DynamicArray{
        capacity: capacity,
        size: 0,
        array: make([]int, 0, capacity),
    }
    return da
}

func (da *DynamicArray) Get(i int) int {
    if i < 0 || i >= da.size {
        return 0
    }
    return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
      if i < 0 || i >= da.size {
        return
    }
    da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.size == da.capacity {
        da.resize()
    }
    da.array = append(da.array, n);    
    da.size++;
}

func (da *DynamicArray) Popback() int {
    da.size--
    tempnum := da.array[da.size]
    return tempnum
}

func (da *DynamicArray) resize() {
    da.capacity *= 2;
    newArray := make([]int, da.size, da.capacity)
    copy(newArray, da.array)
    da.array = newArray
}

func (da *DynamicArray) GetSize() int {
    return da.size
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
