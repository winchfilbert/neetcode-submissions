class Solution {
    public HashSet<Integer> numbers = new HashSet<Integer>();

    public boolean hasDuplicate(int[] nums) {
        for (int i = 0 ; i < nums.length ; i++){
            if(numbers.contains(nums[i]) == true){
                return true;
            };
            numbers.add(nums[i]);
        }
        return false;
    }
}