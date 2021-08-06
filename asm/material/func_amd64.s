
/*
    basics
*/
TEXT ·add32(SB),$0-12   // MEMO: なぜか引数と返り値の間には謎のallignmentが入るみたい。引数は4byteの変数1つだけど、8byte分使うことになる
                        // 追記: 「GoのABIは戻り値の開始位置はワード長でアラインされるようです」ということらしい
    MOVL a+0(FP),AX     // 0(FP)は第一引数, 
    ADDL $2, AX
    MOVL AX, ret+8(FP)
    RET

TEXT ·add32_2(SB),$0-12 // MEMO: この場合は引数がちょうど8byteなので、直上にそのまま返り値の4byteが乗る.
    MOVL a+0(FP),AX     // 0(FP)は第1引数, 
    MOVL b+4(FP),DX     // 4(FP)は第2引数, 
    ADDL DX, AX
    MOVL AX, ret+8(FP)
    RET

TEXT ·add64(SB),$0-16   // MEMO: 引数:返り値が8byte:8byte
    MOVQ a+0(FP),AX     // 8byteのデータを扱う時はMOVQ命令を使う。
    ADDL $2, AX
    MOVQ AX, ret+8(FP)
    RET

/*
    slice
*/
TEXT ·sliceLen(SB),$0-32    // MEMO: sliceは「1. arrayへの先頭ptr(8byte), 2.len(8byte), 3. cap(8byte)」の24byteで表現される
    MOVQ b_len+8(FP),AX
    MOVQ AX, ret+24(FP)
    RET

TEXT ·sliceCap(SB),$0-32
    MOVQ b_cap+16(FP),AX
    MOVQ AX, ret+24(FP)
    RET

TEXT ·sliceBase(SB),$0-32
    MOVQ b+0(FP),AX
    MOVQ AX, ret+24(FP)
    RET

/*
    caller
*/
TEXT ·caller(SB),$24-16
    MOVL i+0(FP), AX
    ADDL $2, AX
    MOVL AX, ret1+8(FP)

        // prologue
        SUBQ $24, SP    // neg関数の引数と戻り値サイズ+BPレジスタの退避先を確保
        MOVQ BP, 16(SP) // 現在のBPレジスタをpush
        LEAQ 16(SP), BP // BPレジスタを新しいスタックに更新

        // contents
        MOVQ AX, (SP)   // MEMO: なんでここは(FP)出ないのだろう。
        CALL ·callee(SB)   // main.negを呼ぶ
        MOVL 8(SP), AX  // main.negの戻り値をAXレジスタに取り出す

        // epilogue
        MOVQ 16(SP), BP // 退避していたBPレジスタをpop
        ADDQ $24, SP    // スタックサイズを戻す

    MOVL AX, ret2+12(FP) // 2番目の戻り値として返す
    RET

/*
    <Layout>

    * stack flame
        * callee arg -> 8byte
        * callee ret -> 8byte
        * BP         -> 8byte
            => 24byte
    * stack size
        * caller arg -> 8byte
        * caller ret -> 16byte
            => 24byte
*/

TEXT ·caller2(SB),$24-24
    // get first arg and ret it.
    MOVQ i+0(FP),AX
    MOVQ AX, ret1+8(FP)

    // get second arg and set it AX
    MOVQ i+0(FP),AX

    // start prologue
    SUBQ $24,SP
    MOVQ BP,16(SP)
    LEAQ 16(SP),BP

    // contents
    MOVQ AX,(SP)
    CALL ·callee2(SB)
    // resultをAXに入れる
    MOVQ 8(SP),AX
    // epilogue
    MOVQ 16(SP),BP
    ADDQ $24,SP

    // ret caller
    MOVQ AX,ret2+16(FP)
    RET

