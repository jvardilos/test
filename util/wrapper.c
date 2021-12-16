#include <stdio.h>
#include <string.h>

char * rubber_(char *, int *);

char * Run(char input[]) {

    int len = strlen(input);
    char * output = rubber_(input, &len);

    return output;
}