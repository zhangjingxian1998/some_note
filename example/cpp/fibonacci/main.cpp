#include<iostream>
#include <sys/time.h>
#include"head/fibonacci.h"
using namespace std;

int main(){
    struct timeval t1,t2;
    double timeuse;

    int i;
    int result;
    cout<<"Please input the index of fibonacci:";
    cin>>i;
    gettimeofday(&t1,NULL);
    result = Fibonacci_recursion(i);
    gettimeofday(&t2,NULL);
    timeuse = (t2.tv_sec - t1.tv_sec) + (double)(t2.tv_usec - t1.tv_usec)/1000000.0;
    cout<<"time = "<<timeuse<<endl;  //输出时间（单位：ｓ）
    cout<<endl<<"the result is:"<<result<<endl;
    gettimeofday(&t1,NULL);
    result = Fibonacci_memory_save_start(i);
    gettimeofday(&t2,NULL);
    cout<<endl<<"the result is:"<<result<<endl;
    timeuse = (t2.tv_sec - t1.tv_sec) + (double)(t2.tv_usec - t1.tv_usec)/1000000.0;
    cout<<"time = "<<timeuse<<endl;  //输出时间（单位：ｓ）
    return 0;
}