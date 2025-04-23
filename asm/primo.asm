%INCLUDE "io.inc"

SECTION .data
num DD 2
primo DD 1

SECTION .text
global CMAIN
CMAIN:
    MOV EAX, [num]
    CMP EAX, 2
    JL nprimo
    JE sprimo
    MOV ECX, 2

primoloop:
    MOV EBX, [num]
    MOV EDX, 0
    DIV ECX

    CMP EDX, 0
    JE nprimo

    MOV EAX, ECX
    ADD EAX, 1
    MOV ECX, EAX
    MOV EAX, [num]
    CMP ECX, EAX
    JL primoloop

    JMP sprimo

nprimo:
    MOV [primo], DWORD 0

sprimo:
    PRINT_UDEC 4, [primo]
    PRINT_CHAR 10

    MOV EAX, 0
    RET