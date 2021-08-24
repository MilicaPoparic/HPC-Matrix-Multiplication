import timeit, time
from mpi4py import MPI
import math

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

def add_and_multiply(a, b, c, size):
    for i in range(0, size):
        for j in range(0, size):
            c[i][j] += a[i][j] * b[i][j]

if __name__ == '__main__':
    comm = MPI.COMM_WORLD
    rank = comm.Get_rank()
    step = 0
    n = 4
    p = 5
    p_sqrt = int(math.sqrt(p-1))
    p_sqrt_2 = p_sqrt ** 2
    block_dim = int(n / p_sqrt)
    dim1 = 0
    dim2 = block_dim
    a = [[15, -11, -12, 12], [-15, -2, 15, -15], [12, 14, -12, -6], [-1, -8, 16, -13]]
    b = [[0, 15, 14, 9], [-3, -7, -12, -4], [10, 10, -16, 15], [-13, -3, 9, 3]]

    # print(a[0][0:2])
    #ako budem slala celu c matricu sta dobijam tima
    #moram da saljem parcice koje izvlacim da bi ostali procesi radili shift i opet dalje mnozili :) samim tim mi treba i samo taj deo c1

    c =[[0 for i in range(n)] for j in range(n)]

    if rank == 0:
        dest = 0
        a, b = step_one(a, b, n) #zapucana dimenzija i ovde
        for i in range(n): #postaviti da bude n a ne ovo
            a_block, b_block = [], []
            for j in range(dim1, dim2):
                a_block.append(a[j][step:step + block_dim])
                b_block.append(b[j][step:step + block_dim])
            c_block = [[ 0 for i in range(block_dim) ] for j in range(block_dim)]
            data = [a_block, b_block, c_block]
            comm.send(data, dest=i+1, tag=1)
            step = step + block_dim
            if (i + 1) % block_dim == 0:
                step = 0
                dim1 += block_dim
                dim2 += block_dim

        #primam c blokove i sklapam resenje
        rows1, rows2=[], []
        for i in range(4):
            dest = i+1
            mtx = comm.recv(source=dest, tag=dest)
            print(mtx, dest)


            #treba nekako da sacuvam to sto mi stigne da bih posle sabirala sa ostatkom!
            #dobijam svaki deo svake c matrice i ne znam kako da sabiram adekvatne delove?
            # print("C BLOK", mtx)
        print(rows1)
    else:
        #treba else pa svi ostali koraci ovde
        data = comm.recv(source=0, tag=1)
        # print(data[0], " Data 0")
        result = []
        #mnozenje i sabiranje
        for t in range(n):
            add_and_multiply(data[0], data[1], data[2], block_dim)
            # print(data[2], rank)
            result.append(data[2])

            left_shift_dest = rank-1+p_sqrt if (rank-1) % p_sqrt == 0 else rank - 1
            left_shift_source = rank + 1 - p_sqrt if rank % p_sqrt == 0 else rank + 1

            up_shift_dest = rank - p_sqrt if rank - p_sqrt > 0 else rank + block_dim*(block_dim-1)
            up_shift_source = rank - block_dim*(block_dim-1) if (rank+p_sqrt) > p_sqrt_2 else rank + p_sqrt

            #left_shift
            comm.send([i[0] for i in data[0]], dest=left_shift_dest, tag=left_shift_dest)
            new_col = comm.recv(source=left_shift_source, tag=rank)
            for i in range(block_dim):
                data[0][i] = data[0][i][1:] + [new_col[i]]
            # print("SHIFTOVANOOOOO" , data[0])

            #up_shift
            comm.send(data[1][0], dest=up_shift_dest, tag=up_shift_dest)
            new_row = comm.recv(source=up_shift_source, tag=rank)
            data[1] = data[1][1:] + [new_row]
            # print("SIFT NA GORE", data[1])

        req = comm.send(data[2], dest=0, tag=rank)
    #sad treba da izmnozim i da saljem dalje na sift i mnozenje
    #treba da izmnozim blokove i onda da saljem na sift i ponovim proces i da saljem


