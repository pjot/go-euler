import math
from decimal import Decimal

def continued_fraction(s):
    e = []
    m = 0
    d = 1
    a = int(math.floor(math.sqrt(s)))
    limit = int(math.log(s) * math.sqrt(s)) + 4
    for i in range(limit):
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
        period = True
        for i in range(p):
            if s[i] != s[i + p]:
                period = False
        if period:
            return p
        p += 1
    return 0

odd_periods = 0
for i in range(2, 10001):
    expansion = continued_fraction(i)
    period = find_period(expansion)
    if period > 0 and period % 2 == 1:
        odd_periods += 1

print 'odd periods', odd_periods
