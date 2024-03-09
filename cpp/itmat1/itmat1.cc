#include <iostream>
#include <vector>
#include <cstdlib>
#include <ctime>

// Function to generate a random matrix
void generateRandomMatrix(int rows, int cols, std::vector<std::vector<int>>& matrix) {
    matrix.resize(rows, std::vector<int>(cols));
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            matrix[i][j] = rand() % 2 + 1; // Generate random number between 1 and 2
        }
    }
}

// Function to multiply two matrices
std::vector<std::vector<int>> multiplyMatrices(const std::vector<std::vector<int>>& mat1, const std::vector<std::vector<int>>& mat2) {
    int rows1 = mat1.size();
    int cols1 = mat1[0].size();
    int cols2 = mat2[0].size();

    std::vector<std::vector<int>> result(rows1, std::vector<int>(cols2, 0));

    for (int i = 0; i < rows1; i++) {
        for (int j = 0; j < cols2; j++) {
            for (int k = 0; k < cols1; k++) {
                result[i][j] += mat1[i][k] * mat2[k][j];
            }
        }
    }

    return result;
}

int main() {
    srand(time(0)); // Seed for random number generator using current time
    int sizes[4] = {1000, 1500, 2000, 2500};
    for (int i = 0; i < 4; i++){
        const int SIZE = sizes[i]; // Size of matrices
        std::vector<std::vector<int>> matrix1, matrix2, result;
        // Generate random matrices
        generateRandomMatrix(SIZE, SIZE, matrix1);
        generateRandomMatrix(SIZE, SIZE, matrix2);
        std::clock_t start = std::clock();
        // Multiply matrices
        result = multiplyMatrices(matrix1, matrix2);
        std::clock_t end = std::clock();
        double elapsed_time = static_cast<double>(end - start) / CLOCKS_PER_SEC;
        std::cout << "Elapsed time for " << SIZE << " x " << SIZE <<" matrix multiplication: " << elapsed_time << " seconds." << std::endl;
    }
    
    return 0;
}
