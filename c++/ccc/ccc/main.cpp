//
//  main.cpp
//  ccc
//
//  Created by 晓洋 on 2020/10/26.
//

#include <iostream>

int main(int argc, const char * argv[]) {
    // insert code here...
    for(int i = 1;i < 100;i++)
    {
        int a = 466 * i;
        if(a % 100 == 4)
        {
            std::cout<<i<<std::endl;
        }
    }
    return 0;
}
