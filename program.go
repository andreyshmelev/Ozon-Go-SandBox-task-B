package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var datasetsCount int
	fmt.Fscan(in, &datasetsCount)

	//	fmt.Println("datasetsCount + ", datasetsCount)

	for i := 0; i < datasetsCount; i++ {

		var howManyGoodsCount string

		fmt.Fscan(in, &howManyGoodsCount)

		//		fmt.Println("howManyGoodsCount + ", howManyGoodsCount)

		var indexes [100000]int

		var howManyGoodsCountInteger int
		howManyGoodsCountInteger, _ = strconv.Atoi(howManyGoodsCount)

		goodPrices := make([]int, howManyGoodsCountInteger)

		for ik := 0; ik < howManyGoodsCountInteger; ik++ {
			fmt.Fscan(in, &goodPrices[ik])
			indexes[goodPrices[ik]]++
		}
		//	fmt.Println("goodPrices : ", goodPrices)

		var total = 0

		for i := 0; i < 100000; i++ {
			var it = 0

			if indexes[i] > 0 {
				for {
					total = total + i
					indexes[i]--
					it++
					if it == 3 {
						it = 0
						total = total - i
					}
					if indexes[i] == 0 {
						break
					}
				}
			}
		}
		fmt.Println(total)
	}
}
