#include <gperftools/profiler.h>
#include <iostream>
#include <fstream>
#include <sstream>
#include <ctime>

using namespace std;

bool isPrime(long long n){
    if (n == 1 | n == 0) return false;
    for (long long i=2; i<n; i++){
        if (n%i == 0) return false;
    }
    return true;
}

int main(){
    long long N = 400000;
    ProfilerStart("prof/primenum.cprof");
    clock_t start = clock();
    for (long long i = 1; i <= N; i++){
        if (isPrime(i)) {
            //cout << i << " ";
        }
    }
    clock_t end = clock();
    ProfilerStop();
    std::cout<< "prime numbers up to " << N << ": "<< (float)(end - start)/ CLOCKS_PER_SEC << " "<< std::endl;;
    return 0;
}