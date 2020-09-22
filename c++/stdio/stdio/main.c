//
//  main.c
//  stdio
//
//  Created by 晓洋 on 2020/9/18.
//

#include <stdio.h>

int main(int argc, const char * argv[]) {
    // insert code here...
    //int remove(const char *) if success return 0 else return !0
    if (remove("/Volumes/Data/coding/c++/stdio/stdio/test.txt") != 0) {
        perror("error delete file!\n");
    }
    else
        printf("success delete file!\n");
    //int rename(const char *__old, const char *__new) if rename success return 0 else return !0
    int result;
    char oldname[] = "/Volumes/Data/coding/c++/stdio/stdio/oldname.txt";
    char newname[] = "/Volumes/Data/coding/c++/stdio/stdio/newname.txt";
    result = rename(oldname, newname);
    if (result == 0) {
        printf("FIle rename success\n");
    }
    else {
        perror("File rename error\n");
    }
    //int fprintf(FILE *restrict, const char *restrict, ...)
    FILE *pFile;
    int n = 100;
    char buffer[27];
    char name[100] = "woshinidie";
    pFile = fopen("/Volumes/Data/coding/c++/stdio/stdio/newname.txt", "w+");
    printf("%d",fprintf(pFile, "Name = %s, num = %d, num16 = %x",name,n,n));
    //void rewind(FILE *)   将文件指针指到开头，同时清除和流相关的错误和EOF标记
    rewind(pFile);
    //size_t fread(void *restrict __ptr, size_t __size, size_t __nitems, FILE *restrict __stream) return the length of str
    fread(buffer, 1, 26, pFile);
    buffer[26] = '\0';
    puts(buffer);
    char str[] = "ddd";
    //size_t fwrite(const void *restrict __ptr, size_t __size, size_t __nitems, FILE *restrict __stream)    return
    fwrite(str, 1, sizeof(str) - 1, pFile);
    fclose(pFile);
    return 0;
}
