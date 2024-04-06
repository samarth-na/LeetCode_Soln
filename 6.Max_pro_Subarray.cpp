/* 152. Maximum Product Subarray

Given an integer array nums, find a subarray that has the largest product, and
return the product.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

class Solution {
public:
  int maxProduct(vector<int> &nums) {
    int maxPro = nums[0];
    int curPro = 1;

    for (int n : nums) {
      curPro = max(curPro, 0);
      curPro *= n;
      maxPro = max(maxPro, curPro);
    }
    return maxPro;
  }
};
