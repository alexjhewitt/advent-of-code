
#include <stdio.h>

int main() {
    char ch;
    int floor = 0;
    int firstNegative = -1;
    int count = 0;

    FILE *file = fopen("parentheses.txt", "r");

    if (file == NULL) {
        printf("Error: Could not open file.\n");
         return 1;
    }

    while ((ch = fgetc(file)) != EOF) {
        count += 1;
        if (ch == '(') {
            floor += 1;
        } else if (ch == ')') {
            floor -= 1;
        }
        if (floor < 0 && firstNegative == -1) {
            firstNegative = count;
        }

    }

    // Close the file
    fclose(file);
    printf("floor taken: %d\n", floor);
    printf("first entered basement: %d", firstNegative);

    return 0;
}
