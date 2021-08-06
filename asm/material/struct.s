#include "go_asm.h"

TEXT ·addPoint(SB),$0-24
    MOVQ p+Point_X(FP),AX
    MOVQ p+Point_Y(FP),DX
    ADDQ DX,AX
    MOVQ AX,ret+16(FP)
    RET
