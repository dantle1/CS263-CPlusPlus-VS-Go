import numpy as np
import sys 

if __name__ == '__main__':
    f = open('matrix.in', 'w')
    m, n, l = sys.stdin.readline().split()
    f.write(" ".join([m,n,l]) + "\n")

    # A = M(m x n), B = M(n x l)
    A = np.random.uniform(1, 2, (int(m),int(n)))
    B = np.random.uniform(1, 2, (int(n),int(l)))

    np.savetxt(f, A, fmt='%f', delimiter = ' ')
    np.savetxt(f, B, fmt='%f', delimiter = ' ')

    f.close()
