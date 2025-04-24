#include <stdio.h>
int main() {
    int num, i, numorig = 1;

    printf("Insira um número: ");
    scanf("%d", &num);
    
    numorig = num;
    
    if (num < 0) {
        printf("Número inválido\n");
        return 1;
    }
    for (i = num - 1; i > 1; i--) {
        num = num * i;
    }
    printf("O fatorial de %d é: %d\n", numorig, num);
    return 0;
}