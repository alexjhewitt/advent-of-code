#include <stdio.h>

int main(void) {
    FILE *f = fopen("./parentheses.txt", "r");
    char c;

    if (NULL == f) {
        printf("file can't be opened\n");
        return 1;
    }

    do {
        c = getc(f);
        printf("%c", c);
    }while (c != EOF);

    fclose(f);
    return 0;
}