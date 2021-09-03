import time
from util import add_and_multiply, step_one, write_to_file, write
import numpy

def sequential(a, b, size):
    start_time = time.time()
    write("sequential.txt", "Matrices A, B", a, b)
    c = [[0 for row in range(size)] for row in range(size)]
    a, b = step_one(a, b, size)
    c = add_and_multiply(a, b, c, size)
    write_to_file('sequential.txt', 1, a, b, c)

    # other steps (transformations)
    for j in range(1, size):
        for i in range(0, size):
            a[i] = a[i][1:] + a[i][:1]
        b = b[1:] + b[:1]
        c = add_and_multiply(a, b, c, size)

        write_to_file('sequential.txt', j + 1, a, b, c)
    end_time = time.time()
    time1 = end_time - start_time

    print("_______________________")
    for row in c:
        print(row)

    print('Process finished in: ', time1)