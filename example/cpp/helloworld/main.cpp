#include"head/hello.h"
extern void Print_hello_extern();
int main(){
    Print_hello();
    Print_hello_extern();
    return 0;
}