package go_reflect

import (
	"reflect"
	"testing"

	go_ref "github.com/goccy/go-reflect"
)

// AIM: go-reflectのパフォーマンスがいいことを確認する。
// パフォーマンスの違いが出る直接的な要因としては、reflect packageでは
// 内部でescapes(i)という処理で明示的に引数をheapに逃してるため。
// (わざわざheapに逃す理由はvalue.goのchanrecvのコメントに書いてあった。)
//
// また、直接コードには書いてないor一般的なはなしで、実験してる中で下記の知見を得た.
//
// 1.関数の引数にifaceを渡すこと自体が遅くなるというわけではなさそうだった。
//   (多分、)コンパイルされた状態では,ifaceが渡されたcallee側ではiface用にstackに領域が
//   が確保されていて(具象型のレイアウトではない)、恐らく違いはその部分だけ。
//   ここにおいて、速度に違いが出る部分は確かになさそう.
// 2.コンパイラ(GC)はコンパイル時にその変数のlifetimeを決定できない時には基本的にheapに逃すようにする.
//   (これが実行時にmemを確保しようとするのか、コンパイル時に確定されるのかはよくわからない)
//   逆に言えばglobal変数など、明らかにscope内の変数(そのscope内では明らかにallocateが発生しない変数)については
//   (当たり前だが)高速に処理ができる.これは、例えばa変数をglobal変数にしたら確認ができた.
//   ちなみに、この場合はなぜかGoのライブラリの方が早い結果になった。

func valueFromReflect(v interface{}) {
	reflect.ValueOf(v).Elem()
}

func valueFromGoReflect(v interface{}) {
	go_ref.ValueNoEscapeOf(v).Elem()
}

type aa struct {
	name string
	age  int
}

func BenchmarkGoPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := aa{}
		valueFromReflect(&a)
	}
}

func BenchmarkLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := aa{}
		valueFromGoReflect(&a)
	}
}
