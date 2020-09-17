//
//  main.c
//  test
//
//  Created by 晓洋 on 2020/9/16.
//
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <assert.h>

int main(int argc, const char * argv[]) {
    // insert code here...
    char str[] = "5g. \t";
    int i = 0;
    //assert(i!=0); 调试使用，传入bool值
    //int isalnum(int c)    将传入的c强制转换为int or EOF   判断是否为数字或者字母
    while (isalnum(str[i])) {
        ++i;
    }
    printf("%d\n",i);
    //int isalpha(int c)    将传入的c强制转换为int or EOF    判断是否为字母
    i = 0;
    while(str[i]){
        if (isalpha(str[i])) {
            printf("%c\t",str[i]);
        } else {
            printf("%c\t",str[i]);
        }
        ++i;
    }
    //int isblank(int c)    检查的字符转换为int or EOF，\t or 空格为非0
    i = 0;
    char c;
    while (str[i]) {
        c = str[i];
        if (isblank(c)) {
            printf("\\tor ");
        }
        ++i;
    }
    //int iscntrl(int c)    检查字符是否为控制字符
    //int isdigit(int c)    检查字符是否为十进制数字
    //int isgraph(int c)    检查字符是否具有图形表示 字母数字和标点符号
    //int islower(int c)    检查字符是否为小写字母
    //int isprint(int c)    检查字符是否可打印 和iscntrl相反
    //int ispunct(int c)    检查字符是否为标点符号! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~
    //int isspace(int c)    检查字符是否为空格   ' ' 0x20 空格 '\t' 0x09 TAB '\n' 0x0a LF '\v' 0x0b VT '\F' 0x0c FF '\r' 0x0d CR
    //int isupper(int c)    检查字符是否为大写字母
    //int isxdigit(int c)   检查字符是否为十六进制数字 0 1 2 3 4 5 6 7 8 9 a b c d e f A B C D E F
    //int tolower(int _c)   将大写字母转换为小写
    //int toupper(int _c)   将小写字母转换为大写
    printf("%d",toascii(str[0]));
    FILE *pFile;
    pFile = fopen("/Volumes/Data/coding/c++/test/test/test.txt", "w+");
    fprintf(pFile, "test fprintf");
    fputs("test fputs", pFile);
    fclose(pFile);
    printf("\n");
    return 0;
}
