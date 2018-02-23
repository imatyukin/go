// Упражнение 1.7.
// Вызов функции io.Copy(dst,src) выполняет чтение src и запись в dst. Воспользуйтесь ею вместо ioutil.ReadAll для
// копирования тела ответа в поток os.Stdout без необходимости выделения достаточно большого для хранения всего ответа
// буфера. Не забудьте проверить, не произошла ли ошибка при вызове io.Сору.

//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

//!-
