"".add STEXT nosplit size=20 args=0x10 locals=0x0
	0x0000 00000 (main.go:4)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $0-16	;; stackframeはゼロ、ただstackは16byte引き下げてね。
	0x0000 00000 (main.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:5)	MOVL	"".secondArg+12(SP), AX	;; SPから12byte上のアドレスを示す(多分、CALLした段階で↑の16byte分、SPが下げられてる).ちなみにこの記法は名前付き引数参照って書かれてたやつ.
	0x0004 00004 (main.go:5)	MOVL	"".firstArg+8(SP), CX	;; firstArgが後に取り出されう
	0x0008 00008 (main.go:5)	ADDL	CX, AX
	0x000a 00010 (main.go:5)	MOVL	AX, "".~r2+16(SP)		;; 返り値1
	0x000e 00014 (main.go:5)	MOVB	$1, "".~r3+20(SP)		;; 返り値2(即値)
	0x0013 00019 (main.go:5)	RET								;; SP+0 ~ SP+8の区間って何に使われてたんや???? 追記：最初の引数 a は 0(SP) ではなく 8(SP) に配置されます。 これは caller が CALL 命令のときにリターン先のアドレスを 0(SP) に格納しているからです。
																;; RETはほとんどの場合、0(SP) に格納したリターン先のアドレスをPOPしてそこにジャンプします。
	0x0000 8b 44 24 0c 8b 4c 24 08 01 c8 89 44 24 10 c6 44  .D$..L$....D$..D
	0x0010 24 14 01 c3                                      $...
"".main STEXT size=66 args=0x0 locals=0x18
	0x0000 00000 (main.go:8)	TEXT	"".main(SB), ABIInternal, $24-0	;; 今度は24byteのstack flameをとり、stack自体はしたに伸ばさないらしい(引数を受け取らず、返り値も返さない)。
	0x0000 00000 (main.go:8)	MOVQ	(TLS), CX	;; TLSとは？？
	0x0009 00009 (main.go:8)	CMPQ	SP, 16(CX)  ;; SPとCXから16先の値を比較してる
	0x000d 00013 (main.go:8)	PCDATA	$0, $-2		;; ?
	0x000d 00013 (main.go:8)	JLS	58				;; ?
	0x000f 00015 (main.go:8)	PCDATA	$0, $-1		;; ?
	0x000f 00015 (main.go:8)	SUBQ	$24, SP		;; ここからの3行、prologueな感じ
	0x0013 00019 (main.go:8)	MOVQ	BP, 16(SP)	;; BPがさす値(恐らくこれもアドレス)を16(SP)に入れる！ 8 bytes (16(SP)-24(SP)) は フレームポインタBP の現在の値を保存するのに使われます。これによってスタックの巻き戻し(呼び出し下の関数を辿ること)が可能になりデバッグ時に便利です。
	0x0018 00024 (main.go:8)	LEAQ	16(SP), BP  ;; BPにSP+16のアドレスを入れる
	;; 上側の8byteはBPのために使われて、下位の16byteはaddのために使われそう
	0x001d 00029 (main.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:8)	MOVQ	$137438953482, AX	;; まってこれ引数の値らしいwwwwww int32を引数にしてるからこういう(8byteに2つの引数を入れるような)書き方になってるのか
	0x0027 00039 (main.go:8)	MOVQ	AX, (SP)			;; SPをとんでもない値に変更してるように見えるgあ。。
	0x002b 00043 (main.go:8)	PCDATA	$1, $0				;; ? GCに対する命令らしくて、そこまで気にしなくても良さげ
	0x002b 00043 (main.go:8)	CALL	"".add(SB)			;; addを呼ぶ
	0x0030 00048 (main.go:8)	MOVQ	16(SP), BP			;; epilogue?
	0x0035 00053 (main.go:8)	ADDQ	$24, SP
	0x0039 00057 (main.go:8)	RET							;; main自体はここで終了してそう。
	0x003a 00058 (main.go:8)	NOP
	0x003a 00058 (main.go:8)	PCDATA	$1, $-1
	0x003a 00058 (main.go:8)	PCDATA	$0, $-2
	0x003a 00058 (main.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x003f 00063 (main.go:8)	PCDATA	$0, $-1
	0x003f 00063 (main.go:8)	NOP
	0x0040 00064 (main.go:8)	JMP	0
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
