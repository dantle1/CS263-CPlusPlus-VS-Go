import numpy as np
import sys 

if __name__ == '__main__':
    f = open('graph.in', 'w')
    n = int(sys.stdin.readline())
    f.write(str(n) + '\n')

    E = np.random.randint(0, n - 1, (2 * n, 2))
    E[:n,0] = np.arange(0, n, 1)
    E[n:,0] = np.arange(0, n, 1)
    for i in range(2 * n):
        if E[i][0] == E[i][1]:
            E[i][1] = (E[i][1] + 1) % n
    np.savetxt(f, E, fmt="%d", delimiter=' ')

    f.close()