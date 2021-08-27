def step_one(a, b, size):
    for i in range(1, size):
        a[i] = a[i][i:] + a[i][:i]
        column = [row[i] for row in b]
        shifted = column[i:] + column[:i]
        for j in range(0, size):
            b[j][i] = shifted[j]
    return a, b

def add_and_multiply(a, b, c, size):
    for i in range(0, size):
        for j in range(0, size):
            c[i][j] += a[i][j] * b[i][j]
    return c