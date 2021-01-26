#!/bin/python3

import math
import os
import random
import re
import sys
    
# Complete the minimumBribes function below.
def minimumBribes(q):
    
    # Starting values
    numBribes = 0
    isChaotic = False
    

    for index, value in enumerate(q):
        
        # If the value has moved more than 2 spaces mark this array
        # as chaotic and bail
        if (value > index + 3):
            isChaotic = True
            break
        else:
            
            # Loop through from val - 2 to index and flag any bribes that
            # occured after an item has been bribed
            for j in range(max(value - 2, 0), index):
                if (q[j] > value):
                    numBribes += 1
       
    if (isChaotic):
        print("Too chaotic")
    else:
        print(numBribes)
    

if __name__ == '__main__':
    t = int(input())

    for t_itr in range(t):
        n = int(input())

        q = list(map(int, input().rstrip().split()))

        minimumBribes(q)
