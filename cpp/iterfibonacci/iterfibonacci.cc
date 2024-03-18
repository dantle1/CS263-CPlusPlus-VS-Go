#include <gperftools/profiler.h>
#include <iostream>

using namespace std;

int fibonacci(int n){
    int a = 0, b = 1;
    int sum = a + b;
    for (int i = 3; i <= n; i++){
       a = b;
       b = sum;
       sum = a+b;
    }
    return sum;
}

int main(){
    int n = 10;
    ProfilerStart("prof/iterfibonacci.cprof");
    int fib = fibonacci(n);
    ProfilerStop();
    cout << "fibonacci for " << n << ": " << fib << endl;
}