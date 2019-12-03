import math
from decimal import Decimal

def continued_fraction(s):
    e = []
    m = 0
    d = 1
    a = int(math.floor(math.sqrt(s)))
    while True:
        m = d * a - m
        d = (s - m * m) / d
        if d == 0:
            return 0
        a = int(math.floor((math.sqrt(s) + m) / d))
        v = (m, d)
        if v in e:
            return len(e) - e.index(v)
        e.append((m, d))

odd_periods = 0
for i in range(2, 10001):
    period = continued_fraction(i)
    if period > 0 and period % 2 == 1:
        odd_periods += 1

print 'odd periods', odd_periods
