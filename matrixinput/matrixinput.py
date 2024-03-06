import numpy as np
import pandas as pd
import sys

def main():
    file1 = open('matrixinput1.txt', 'a')
    file2 = open('matrixinput2.txt', 'a')
    # number of rows in first matrix
    print("Enter m: ")
    m = int(sys.stdin.readline())
    file1.write(str(m) + "\n")
    # number of columns in first matrix and rows in second matrix
    print("Enter n: ")
    n = int(sys.stdin.readline())
    file1.write(str(n) + "\n")
    file2.write(str(n) + "\n")
    # number of columns in second matrix
    print("Enter p: ")
    p = int(sys.stdin.readline())
    file2.write(str(p) + "\n")
    # first matrix: m x n
    matrix1 = np.random.randint(2, size=(m,n))
    # second matrix: n x p
    matrix2 = np.random.randint(2, size=(n,p))
    np.savetxt(file1, matrix1, fmt='%d', delimiter = ",", header="[", footer="]", comments = "")
    np.savetxt(file2, matrix2, fmt='%d', delimiter = ",", header="[", footer="]", comments = "")
    file1.close()
    file2.close()

if __name__== "__main__":
    main()