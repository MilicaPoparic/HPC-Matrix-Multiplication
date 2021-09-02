import timeit, time
from mpi4py import MPI
import math
import numpy as np
from util import add_and_multiply, step_one, write_to_file, write
from main import a, b, n

if __name__ == '__main__':
    comm = MPI.COMM_WORLD
    rank = comm.Get_rank()
    p = comm.Get_size()
    step, dim1, s1 = 0, 0, 0
    p_sqrt = int(math.sqrt(p-1))
    p_sqrt_2 = p_sqrt ** 2
    range_per_row = int((p - 1) / p_sqrt)
    block_dim = int(n / p_sqrt)
    dim2 = block_dim
    s2 = range_per_row
    rows = {}
    c = [[0 for i in range(n)] for j in range(n)]

    if rank == 0:
        dest = 0
        write("parallel.txt","matrices a, b", a, b)
        start_time = time.time()
        a, b = step_one(a, b, n)
        for i in range(n):
            a_block, b_block, data = [], [], []
            for j in range(dim1, dim2):
                a_block.append(a[j][step:step + block_dim])
                b_block.append(b[j][step:step + block_dim])
            if len(a_block[block_dim-1]) == block_dim:
                print(a_block, " a BLOCK")
                c_block = [[0 for i in range(block_dim)] for j in range(block_dim)]
                data = [a_block, b_block, c_block]
                dest += 1
                if dest == p:
                    dest = 1
                comm.send(data, dest=dest, tag=1)
            step = step + block_dim
            if (i + 1) % block_dim == 0:
                step = 0
                dim1 += block_dim
                dim2 += block_dim

        c_blocks = []
        d = 0
        for i in range(p-1):
            d = i + 1
            mtx = comm.recv(source=d, tag=d)
            c_blocks.append(mtx)

        for l in range(range_per_row):
            rows[l] = []
            for m in range(s1, s2):
                rows[l].append(c_blocks[m])
            s1 += range_per_row
            s2 += range_per_row

        end_time = time.time()
        print(np.bmat([rows[i] for i in rows.keys()]))
        print("Process finished in ", end_time - start_time)
        write("parallel.txt","result: ",np.bmat([rows[i] for i in rows.keys()]), '-------------------------')
    else:
        data = comm.recv(source=0, tag=1)
        for t in range(n):
            add_and_multiply(data[0], data[1], data[2], block_dim)
            write_to_file('parallel.txt',t+rank, data[0], data[1], data[2])
            left_shift_dest = rank-1+p_sqrt if (rank-1) % p_sqrt == 0 else rank - 1
            left_shift_source = rank + 1 - p_sqrt if rank % p_sqrt == 0 else rank + 1
            comm.send([i[0] for i in data[0]], dest=left_shift_dest, tag=left_shift_dest)
            new_col = comm.recv(source=left_shift_source, tag=rank)
            for i in range(block_dim):
                data[0][i] = data[0][i][1:] + [new_col[i]]

            up_shift_dest = rank - p_sqrt if rank - p_sqrt > 0 else rank + range_per_row * (range_per_row - 1)
            up_shift_source = rank - range_per_row * (range_per_row - 1) if (rank + p_sqrt) > p_sqrt_2 else rank + p_sqrt
            comm.send(data[1][0], dest=up_shift_dest, tag=up_shift_dest)
            new_row = comm.recv(source=up_shift_source, tag=rank)
            data[1] = data[1][1:] + [new_row]

        req = comm.send(data[2], dest=0, tag=rank)

