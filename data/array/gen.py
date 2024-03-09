import numpy as np
import sys 

if __name__ == '__main__':
    f = open('array.in', 'w')
    n = int(sys.stdin.readline())
    f.write(str(n) + '\n')
    A = np.random.randint(1, 10, (1, n))
    np.savetxt(f, A, fmt="%d", delimiter=' ')

    f.close()