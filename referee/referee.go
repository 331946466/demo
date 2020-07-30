package referee

import "fmt"

func Referee(Grade int, Poker string, doubleArray [][]int, Player string, Grade2 int, Poker2 string, doubleArray2 [][]int, Player2 string) {
	if Grade > Grade2 {
		fmt.Println("胜利者是:", Player)
	}
	if Grade < Grade2 {
		fmt.Println("胜利者是:", Player2)
	}

	if Grade2 == Grade {
		//同为皇家同花时
		if Grade == 10 {
			fmt.Println("平局————", Poker)
		}
		//同为同花顺时
		if Grade == 9 {
			if doubleArray[0][0] > doubleArray2[0][0] && doubleArray[0][1] > doubleArray2[0][1] {
				fmt.Println("胜利者是:", Player)
			}
			if doubleArray[0][0] < doubleArray2[0][0] && doubleArray[0][1] < doubleArray2[0][1] {
				fmt.Println("胜利者是:", Player2)
			}
			if doubleArray[0][0] == doubleArray2[0][0] && doubleArray[0][1] == doubleArray2[0][1] {
				fmt.Println("平局")
			}
		}
		//同为四条时
		if Grade == 8 {
			var FirstPoker, FirstPoker2, AlonePoker, AlonePoker2 int
			for i := 0; i < len(doubleArray)/2; i++ {
				if doubleArray[i][0] == doubleArray[i+3][0] && doubleArray[i][0] != 0 {
					FirstPoker = doubleArray[i][0]
					for j := 0; j < len(doubleArray)/2; j++ {
						for k := 0; k < len(doubleArray)/2; k++ {
							if doubleArray[j][0] != doubleArray[i][0] && doubleArray[j][0] != 0 && doubleArray[j][0] > doubleArray[k][0] {
								AlonePoker = doubleArray[j][0]
							}
						}
					}
				}
				if doubleArray2[i][0] == doubleArray2[i+3][0] && doubleArray2[i][0] != 0 {
					FirstPoker2 = doubleArray2[i][0]
					for j := 0; j < len(doubleArray)/2; j++ {
						for k := 0; k < len(doubleArray)/2; k++ {
							if doubleArray2[j][0] != doubleArray[i][0] && doubleArray2[j][0] != 0 && doubleArray2[j][0] > doubleArray2[k][0] {
								AlonePoker2 = doubleArray2[j][0]
							}
						}
					}
				}
				if FirstPoker > FirstPoker2 {
					fmt.Println("胜利者是：", Player)
				}
				if FirstPoker < FirstPoker2 {
					fmt.Println("胜利者是:", Player2)
				}
				if FirstPoker == FirstPoker2 {
					if AlonePoker > AlonePoker2 {
						fmt.Println("胜利者是：", Player)
					}
					if AlonePoker < AlonePoker2 {
						fmt.Println("胜利者是：", Player2)
					}
					if AlonePoker == AlonePoker2 {
						fmt.Println("平局————")
					}
				}
			}
		}
		//同为满堂彩
		/*if Grade==7{
			for
					}
				}
		}*/
	}
}