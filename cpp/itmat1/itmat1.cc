#include <iostream>
#include <vector>
#include <cstdlib>
#include <ctime>
#include <fstream>
#include <string>
#include <sstream>
#include <gperftools/profiler.h>

// Function to multiply two matrices
std::vector<std::vector<double>> multiplyMatrices(const std::vector<std::vector<double>>& mat1, const std::vector<std::vector<double>>& mat2) {
    int rows1 = mat1.size();
    int cols1 = mat1[0].size();
    int cols2 = mat2[0].size();

    std::vector<std::vector<double>> result(rows1, std::vector<double>(cols2, 0));

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

    ProfilerStart("prof/itmat1.cprof");

    std::vector<std::vector<double>> result = multiplyMatrices(A, B);

    ProfilerStop();

    // for (auto &row : result) {
    //     for (const auto &val : row) {
    //         std::cout << val << "\n";
    //     }
    // }

    f.close();
    
    return 0;
}
