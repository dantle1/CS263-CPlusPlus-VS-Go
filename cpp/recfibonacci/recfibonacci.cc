// Fibonacci Series using Recursion
#include <bits/stdc++.h>
#include <gperftools/profiler.h>
using namespace std;
 
int fib(int n)
{
    if (n <= 1)
        return n;
    return fib(n - 1) + fib(n - 2);
}

int main(){
    int N = 1000;
    ProfilerStart("prof/recfibonacci.cprof");
    cout << fib(N) << endl;
    ProfilerStop();
    return 0;
}