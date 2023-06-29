#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int min(int arr[3])
{ 
    int min = arr[0];
    int index;

    for (int x = 0; x<3; x++) {
        if (arr[x] < min) {
            min = arr[x];
            index = x;
        }
    }

    return index;
}

int main()
{
    FILE *fp;
    char line[10];
    int total = 0;
    int top3[3] = {0};
    int mini;

    fp = fopen("./c/input.txt", "r");

    if(fp == NULL) {
        perror("Error opening file");
        return(-1);
    }

    while ((fgets(line, sizeof line, fp)) != NULL ) {
        if (line[1] == '\0') {
            mini = min(top3);
            if (top3[mini] < total) {
                top3[mini] = total; 
            }
            total = 0;
        }
        total += atoi(line);
    }

    int final;
    for (int i = 0; i<3; i++) {
        final += top3[i];
    }
        

    printf("THE MAX CALORIES ARE: %d!\n", final);

    fclose(fp);

    return 0;
}
