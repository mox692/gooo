PKG=main
TARG=a.out
AR=main.a

all: $(TARG)

$(TARG): $(AR)
	go tool link $(AR)

$(AR): main.o func_amd64.o
	go tool pack c $@ $^

main.o: main.go symabis
	go tool compile -symabis symabis -p $(PKG) $<

symabis: *.s
	go tool asm -gensymabis -o $@ $^

func_amd64.o: func_amd64.s go_asm.h
	go tool asm -p $(PKG) $<

go_asm.h: *.go
	go tool compile -asmhdr $@ $^

clean:;
	rm *.o *.s $(TARG) $(AR) go_asm.h symabis
