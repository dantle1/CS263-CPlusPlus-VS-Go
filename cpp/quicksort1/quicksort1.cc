#include <bits/stdc++.h>
#include <gperftools/profiler.h>
#include <ctime>
using namespace std;

int partition(std::vector<long long>& arr, int low, int high)
{
    int pivot = arr[high];
    int i = (low-1);

    for(int j=low; j<=high; j++){
        if(arr[j] < pivot){
            i++;
            swap(arr[i],arr[j]);
        }
    }
    swap(arr[i+1], arr[high]);
    return (i+1);
}

void quickSort(std::vector<long long>& arr, int low, int high){
    if(low < high){
        int pivot = partition(arr,low,high);
        quickSort(arr,low,pivot-1);
        quickSort(arr,pivot+1,high);
    }
}

int main(){
    fstream f("../data/array/array.in");
    if (!f.good()){
        std::cerr << "double check file path\n";
        return 1;
    }

    string line;
    getline(f,line);
    istringstream arr_length(line);
    int n;
    arr_length >> n;
    //cout << n << endl;

    long long num;
    getline(f,line);
    std::istringstream iss(line);
    std::vector<long long> arr;
    while (iss >> num){
        arr.push_back(num);
    }
    
    clock_t start = clock();
    quickSort(arr, 0, n-1);
    clock_t end = clock();
    
    for (auto &element : arr){
        cout << element << endl;
    }

    //cout<< "Quicksort time for " << n << " elements: "<< (float)(end - start)/ CLOCKS_PER_SEC << " "<<endl;;
    f.close();
    return 0;
}