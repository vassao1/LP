#include <stdio.h>

int main() {
    float fahrenheit, celsius;

    printf("Insira a temperatura em Fahrenheit: ");
    scanf("%f", &fahrenheit);
    celsius = ((fahrenheit - 32) * 5) / 9;
    printf("A temperatura em Celsius é: %.2f", celsius);

    return 0;
}