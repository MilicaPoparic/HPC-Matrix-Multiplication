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

def write_to_file(filename, shifting_step, a, b, c):
    file = open(filename, "a")  # append mode

    file.write('Shifting step:' + str(shifting_step) + '\n')
    file.write('A:' + str(shifting_step) + ': ' + str(a) + '\n')
    file.write('B:' + str(shifting_step) + ': ' + str(b) + '\n')
    file.write('C:' + str(shifting_step) + ': ' + str(c) + '\n')

    file.close()

def write(filename, marker, c, d):
    file = open(filename, "a")  # append mode
    file.write(marker + '\n')
    file.write(str(c) + '\n')
    file.write(str(d) + '\n')
    file.close()
