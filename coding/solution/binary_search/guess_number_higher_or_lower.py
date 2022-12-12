class Solution:
    def guessNumber(self, n: int) -> int:
        low = 1
        high = n

        while (low <= high):
            mid = low + (high - low) // 2
            guessed = self.guess(mid)
            if (guessed == 0):
                return mid 
            elif guessed < 0:
                high = mid - 1
            else:
                low = mid + 1
        return -1

    def guess(num: int) -> int: 
        if num > 256:
            return -1
        elif num < 256:
            return 1
        else:
            return 0  



print(Solution.guessNumber(Solution, 20000))