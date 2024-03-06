#include <iostream>
#include <vector>
#include <cstdlib>
#include <ctime>

// Function to generate a random matrix
void generateRandomMatrix(int rows, int cols, std::vector<std::vector<int>>& matrix) {
    matrix.resize(rows, std::vector<int>(cols));
    for (int i = 0; i < rows; ++i) {
        for (int j = 0; j < cols; ++j) {
            matrix[i][j] = rand() % 2 + 1; // Generate random number between 1 and 2
        }
    }
}

// Function to add two matrices
std::vector<std::vector<int>> addMatrices(const std::vector<std::vector<int>>& mat1, const std::vector<std::vector<int>>& mat2) {
    int rows = mat1.size();
    int cols = mat1[0].size();
    std::vector<std::vector<int>> result(rows, std::vector<int>(cols, 0));

    for (int i = 0; i < rows; ++i) {
        for (int j = 0; j < cols; ++j) {
            result[i][j] = mat1[i][j] + mat2[i][j];
        }
    }

    return result;
}

// Function to subtract two matrices
std::vector<std::vector<int>> subtractMatrices(const std::vector<std::vector<int>>& mat1, const std::vector<std::vector<int>>& mat2) {
    int rows = mat1.size();
    int cols = mat1[0].size();
    std::vector<std::vector<int>> result(rows, std::vector<int>(cols, 0));

    for (int i = 0; i < rows; ++i) {
        for (int j = 0; j < cols; ++j) {
            result[i][j] = mat1[i][j] - mat2[i][j];
        }
    }

    return result;
}

// Function to multiply two matrices using Strassen's algorithm
std::vector<std::vector<int>> multiplyMatrices(const std::vector<std::vector<int>>& mat1, const std::vector<std::vector<int>>& mat2) {
    int rows1 = mat1.size();
    int cols1 = mat1[0].size();
    int cols2 = mat2[0].size();

    // Base case: if matrices are 1x1, perform simple multiplication
    if (rows1 == 1 && cols1 == 1 && cols2 == 1) {
        std::vector<std::vector<int>> result(1, std::vector<int>(1, 0));
        result[0][0] = mat1[0][0] * mat2[0][0];
        return result;
    }

    // Split matrices into quadrants
    int newRows = rows1 / 2;
    int newCols = cols1 / 2;

    std::vector<std::vector<int>> A11(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> A12(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> A21(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> A22(newRows, std::vector<int>(newCols));

    std::vector<std::vector<int>> B11(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> B12(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> B21(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> B22(newRows, std::vector<int>(newCols));

    std::vector<std::vector<int>> C11(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> C12(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> C21(newRows, std::vector<int>(newCols));
    std::vector<std::vector<int>> C22(newRows, std::vector<int>(newCols));

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
    std::vector<std::vector<int>> P1 = multiplyMatrices(A11, subtractMatrices(B12, B22));
    std::vector<std::vector<int>> P2 = multiplyMatrices(addMatrices(A11, A12), B22);
    std::vector<std::vector<int>> P3 = multiplyMatrices(addMatrices(A21, A22), B11);
    std::vector<std::vector<int>> P4 = multiplyMatrices(A22, subtractMatrices(B21, B11));
    std::vector<std::vector<int>> P5 = multiplyMatrices(addMatrices(A11, A22), addMatrices(B11, B22));
    std::vector<std::vector<int>> P6 = multiplyMatrices(subtractMatrices(A12, A22), addMatrices(B21, B22));
    std::vector<std::vector<int>> P7 = multiplyMatrices(subtractMatrices(A11, A21), addMatrices(B11, B12));

    // Computing result quadrants
    C11 = addMatrices(subtractMatrices(addMatrices(P5, P4), P2), P6);
    C12 = addMatrices(P1, P2);
    C21 = addMatrices(P3, P4);
    C22 = subtractMatrices(subtractMatrices(addMatrices(P1, P5), P3), P7);

    // Combining result quadrants into the final result matrix
    std::vector<std::vector<int>> result(rows1, std::vector<int>(cols2));
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
    srand(time(0)); // Seed for random number generator using current time
    int sizes[3] = {256, 512, 1024};
    for (int i = 0; i < 3; i++){
        const int SIZE = sizes[i]; // Size of matrices
        std::vector<std::vector<int>> matrix1, matrix2, result;
        // Generate random matrices
        generateRandomMatrix(SIZE, SIZE, matrix1);
        generateRandomMatrix(SIZE, SIZE, matrix2);
        std::clock_t start = std::clock();
        // Multiply matrices using Strassen's algorithm
        result = multiplyMatrices(matrix1, matrix2);
        std::clock_t end = std::clock();
        double elapsed_time = static_cast<double>(end - start) / CLOCKS_PER_SEC;
        std::cout << "Elapsed time for " << SIZE << " x " << SIZE <<" matrix multiplication: " << elapsed_time << " seconds." << std::endl;
    }
    
    return 0;
}
