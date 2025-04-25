#include <stdio.h>
#include <string.h>

int main() {
    char mainstr[100];
    printf("Digite uma string: ");
    fgets(mainstr, sizeof(mainstr), stdin);
    mainstr[strcspn(mainstr, "\n")] = '\0';

    char substr[50];
    printf("Digite a substring: ");
    fgets(substr, sizeof(substr), stdin);
    substr[strcspn(substr, "\n")] = '\0';

    int i, contador = 0;

    for (i = 0; i <= strlen(mainstr) - strlen(substr); i++) {
        if (mainstr[i] == substr[0]) {
            int j;
            for (j = 0; j < strlen(substr); j++) {
                if (mainstr[i + j] != substr[j]) {
                    break;
                    }
                }
            if (j == strlen(substr)) {
                contador++;
            }
        }
    }
    printf("A substring %s aparece %d vezes na string %s.", substr, contador, mainstr);
    return 0;
}