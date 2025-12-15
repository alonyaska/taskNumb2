from itertools import product, permutations, repeat


def f(k,l,m,n):
    return  (not n) or (k and (not m)) or (l == m)


for x1,x2,x3,x4 in product([0,1], repeat = 4):
    t = (
        (0,1,0,x1,0),
        (x2,0,0,x3,0),
        (1,0,x4,1,0),


)
    if len(t) ==len(set(t)):
        for p in permutations('klmn', r=4):
            if all(f(**dict(zip(p, line))) == line[-1] for line in t):
                print(*p)

