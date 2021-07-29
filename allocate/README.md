* go test -bench . -benchmem -gcflags "-m -l"
* go build -gcflags '-m' ./*.go

* 全体的な傾向(雑)として「関数内で確保した変数の参照を、その関数が返してしまうと(その関数がreturnしても参照されうるとコンパイラが判断するため)、ヒープに確保されてしまう.」