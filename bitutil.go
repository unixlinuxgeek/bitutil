// Упражнение 4.1:
//
// Напишите функцию, которая подсчитывает количество битов,
// различных в двух дайджестах SHA256 (см. PopCount в разделе 2.6.2).

package bitutil

import "fmt"

//goland:noinspection ALL
func Diff(x, y [32]byte) int {
	d := 0
	for a := 0; a < 32; a++ {
		if x[a] != y[a] {
			d++
			//fmt.Printf("%d: x[a]: %b\t", a, x[a])
		}
	}
	fmt.Printf("Количество различающихся битов: %d\n", d)
	return d
}
