//
//  main.c
//  ttt'
//
//  Created by 晓洋 on 2020/10/8.
//

#include <stdio.h>
#include <time.h>
int main ()
{
    time_t x;
    int hour,minute,second;
    struct tm * lt;//定义一个结构体
    //结构体定义在下面方式
    //struct tm {
    //int    tm_sec;        /* seconds after the minute [0-60] */
    //int    tm_min;        /* minutes after the hour [0-59] */
    //int    tm_hour;    /* hours since midnight [0-23] */
    //int    tm_mday;    /* day of the month [1-31] */
    //int    tm_mon;        /* months since January [0-11] */
    //int    tm_year;    /* years since 1900 */
//    int    tm_wday;    /* days since Sunday [0-6] */
//    int    tm_yday;    /* days since January 1 [0-365] */
//    int    tm_isdst;    /* Daylight Savings Time flag */
//    long    tm_gmtoff;    /* offset from UTC in seconds */
//    char    *tm_zone;    /* timezone abbreviation */
//};
    time (&x);
    //time_t time(time_t *time)  返回自纪元 Epoch（1970-01-01 00:00:00 UTC）起经过的时间，以秒为单位。
    lt = localtime (&x);
    //struct tm* localtime(const time_t *timer) 将time转换为tm
    //将结构体成员变量分别赋值给hour,minute,second
    hour = lt->tm_hour;
    minute = lt->tm_min;
    second = lt->tm_sec;
    printf( "%02d:%02d:%02d\n",hour,minute,second);
    return 0;
}
