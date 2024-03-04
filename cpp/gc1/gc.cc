/* Analogue to the implementation in gc.go */

#include <string>
#include <vector>
#include <iostream>
#include <fstream>

enum Color {
    WHITE,
    GRAY,
    BLACK
};

struct Object {
    Color Marked = WHITE;
    int Left = -1;
    int Right = -1;
    bool Present = true;
};

int numcycles = 10000000;
float heapprop = 0.85;
std::string infile = "../python/reg-graph-data/big-graph.txt";
std::vector<Object*> heap;
std::vector<int> rootSet;

void initHeap(size_t size) {
    heap.resize(size);
    for (size_t i = 0; i < size; i++) {
        heap[i] = new Object();
    }
}

void resetColors() {
    for (size_t i = 0; i < heap.size(); i++) {
        heap[i]->Marked = WHITE;
        heap[i]->Present = true;
    }
}

void initialMark() {
    for (int i : rootSet) {
        heap[i]->Marked = GRAY;
    }
}

void mark(int i) {
    Object *obj = heap[i];
    if (obj->Marked == BLACK)
        return;
    
    obj->Marked = GRAY;

    if (obj->Left != -1 && heap[obj->Left]->Marked != GRAY)
        mark(obj->Left);
    if (obj->Right != -1 && heap[obj->Right]->Marked != GRAY)
        mark(obj->Right);

    obj->Marked = BLACK;
}

void Mark() {
    for (int i : rootSet) {
        Object *obj = heap[i];
        if (obj->Marked == GRAY) {
            mark(i);
        }
    }
}

void sweep() {
    for (size_t i = 0; i < heap.size(); i++) {
        Object *obj = heap[i];
        if (obj->Marked == WHITE)
            obj->Present = false;
    }
}

void gcCycle() {
    resetColors();

    initialMark();

    Mark();

    sweep();
}

void printObjectStatus();

int main(int argc, char* argv[]) {
    std::ifstream input;
    input.open(infile);

    int n, u, v, w;
    input >> n;

    initHeap(size_t((float) n / heapprop));

    while (input >> u >> v >> w) {
        if (heap[u]->Left == -1)
            heap[u]->Left = v;
        else
            heap[u]->Right = v;
    }

    int rootsetsize = int(float(n) / 5.);

    int interval = n / rootsetsize;
    for (size_t i = 0; i < rootsetsize; i++) {
        rootSet.push_back(interval*i);
    }

    for (size_t i = 0; i < numcycles; i++) {
        gcCycle();
        // printf("After GC cycle %d:\n", i+1);
        // printObjectStatus();
    }

    input.close();
    return 0;
}

void printObjectStatus() {
    for (size_t i = 0; i < heap.size(); i++) {
        Object *obj = heap[i];
        if (!obj->Present)
            printf("Object %d: Freed\n", i);
        else if (obj->Marked == BLACK)
            printf("Object %d: Marked (Black)\n", i);
    }
}