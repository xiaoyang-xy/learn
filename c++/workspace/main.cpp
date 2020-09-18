#include <iostream>
#include <climits>
int main() {
    std::cout << INT_MAX << "\t" << INT_MIN << std::endl;
    std::cout << CHAR_BIT << "\t" << std::endl;
    std::cout << SCHAR_MAX << "\t" << SCHAR_MIN << std::endl;
    std::cout << UCHAR_MAX << "\t" << CHAR_MAX << "\t"  << CHAR_MIN << std::endl;
    std::cout << MB_LEN_MAX << "\t" << SHRT_MIN << "\t" << SHRT_MAX << std::endl;
    std::cout << USHRT_MAX << "\t" << UINT_MAX << std::endl;
    std::cout << LONG_MAX << "\t" << LONG_MIN << std::endl;
    std::cout << ULLONG_MAX << "\t" << ULONG_MAX << std::endl;
    std::cout << LLONG_MAX << "\t" << LLONG_MIN << std::endl;
    return 0;
}
