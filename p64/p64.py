import math
from decimal import Decimal

def continued_fraction(s):
    e = []
    m = 0
    d = 1
    a = int(math.floor(math.sqrt(s)))
    for i in range(40):
        m = d * a - m
        d = (s - m * m) / d
        if d == 0:
            return e
        a = int(math.floor((math.sqrt(s) + m) / d))
        e.append((a, m, d))
    return e

def find_period(s):
    p = 1
    while p <= (len(s) / 2):
        print p
        p += 1

expansion = continued_fraction(7)
print expansion
find_period(expansion)

'''
for i in range(2, 14):
    expansion = continued_fraction(i)
    print i, expansion
'''
