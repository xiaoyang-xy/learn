//
//  main.c
//  time
//
//  Created by 晓洋 on 2020/9/17.
//

#include <stdio.h>
#include <time.h>

int main(int argc, const char * argv[]) {
    // insert code here...
    //time_t   int64 long
    time_t curtime;
    //clock_t   unsigned long   clock()  return 处理器时钟所使用的时间。为了获取 CPU 所使用的秒数，您需要除以 CLOCKS_PER_SEC
    clock_t t = clock();
    printf("%ld\n",(double)t/CLOCKS_PER_SEC);
    //time_t time(time_t *time)  自纪元 Epoch（1970-01-01 00:00:00 UTC）起经过的时间，以秒为单位。
    time(&curtime);
    printf("%ld\n",curtime);
    //char *ctime(const time_t *)  返回一个本地时间Thu Sep 17 13:45:10 2020
    printf("%s\n",ctime(&curtime));
    t = clock() - t;
    //CLOCKS_PER_SEC    表示每秒的处理器时钟个数。
    printf("%lf\n",(double)t/CLOCKS_PER_SEC);
    //double difftime(end, beginning) 返回end到beginning之间的秒数
    //time_t mktime(struct tm *)    将tm结构体转换为time_t
    time_t now;
    struct tm newyear;
    double seconds;
    time(&now);
    //struct tm* localtime(const time_t *timer) 将time转换为tm
    newyear = *localtime(&now);
    newyear.tm_hour = 0;newyear.tm_min = 0;newyear.tm_sec = 0;newyear.tm_mday = 1;newyear.tm_mon = 0;
    seconds = difftime(now, mktime(&newyear));
    printf("%lf\n",seconds);
    //char *asctime(const struct tm * timeptr)  将tm转换为char *
    newyear = *localtime(&now);
    printf("%s",asctime(&newyear));
    //struct tm* gmtime(const time_t *)  utc时间
    printf("%s",asctime(gmtime(&now)));
    //size_t strftime(char *restrict, SIZE_T, const char *restrict, const struct tm *restrict)
    //将tm格式化输出传给buffer  返回buffer长度
    time_t rawtime;
    struct tm *timeinfo;
    char buffer[80];
    time(&rawtime);
    timeinfo = localtime(&rawtime);
    strftime(buffer, 80, "Now it's %I:%M%p.", timeinfo);
    puts(buffer);
    return 0;
}
