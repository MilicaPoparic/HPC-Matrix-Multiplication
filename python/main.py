import timeit, time, os
from mpi4py import MPI

def step_one(a, b, size):
    for i in range(1, size):
        a[i] = a[i][i:] + a[i][:i]

        column = [row[i] for row in b]
        shifted = column[i:] + column[:i]
        for j in range(0, size):
            b[j][i] = shifted[j]
    return a, b


def multiply(a, b, size):
    c = [[0 for row in range(size)] for row in range(size)]
    for i in range(0, size):
        for j in range(0, size):
            c[i][j] = a[i][j] * b[i][j]
    return c


def add(a, b, size):
    c = [[0 for row in range(size)] for row in range(size)]
    for i in range(0, size):
        for j in range(0, size):
            c[i][j] = a[i][j] + b[i][j]
    return c


def sequential(size):
    start_time = time.time()
    c = [[0 for row in range(size)] for row in range(size)]
    mtx1 = [[-4, 1, 7, 11], [-2, 0, -1, 10], [-5, 0, 0, 10], [0, 2, 9, 1]]
    mtx2 = [[1, 1, 1, 1], [2, 2, 2, 2], [3, 3, 3, 3], [4, 4, 4, 4]]
    a, b = step_one(mtx1, mtx2, size)
    c = multiply(a, b, size)

    # other steps (transformations)
    for i in range(1, size):
        for i in range(0, size):
            a[i] = a[i][1:] + a[i][:1]
        b = b[1:] + b[:1]
        cn = multiply(a, b, size)
        c = add(c, cn, size)

    end_time = time.time()
    time1 = end_time - start_time

    for row in c:
        print(row)

    print('Process finished in: ', time1)


if __name__ == '__main__':
    #sequential(4)
    p = 5
    # a = [[15, -11, -12, 12], [-15, -2, 15, -15], [12, 14, -12, -6], [-1, -8, 16, -13]]
    # b = [[0, 15, 14, 9], [-3, -7, -12, -4], [10, 10, -16, 15], [-13, -3, 9, 3]]
    # step = 0
    os.system("mpiexec -n {0} python -m mpi4py parallel.py".format(p))
    # dim1 = 0
    # dim2 = 2
    #j prva dva puta 0 posle dva puta 1

    # for i in range(4):
    #     a_block, b_block = [], []
    #     for j in range(dim1, dim2):  # izmeniti da bude n
    #         a_block.append(a[j][step:step + 2])
    #         b_block.append(b[j][step:step + 2])
    #     print(i, a_block)
    #     step = step + 2
    #     if (i+1) % 2 == 0:
    #         step = 0
    #         dim1 += 2
    #         dim2 += 2




