#include <stdio.h>

int main() {
    int num, i, cont;

    printf("Insira quantos ímpares você quer que sejam impressos: ");
    scanf("%d", &num);

    for (i = 1, cont = 1; cont <= num; i++) {
        if (i % 2 != 0) {
            printf("%d ", i);
            cont++;
        }
    }

    return 0;
}