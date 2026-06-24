class DynamicArray {
    public int[] arr;
    public int capacity;
    public int size;
    
    public DynamicArray(int capacity) {
        this.capacity = capacity;
        this.arr = (int[]) new int[capacity];
        this.size = 0;
    }

    public int get(int i) {
        if (i < 0 || i > size){
            throw new IndexOutOfBoundsException();
        }
        return arr[i];
    }

    public void set(int i, int n) {
         if (i < 0 || i > size){
            throw new IndexOutOfBoundsException();
        }
        arr[i] = n;
    }

    public void pushback(int n) {
        if(capacity == size){
            resize();
        }

        arr[size] = n;
        size++;
    }

    public int popback() {
        size--;
        int poppedNumber = arr[size];
        arr[size] = 0;
        return poppedNumber;
    }  

    private void resize() {
        this.capacity *= 2;
        int[] newarr = new int[capacity];
        for(int i = 0 ; i < size; i++){
            newarr[i] = arr[i];
        }
        this.arr = newarr;
    }

    public int getSize() {
        return size;
    }

    public int getCapacity() {
        return capacity;
    }
}
