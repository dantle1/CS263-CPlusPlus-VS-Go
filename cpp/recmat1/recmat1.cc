#include <iostream>
#include <vector>
#include <cstdlib>
#include <ctime>
#include <fstream>
#include <string>
#include <sstream>
#include <gperftools/profiler.h>

// Function to add two matrices
std::vector<std::vector<double>> add(const std::vector<std::vector<double>>& mat1, const std::vector<std::vector<double>>& mat2) {
    int rows = mat1.size();
    int cols = mat1[0].size();
    std::vector<std::vector<double>> result(rows, std::vector<double>(cols, 0));

    for (int i = 0; i < rows; ++i) {
        for (int j = 0; j < cols; ++j) {
            result[i][j] = mat1[i][j] + mat2[i][j];
        }
    }

    return result;
}

// Function to subtract two matrices
std::vector<std::vector<double>> sub(const std::vector<std::vector<double>>& mat1, const std::vector<std::vector<double>>& mat2) {
    int rows = mat1.size();
    int cols = mat1[0].size();
    std::vector<std::vector<double>> result(rows, std::vector<double>(cols, 0));

    for (int i = 0; i < rows; ++i) {
        for (int j = 0; j < cols; ++j) {
            result[i][j] = mat1[i][j] - mat2[i][j];
        }
    }

    return result;
}

// Function to multiply two matrices using Strassen's algorithm
std::vector<std::vector<double>> multiplyMatrices(const std::vector<std::vector<double>>& mat1, const std::vector<std::vector<double>>& mat2) {
    int rows1 = mat1.size();
    int cols1 = mat1[0].size();
    int cols2 = mat2[0].size();

    // Base case: if matrices are 1x1, perform simple multiplication
    if (rows1 == 1 && cols1 == 1 && cols2 == 1) {
        std::vector<std::vector<double>> result(1, std::vector<double>(1, 0));
        result[0][0] = mat1[0][0] * mat2[0][0];
        return result;
    }

    // Split matrices into quadrants
    int newRows = rows1 / 2;
    int newCols = cols1 / 2;

    std::vector<std::vector<double>> A11(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> A12(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> A21(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> A22(newRows, std::vector<double>(newCols));

    std::vector<std::vector<double>> B11(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> B12(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> B21(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> B22(newRows, std::vector<double>(newCols));

    std::vector<std::vector<double>> C11(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> C12(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> C21(newRows, std::vector<double>(newCols));
    std::vector<std::vector<double>> C22(newRows, std::vector<double>(newCols));

    // Splitting input matrices into quadrants
    for (int i = 0; i < newRows; ++i) {
        for (int j = 0; j < newCols; ++j) {
            A11[i][j] = mat1[i][j];
            A12[i][j] = mat1[i][j + newCols];
            A21[i][j] = mat1[i + newRows][j];
            A22[i][j] = mat1[i + newRows][j + newCols];

            B11[i][j] = mat2[i][j];
            B12[i][j] = mat2[i][j + newCols];
            B21[i][j] = mat2[i + newRows][j];
            B22[i][j] = mat2[i + newRows][j + newCols];
        }
    }

    // Recursive steps for multiplication
    std::vector<std::vector<double>> P1 = multiplyMatrices(A11, sub(B12, B22));
    std::vector<std::vector<double>> P2 = multiplyMatrices(add(A11, A12), B22);
    std::vector<std::vector<double>> P3 = multiplyMatrices(add(A21, A22), B11);
    std::vector<std::vector<double>> P4 = multiplyMatrices(A22, sub(B21, B11));
    std::vector<std::vector<double>> P5 = multiplyMatrices(add(A11, A22), add(B11, B22));
    std::vector<std::vector<double>> P6 = multiplyMatrices(sub(A12, A22), add(B21, B22));
    std::vector<std::vector<double>> P7 = multiplyMatrices(sub(A11, A21), add(B11, B12));

    // Computing result quadrants
    C11 = add(sub(add(P5, P4), P2), P6);
    C12 = add(P1, P2);
    C21 = add(P3, P4);
    C22 = sub(sub(add(P1, P5), P3), P7);

    // Combining result quadrants into the final result matrix
    std::vector<std::vector<double>> result(rows1, std::vector<double>(cols2));
    for (int i = 0; i < newRows; ++i) {
        for (int j = 0; j < newCols; ++j) {
            result[i][j] = C11[i][j];
            result[i][j + newCols] = C12[i][j];
            result[i + newRows][j] = C21[i][j];
            result[i + newRows][j + newCols] = C22[i][j];
        }
    }

    return result;
}

int main() {
    std::fstream f("../data/matrix/matrix.in");

    if (!f.good()) {
        std::cerr << "double check file path\n";
        return 1;
    }

    std::string line;
    std::getline(f, line);
    std::istringstream dimensions(line);
    int m, n, l;
    dimensions >> m >> n >> l;

    std::vector<std::vector<double>> A(m, std::vector<double>(n, 0)), B(n, std::vector<double>(l, 0));
    for (int i = 0; i < m && std::getline(f, line); i++) {
        std::istringstream row(line);
        for (int j = 0; j < n; j++) {
            row >> A[i][j];
        }
    }
    for (int i = 0; i < n && std::getline(f, line); i++) {
        std::istringstream row(line);
        for (int j = 0; j < l; j++) {
            row >> B[i][j];
        }
    }

    std::vector<std::vector<double>> result = multiplyMatrices(A, B);

    for (auto &row : result) {
        for (const auto &val : row) {
            std::cout << val << "\n";
        }
    }

    f.close();
    
    return 0;
}
