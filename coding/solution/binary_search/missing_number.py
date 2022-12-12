from typing import List

#not optimal complexity O(n*logn)
class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        start = 0
        while start < len(nums):
            num = nums[start] 
            if num < len(nums) and num != start:
                nums[start], nums[num] = nums[num], nums[start]
            else:
                start += 1

        for (i, n) in enumerate(nums): 
            if i != n:
                return i
        return len(nums)

            
print(Solution.missingNumber(Solution, [9,6,4,2,3,5,7,0,1]))

print(0^0^1)
