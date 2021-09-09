package effectivego

import (
	"os"
	"io"
)

func Contents(filename string) (string, error){
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	// ここにクローズと書くことでファイルを処理後閉めることを忘れさせない。
	// PHPだと、file.open →　長い処理　→　file.close　と書けてしまい、あとでソースに修正を入れた後、closeを忘れるかもしれない。
	// Goであれば近くにcloseがあってわかりやすい。
	// 所感：f.close f.openを使う前にドキュメント読みにいくだろうからまず忘れない気がする・・・・
	defer f.Close()

	// 変数名　[]型
	var result []byte

	buf := make([]byte, 100)

	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if(err != nil){
			if err == io.EOF {
				break
			}
			return "", err
		}
	}

	// 文字列型に入れると配列を文字にできる・・？
	return string(result), err
}