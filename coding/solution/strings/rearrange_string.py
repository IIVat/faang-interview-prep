from queue import PriorityQueue
from collections import Counter
from heapq import (heappop, heappush)

class Solution:
    def rearrangeString(s: str, k: int) -> str:
        if k == 0:
            return ""

        counter = Counter(s)

        heap = []

        for (e, c) in counter.most_common():
            heappush(heap, (-c, e))

        res = []
    

        counter = 0
        while heap:
            tmp = []
            for _ in range(min(k, len(s) - len(res))):
                counter += 1
                if not heap:
                  print(counter)
                  return ""
                count,m = heappop(heap)
                res.append(m)
                if count < -1:  
                    tmp.append((count + 1, m))
            for e in tmp:
                heappush(heap, e)

        print(counter)


        if len(res) == len(s):
            return "".join(res)
         

        return ""
        # .join(res) if len(res) == len(s) else ""



print(Solution.rearrangeString("aaaaaaaaaaaa", 3))

