#include <bits/stdc++.h>
#include <gperftools/profiler.h>
#include <ctime>
using namespace std;

void merge(vector<long long> &array, int const left, int const mid,
           int const right)
{
    long long const subArrayOne = mid - left + 1;
    long long const subArrayTwo = right - mid;
 
    // Create temp arrays
    auto *leftArray = new long long[subArrayOne],
         *rightArray = new long long[subArrayTwo];
 
    // Copy data to temp arrays leftArray[] and rightArray[]
    for (auto i = 0; i < subArrayOne; i++)
        leftArray[i] = array[left + i];
    for (auto j = 0; j < subArrayTwo; j++)
        rightArray[j] = array[mid + 1 + j];
 
    auto indexOfSubArrayOne = 0, indexOfSubArrayTwo = 0;
    int indexOfMergedArray = left;
 
    // Merge the temp arrays back into array[left..right]
    while (indexOfSubArrayOne < subArrayOne
           && indexOfSubArrayTwo < subArrayTwo) {
        if (leftArray[indexOfSubArrayOne]
            <= rightArray[indexOfSubArrayTwo]) {
            array[indexOfMergedArray]
                = leftArray[indexOfSubArrayOne];
            indexOfSubArrayOne++;
        }
        else {
            array[indexOfMergedArray]
                = rightArray[indexOfSubArrayTwo];
            indexOfSubArrayTwo++;
        }
        indexOfMergedArray++;
    }
 
    // Copy the remaining elements of
    // left[], if there are any
    while (indexOfSubArrayOne < subArrayOne) {
        array[indexOfMergedArray]
            = leftArray[indexOfSubArrayOne];
        indexOfSubArrayOne++;
        indexOfMergedArray++;
    }
 
    // Copy the remaining elements of
    // right[], if there are any
    while (indexOfSubArrayTwo < subArrayTwo) {
        array[indexOfMergedArray]
            = rightArray[indexOfSubArrayTwo];
        indexOfSubArrayTwo++;
        indexOfMergedArray++;
    }
    delete[] leftArray;
    delete[] rightArray;
}
 
// begin is for left index and end is right index
// of the sub-array of arr to be sorted
void mergeSort(vector<long long> &array, int const begin, int const end)
{
    if (begin >= end)
        return;
 
    int mid = begin + (end - begin) / 2;
    mergeSort(array, begin, mid);
    mergeSort(array, mid + 1, end);
    merge(array, begin, mid, end);
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
    
    //clock_t start = clock();
    mergeSort(arr,0,n-1);
    //clock_t end = clock();
    
    for (auto &element : arr){
        cout << element << endl;
    }

    //cout<< "Quicksort time for " << n << " elements: "<< (float)(end - start)/ CLOCKS_PER_SEC << " "<<endl;;
    f.close();
    return 0;
}