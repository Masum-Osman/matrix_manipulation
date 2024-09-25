// test_add.cpp
#include <iostream>

extern "C" {
    int add(int a, int b);
}

int main() {
    std::cout << "5 + 10 = " << add(5, 10) << std::endl;
    return 0;
}
