#!/bin/python3

import math
import os
import random
import re
import sys

from collections import Counter

# Complete the twoStrings function below.
def twoStrings(s1, s2):
    
    # Get our Counters for each string
    strOneCount = Counter(s1)
    strTwoCount = Counter(s2)
    
    # If you subtract string 1's characters from string 2
    # and string 2 is changed, then we know they have at least
    # 1 shared character
    diffStrs = Counter(s2) - Counter(s1)
    
    if (strTwoCount != diffStrs):
        return "YES"
    else:
        return "NO"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        s1 = input()

        s2 = input()

        result = twoStrings(s1, s2)

        fptr.write(result + '\n')

    fptr.close()
