import sys

t = int(sys.stdin.readline())

arr = []

for k in range(t):

    # number of integers
    n = int(sys.stdin.readline())

    minV = None
    maxV = None

    lst = list(map(int, sys.stdin.readline().split()))
    for x in lst:
        if minV is None or x < minV:
            minV = x
        if maxV is None or x > maxV:
            maxV = x
    
    arr.append(maxV - minV)

for x in arr:
    print(x)
