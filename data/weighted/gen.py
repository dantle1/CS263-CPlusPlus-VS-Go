import numpy as np
import sys 
import networkx

if __name__ == '__main__':
    f = open('weighted.in', 'w')
    n = int(sys.stdin.readline())
    m = np.random.randint(n, n**2)
    f.write(str(n) + ' ' + str(m) + '\n')

    edges = np.random.randint(0, n, (m, 3))
    np.savetxt(f, edges, fmt='%d', delimiter=' ')

    f.close()