#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the compareTriplets function below.
def compareTriplets(a, b):
  
  # Declare and initialize the comparison list
  comparison = [0, 0]
  
  # Compare a and b
  for i in range(len(a)):
    if a[i] > b[i]:
      comparison[0] += 1
    elif b[i] > a[i]:
      comparison[1] += 1
  
  return comparison

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  a = list(map(int, input().rstrip().split()))

  b = list(map(int, input().rstrip().split()))

  result = compareTriplets(a, b)

  fptr.write(' '.join(map(str, result)))
  fptr.write('\n')

  fptr.close()
