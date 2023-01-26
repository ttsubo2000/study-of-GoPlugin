// プラグインとして利用する処理は main パッケージに属している必要がある

package main

import (
	"fmt"

	"github.com/ttsubo2000/study-of-GoPlugin/simple-go-plugin/lib/pkg/strs"
)

func Fn(message string) {
	fmt.Println(strs.Upper(message))
}
