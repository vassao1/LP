#include <stdio.h>

int is_prime(int num) {
    int i;
    if (num == 2) {
        return 1;
    }
    if (num < 1) {
        return 0;
    }
    if (num % 2 == 0) {
        return 0;
    }
    for (i = num - 1; i > 1; i--) {
        if (num % i == 0) {
            return 0;
        }
    }
    return 1;
}

int main() {
    int num;

    printf("Insira um número: ");
    scanf("%d", &num);

    if (is_prime(num) == 1) {
        printf("%d é um número primo!", num);
    } else {
        printf("%d não é um número primo.", num);
    }

    return 0;
}