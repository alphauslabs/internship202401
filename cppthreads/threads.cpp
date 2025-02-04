#include <thread>
#include <string>
#include <iostream>

using namespace std;

void task(string msg) {
    cout << "task starts: " << msg << endl;
    while (1) {}
}

int main() {
    std::thread t1(task, "hello1");
    std::thread t2(task, "hello2");
    std::thread t3(task, "hello3");
    std::thread t4(task, "hello4");
    t1.join();
    t2.join();
    t3.join();
    t4.join();
}
