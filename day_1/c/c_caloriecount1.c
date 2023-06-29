#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main()
{
    FILE *fp;
    char line[10];
    int total = 0;
    int max = 0;

    fp = fopen("./c/input.txt", "r");

    if(fp == NULL) {
        perror("Error opening file");
        return(-1);
    }

    while ((fgets(line, sizeof line, fp)) != NULL ) {
        if (line[1] == '\0') {
            if (max < total) {
                max = total;
            }
            total = 0;
        }
        total += atoi(line);
    }

    printf("THE MAX CALORIES ARE: %d!\n", max);

    fclose(fp);

    return 0;
}
