package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {
	// プラグイン（soファイル）を読み込む
	p, err := plugin.Open("calc/calc.so")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// 変数Pvのシンボルを取得
	pv, err := p.Lookup("Pv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// 関数Pow2()のシンボルを取得
	pow2, err := p.Lookup("Pow2")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// pvを2のべき乗した計算結果を出力
	*pv.(*int) = 3
	p2 := pow2.(func() int)()
	fmt.Printf("pow2: %d * %d = %d\n", *pv.(*int), *pv.(*int), p2)

	// 関数AddSub()のシンボルを取得
	addsub, err := p.Lookup("AddSub")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// 足し算と引き算の計算結果を出力
	x, y := 5, 2
	add, sub := addsub.(func(int, int) (int, int))(x, y)
	fmt.Printf("add: %d + %d = %d\n", x, y, add)
	fmt.Printf("sub: %d - %d = %d\n", x, y, sub)
}
