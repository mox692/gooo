"".add STEXT nosplit size=25 args=0x10 locals=0x0
	0x0000 00000 (main.go:4)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (main.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:5)	MOVL	"".a+8(SP), AX
	0x0004 00004 (main.go:5)	LEAL	(AX)(AX*2), CX
	0x0007 00007 (main.go:6)	MOVL	"".b+12(SP), DX
	0x000b 00011 (main.go:6)	ADDL	DX, AX
	0x000d 00013 (main.go:6)	ADDL	CX, AX
	0x000f 00015 (main.go:6)	MOVL	AX, "".~r2+16(SP)
	0x0013 00019 (main.go:6)	MOVB	$1, "".~r3+20(SP)
	0x0018 00024 (main.go:6)	RET
	0x0000 8b 44 24 08 8d 0c 40 8b 54 24 0c 01 d0 01 c8 89  .D$...@.T$......
	0x0010 44 24 10 c6 44 24 14 01 c3                       D$..D$...
"".main STEXT size=66 args=0x0 locals=0x18
	0x0000 00000 (main.go:9)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (main.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:9)	PCDATA	$0, $-2
	0x000d 00013 (main.go:9)	JLS	58
	0x000f 00015 (main.go:9)	PCDATA	$0, $-1
	0x000f 00015 (main.go:9)	SUBQ	$24, SP
	0x0013 00019 (main.go:9)	MOVQ	BP, 16(SP)
	0x0018 00024 (main.go:9)	LEAQ	16(SP), BP
	0x001d 00029 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	MOVQ	$137438953482, AX
	0x0027 00039 (main.go:9)	MOVQ	AX, (SP)
	0x002b 00043 (main.go:9)	PCDATA	$1, $0
	0x002b 00043 (main.go:9)	CALL	"".add(SB)
	0x0030 00048 (main.go:9)	MOVQ	16(SP), BP
	0x0035 00053 (main.go:9)	ADDQ	$24, SP
	0x0039 00057 (main.go:9)	RET
	0x003a 00058 (main.go:9)	NOP
	0x003a 00058 (main.go:9)	PCDATA	$1, $-1
	0x003a 00058 (main.go:9)	PCDATA	$0, $-2
	0x003a 00058 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x003f 00063 (main.go:9)	PCDATA	$0, $-1
	0x003f 00063 (main.go:9)	NOP
	0x0040 00064 (main.go:9)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 2b 48  dH..%....H;a.v+H
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 b8 0a  ...H.l$.H.l$.H..
	0x0020 00 00 00 20 00 00 00 48 89 04 24 e8 00 00 00 00  ... ...H..$.....
	0x0030 48 8b 6c 24 10 48 83 c4 18 c3 e8 00 00 00 00 90  H.l$.H..........
	0x0040 eb be                                            ..
	rel 5+4 t=17 TLS+0
	rel 44+4 t=8 "".add+0
	rel 59+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
