class Solution {
    public boolean hasDuplicate(int[] nums) {
        int flag = 0;
        for (int i = 0 ; i < nums.length ; i++){
            for (int j = i+1 ; j < nums.length; j++){
                if (nums[i] == nums[j]){
                    flag = 1;
                }
            }
        }
        if (flag == 1){
            return true;
        }
        else{
            return false;
        }
    }
}