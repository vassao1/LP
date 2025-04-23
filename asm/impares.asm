%INCLUDE "io.inc"

SECTION .data
numtoprint DD 9
odd DD 1

SECTION .text
global CMAIN
CMAIN:
    MOV EAX, [numtoprint]
    MOV EBX, 0
    MOV ECX, [odd]

looptoprint:
    PRINT_UDEC 4, ECX
    PRINT_CHAR 10 

    ADD ECX, 2
    MOV EAX, EBX
    ADD EAX, 1
    MOV EBX, EAX
    CMP EBX, [numtoprint]
    JL looptoprint

    MOV EAX, 0
    RET