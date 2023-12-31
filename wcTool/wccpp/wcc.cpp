#include<iostream>
#include<string>
#include<unordered_map>
using namespace std;

class Arguments{

};

unordered_map<string, Arguments*> arg_map; 

int main(int argc, char* argv[])
{
    int i =1;
    while(i<=argc)
    {
        cout<<"the argument no is : "<<i<<" the arg is : "<<argv[i-1]<<endl;
        i++;
    }

    return 0;
}