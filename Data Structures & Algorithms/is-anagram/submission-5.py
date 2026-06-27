from collections import Counter

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        f1 = Counter(s)
        f2 = Counter(t)
        if f1 == f2:
            return True
        
        return False
        