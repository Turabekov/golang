package main

import "fmt"

// 6. Uch xonali son berilgan. Uni o’nliklar xonasidagi raqam bilan birliklar xonasidagi raqamni almashtirishdan xosil bo’lgan sonni toping.
func main() {
	var num int = 456

	first := num / 100
	second := num % 100 / 10
	third := num % 100 % 10

	fmt.Printf("%d%d%d\n", first, third, second)

}
