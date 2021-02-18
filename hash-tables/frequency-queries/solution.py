#!/bin/python3

import math
import os
import random
import re
import sys

# So we can use Counter
from collections import Counter

# Complete the freqQuery function below.
def freqQuery(queries):
    
    db = Counter()
    freqCount = Counter()
    results = []
    
    for operation, value in queries:
        
        # From the problem statement we know that operation
        # is going to exist in the set {1,2,3} so we don't
        # need to account for anything outside of that in our
        # operation handoff
        
        # Insert opertion
        if (operation == 1):
            
            db[freqCount[value]] -= 1
            freqCount[value] += 1
            db[freqCount[value]] += 1
        
        # Delete operation
        elif (operation == 2):
            
            if (freqCount[value] > 0):
                db[freqCount[value]] -= 1
                freqCount[value] -= 1
                db[freqCount[value]] += 1
            
        # Frequency operation
        else:   
            if (db[value] > 0):
                results.append(1)
            else:
                results.append(0) 
    
    return results
                    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    queries = []

    for _ in range(q):
        queries.append(list(map(int, input().rstrip().split())))

    ans = freqQuery(queries)

    fptr.write('\n'.join(map(str, ans)))
    fptr.write('\n')

    fptr.close()
