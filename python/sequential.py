import timeit, time
from mpi4py import MPI
import math, copy
import numpy as np
from util import add_and_multiply, step_one, write_to_file, write

# treba mi lista velicine p da cuva parove a i b i c mozda

blocks = []


def save_blocks(data):
    blocks.append([data[0], data[1], data[2]])
    return blocks


def sequential(a, b, n, p):
    # f = open("resources/sequential_weak25.txt", "a")
    step, dim1, s1 = 0, 0, 0
    p_sqrt = int(math.sqrt(p - 1))
    range_per_row = int((p - 1) / p_sqrt)
    block_dim = int(n / p_sqrt)
    s2 = range_per_row
    rows = {}
    # write("sequential.txt", "matrices a, b", a, b)
    start_time = time.time()
    a, b = step_one(a, b, n)
    a_block, b_block = [], []

    for i in range(0, n, block_dim):
        mtx_a = a[i:i + block_dim]
        mtx_b = b[i:i + block_dim]
        for j in range(0, n, block_dim):
            for k in range(len(mtx_a)):
                a_block.append(mtx_a[k][j:j + block_dim])
                b_block.append(mtx_b[k][j:j + block_dim])
                if (len(a_block) == block_dim):
                    c_block = [[0 for i in range(block_dim)] for j in range(block_dim)]
                    data = [a_block, b_block, c_block]
                    blocks = save_blocks(data)
                    a_block, b_block = [], []

    c_blocks = [[[0 for k in range(block_dim)] for j in range(block_dim)] for i in range(p-1)]

    for m in range(n):
        blocks_shifted = copy.deepcopy(blocks)
        for i in range(p - 1):
            process = i + 1
            add_and_multiply(blocks[i][0], blocks[i][1], blocks[i][2], block_dim)
            # write_to_file('sequential.txt', m + process, blocks[i][0], blocks[i][1], blocks[i][2])
            for s in range(block_dim):
                for k in range(block_dim):
                    c_blocks[i][s][k] += blocks[i][2][s][k]
            left_shift_dest = process - 1 + p_sqrt if (process - 1) % p_sqrt == 0 else process - 1

            new_col = [i[0] for i in blocks[left_shift_dest - 1][0]]
            for j in range(block_dim):
                blocks_shifted[i][0][j] = blocks[i][0][j][1:] + [new_col[j]]

            up_shift_dest = process - p_sqrt if process - p_sqrt > 0 else process + range_per_row * (range_per_row - 1)

            new_row = (blocks[up_shift_dest - 1][1][0])
            blocks_shifted[i][1] = blocks[i][1][1:] + [new_row]


        blocks = blocks_shifted

    for l in range(range_per_row):
        rows[l] = []
        for m in range(s1, s2):
            rows[l].append(c_blocks[m])
        s1 += range_per_row
        s2 += range_per_row

    end_time = time.time()
    print(np.bmat([rows[i] for i in rows.keys()]))
    print("Process finished in ", end_time - start_time)
    # write("sequential.txt", "result: ", np.bmat([rows[i] for i in rows.keys()]), '-------------------------')
    # f.write(str(end_time - start_time) + "\n")
    # f.close()