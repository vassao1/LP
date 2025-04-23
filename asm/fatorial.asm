%INCLUDE "io.inc"

SECTION .data
num DD 6
fat DD 1

SECTION .text
global CMAIN
CMAIN:
    MOV ECX, [num]
    MOV EAX, 1

loop_fatorial:
    MUL ECX
    LOOP loop_fatorial
    MOV [fat], EAX
    PRINT_UDEC 4, [fat] 
    PRINT_CHAR 10

    MOV EAX, 0
    RET