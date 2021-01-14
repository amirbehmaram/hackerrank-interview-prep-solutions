#!/bin/python3

import math
import os
import random
import re
import sys

# Calculate the sum of a single hourglass
def sumSingleHourglass(i, j, arr):
    sum = 0
    
    # Sum Top Row
    sum += arr[i-1][j-1]
    sum += arr[i-1][j]
    sum += arr[i-1][j+1]
    
    # Sum Middle
    sum += arr[i][j]
    
    # Sum Bottom
    sum += arr[i+1][j-1]
    sum += arr[i+1][j]
    sum += arr[i+1][j+1]

    return sum

# Solution:
# Loop through only our center hourglass indeces, then calculate
# each of their values at a time
def hourglassSum(arr):
    index = 1
    maxSum = 0
    
    while(index < (len(arr) - 1)):
        jIndex = 1
        
        while(jIndex < (len(arr[index]) - 1)):
            hourglassSum = sumSingleHourglass(index, jIndex, arr)
            
            # Corner case
            # Always update maxSum to be the first hourglass sum
            # If we have all negative numbers in our array then all of
            # the hourglass sums will be less than our default 0
            if (index == 1 and jIndex == 1):
                maxSum = hourglassSum
                
            # Update the max sum if the calculated sume is greater than our current
            if (hourglassSum > maxSum):
                maxSum = hourglassSum
            
            jIndex += 1
        
        index += 1
    
    return maxSum
            
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    result = hourglassSum(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
