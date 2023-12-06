#include"../head/fibonacci.h"
using namespace std;
int Fibonacci_recursion(int i){
    if(i==1){
        return 1;
    }
    if(i==2){
        return 1;
    }
    return Fibonacci_recursion(i-1) + Fibonacci_recursion(i-2);
}
int Fibonacci_memory_save(int i, int *list){
    if(i==1){
        list[i-1] = 1;
        return 1;
    }
    if(i==2){
        list[i-1] = 1;
        return 1;
    }
    if (list[i-1] != 0){
        return list[i-1];
    }
    int result;
    int result1 = Fibonacci_memory_save(i-1, list);
    int result2 = Fibonacci_memory_save(i-2, list);
    result = result1 + result2;
    list[i-1] = result;
    return result;
}
int Fibonacci_memory_save_start(int i){
    int list[i]={0};
    if(i==1){
        list[i-1]=1;
        return 1;
    }
    if(i==2){
        list[i-1]=1;
        return 1;
    }
    
    if (list[i-1] != 0){
        return list[i-1];
    }
    int result;
    int result1 = Fibonacci_memory_save(i-1, list);
    int result2 = Fibonacci_memory_save(i-2, list);
    result = result1 + result2;
    list[i-1] = result;
    
    return result;
}
