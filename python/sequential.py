import time
from util import add_and_multiply, step_one
import numpy

def sequential(a, b, size):
    start_time = time.time()
    c = [[0 for row in range(size)] for row in range(size)]
    a, b = step_one(a, b, size)
    c = add_and_multiply(a, b, c, size)

    # other steps (transformations)
    for i in range(1, size):
        for i in range(0, size):
            a[i] = a[i][1:] + a[i][:1]
        b = b[1:] + b[:1]
        c = add_and_multiply(a, b, c, size)

    end_time = time.time()
    time1 = end_time - start_time

    for row in c:
        print(row)

    print('Process finished in: ', time1)
