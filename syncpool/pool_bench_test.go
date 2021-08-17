package main

import (
	"fmt"
	"sync"
	"testing"
)

// poolの使用例. intのスライスを使い回す場合に、都度makeする方法と、一度poo;lを作っておいてそれを使い回す物の比較。
// この例では[]intのpool1つしか用いていないが、pool自体を複数作成することも可能。
// また、並行処理の中で共通のpoolを使用することも可能(その場合はNew関数内のatmic性に注意する)
var counts = []int{1, 100, 1000, 100000}

func BenchmarkWithoutPool(b *testing.B) {
	for _, count := range counts {
		b.Run(fmt.Sprintf("count: %d", count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = make([]int, count)
			}
		})
	}
}

func BenchmarkWithPoolVal(b *testing.B) {
	for _, count := range counts {
		b.Run(fmt.Sprintf("count: %d", count), func(b *testing.B) {
			pool := sync.Pool{
				New: func() interface{} {
					// MEMO: ここの要素数は、再allocが発生しない様によく考える必要があると思う
					return make([]int, count)
				},
			}
			for i := 0; i < b.N; i++ {
				// MEMO: poolから, int * count分の領域を取ってくる(この取得コストはかなり小さい)
				val, ok := pool.Get().([]int)
				if !ok {
					b.Errorf("get err\n")
				}
				// sliceを用いた何かしらの処理
				val[count-1] = i
				// sliceの仕様が終わったら再度poolに戻す
				pool.Put(val)
			}
		})
	}
}

func BenchmarkWithPoolPtr(b *testing.B) {
	for _, count := range counts {
		b.Run(fmt.Sprintf("count: %d", count), func(b *testing.B) {
			pool := sync.Pool{
				New: func() interface{} {
					return new([]int)
				},
			}
			for i := 0; i < b.N; i++ {
				val, ok := pool.Get().(*[]int)
				if !ok {
					b.Errorf("get err\n")
				}
				pool.Put(val)
			}
		})
	}
}

// 実行中のzero allocを実現できたcode.(Newで&sを返しているのがポイント)
func BenchmarkWithPoolPtrEff(b *testing.B) {
	for _, count := range counts {
		b.Run(fmt.Sprintf("count: %d", count), func(b *testing.B) {
			pool := sync.Pool{
				New: func() interface{} {
					// MEMO: ここの要素数は、再allocが発生しない様によく考える必要があると思う
					s := make([]int, count)
					return &s
				},
			}
			for i := 0; i < b.N; i++ {
				// MEMO: poolから, int * count分の領域を取ってくる(この取得コストはかなり小さい)
				val, ok := pool.Get().(*[]int)
				if !ok {
					b.Errorf("get err\n")
				}
				(*val)[count-1] = 3
				pool.Put(val)
			}
		})
	}
}
