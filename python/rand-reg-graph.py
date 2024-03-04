import networkx
import sys 

'''We use this to create random 2-out directed graphs. These populate the heap when we test our
gc algorithm, and we have roots point to random heap nodes at each iteration'''
if __name__ == '__main__':
    n, d = [int(k) for k in sys.stdin.readline().split()]
    G = networkx.random_k_out_graph(n, d, 0.2)
    print(n)
    for e in G.edges:
        print(*e)