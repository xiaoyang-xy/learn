//
//  main.c
//  stdlib
//
//  Created by 晓洋 on 2020/9/21.
//
/*
tcpdump -i any -s 0 -w

*/
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <time.h>
#include <string.h>
#include <locale.h>
#include <setjmp.h>
int main(int argc, const char * argv[]) {
    // insert code here...
    double n, m;
    double pi = 3.1415926535;
    char buffer[256];
    fgets(buffer, 256, stdin);
    //double atof(const char *) 将字符串转换为double
    n = atof(buffer);
    //int atoi(const char *) string -> int
    int i = atoi(buffer);
    //long int atol(const char *)   string -> long int
    long int i2 = atol(buffer);
    //long long int atoll(const char *) string -> long long int
    long long int i3 = atoll(buffer);
    m = sin(n*pi/180);
    printf("%lf\n%lf\n%d\n%ld\n%lld\n",n,m,i,i2,i3);
    char szOrbits[] = "365.24    29.53 1.52";
    char *pEnd;
    char *pEnd1;
    /*double d1, d2;
    //double strtod(const char *, char **)  应该是后面如果为双重指针，将一个一个往后走直到空白字符停止，然后指针指向下一个非空白字符 若后面为NULL,
    d1 = strtod(szOrbits, &pEnd);
    d2 = strtod(pEnd, NULL);
    printf("%.2f\n",d1/d2); */
    float f1, f2, f3;
    f1 = strtof(szOrbits, &pEnd);
    f2 = strtof(pEnd, &pEnd1);
    f3 = strtof(pEnd1, NULL);
    //long double strtold(const char *, char **) return long double
    printf("%.2f,%.2f,%.2f\n",f1,f2,f3);
    char szNumbers[] = "2001 60c0c0 -1101110100110100100000 0x6fffff";
    long int li1, li2, li3, li4;
    //long int strtol(const char *__str, char **__endptr, int __base) return num(__base进制) 0为默认格式化
    //long long int strtoll(const char *__str, char **__endptr, int __base) return num(__base进制)
    //unsigned long int strtoul(const char *__str, char **__endptr, int __base) return num(__base进制)
    //unsigned long long int strtoul(const char *__str, char **__endptr, int __base) return num(__base进制)
    li1 = strtol (szNumbers,&pEnd,10);
    li2 = strtol (pEnd,&pEnd,16);
    li3 = strtol (pEnd,&pEnd,2);
    li4 = strtol (pEnd,NULL,0);
    printf ("The decimal equivalents are: %ld, %ld, %ld and %ld.\n", li1, li2, li3, li4);
    for (int i = 0; i < 10; i++) {
        //int rand() 10  return [0,10)
        int randnum = rand() % 100;
        //srand(unsigned int) set rand() = _rand() * i
        srand(i);
        printf("%d\n",randnum);
    }
    
    char *pData = (char *)calloc(10, sizeof(char));
    for (int i = 0; i < 3; i++) {
        pData[i] = 'c';
    }
    printf("%s\n",pData);
    char *pd;
    pd = (char *)malloc(100);
    pd = (char *)realloc(pd, 1);
    strcpy(pd,"mytest");
    printf("%s\t%p\n",pd,pd);
    strcat(pd, ".com");
    printf("%s\t%p\n",pd,pd);
    printf("%p\n",(char *)realloc(pd, 1));
    printf("%s\t%p\n",pd,pd);
    free(pd);
    free(pData);
    
    FILE *pFile;
    pFile = fopen("mytest.txt", "r");
    if (!pFile) {
        fputs("error open file\n", stderr);
        //中断操作 void abort(void)
        //abort();
    }
    //int abs(int i) return the Absolute value of the number i
    //div_t div(int, int)   return the div_t
    div_t diverresult;
    diverresult = div(38, 5);
    printf("%d %d\n",diverresult.quot,diverresult.rem);
    fclose(pFile);
    printf("%s\n",setlocale(LC_ALL, "C"));
    jmp_buf env;
    int val;
    //设置捕获点 longjmp就会报错
    val = setjmp(env);
    printf("val is %d\n",val);
    if (!val) {
        longjmp(env, 1);
    }
    return 0;
}
        
