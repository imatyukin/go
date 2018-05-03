// Упражнение 3.13.
// Напишите объявления const для КВ, MB и так далее до YB настолько компактно, насколько сможете.
package main

import "fmt"

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB)
}
