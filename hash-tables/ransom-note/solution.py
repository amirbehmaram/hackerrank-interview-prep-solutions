#!/bin/python3

import math
import os
import random
import re
import sys

# So we can use Counters
import collections

# Complete the checkMagazine function below.
def checkMagazine(magazine, note):
    
    # Setup out counters
    magCounter = collections.Counter(magazine)
    noteCounter = collections.Counter(note)
    
    # If we subtract the magazine counter from the notecounter
    # and the result is empty then noteCounter is a subset of
    # the magCounter
    if (not noteCounter - magCounter):
        print('Yes')
    else:
        print("No")
    

if __name__ == '__main__':
    mn = input().split()

    m = int(mn[0])

    n = int(mn[1])

    magazine = input().rstrip().split()

    note = input().rstrip().split()

    checkMagazine(magazine, note)
