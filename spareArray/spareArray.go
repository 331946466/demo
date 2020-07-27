package spareArray

import (
	"mygo/testPoker"

	//"mygo/testPoker"
)

func ArrayPoker(BytePoker []byte)(GradeResult int,NameResult string,ArrayResult[][] int) {
	var j int = 0
	var h int = 0
	doubleArray := make([][]int, len(BytePoker))
	for k := 0; k < len(BytePoker); k++ {
		doubleArray[k] = make([]int, 3)
	}
//将牌存入稀疏数组
	for i := -1; i < len(BytePoker)-1; i++ {
		if i%2 == 0 {
			if BytePoker[i] == 65 {
				doubleArray[j][0] = 14
			}
			if BytePoker[i] == 75 {
				doubleArray[j][0] = 13
			}
			if BytePoker[i] == 81 {
				doubleArray[j][0] = 12
			}
			if BytePoker[i] == 74 {
				doubleArray[j][0] = 11
			}
			if BytePoker[i] == 84 {
				doubleArray[j][0] = 10
			}
			if BytePoker[i] == 57 {
				doubleArray[j][0] = 9
			}
			if BytePoker[i] == 56 {
				doubleArray[j][0] = 8
			}
			if BytePoker[i] == 55 {
				doubleArray[j][0] = 7
			}
			if BytePoker[i] == 54 {
				doubleArray[j][0] = 6
			}
			if BytePoker[i] == 53 {
				doubleArray[j][0] = 5
			}
			if BytePoker[i] == 52 {
				doubleArray[j][0] = 4
			}
			if BytePoker[i] == 51 {
				doubleArray[j][0] = 3
			}
			if BytePoker[i] == 50 {
				doubleArray[j][0] = 2
			}
			if BytePoker[i]==88{
				doubleArray[j][0]=1
			}
			if j < len(doubleArray) {
				j++
			}
		}
		if i%2 == 1 {
			if BytePoker[i] == 115 {
				doubleArray[h][1] = 1
			}
			if BytePoker[i] == 104 {
				doubleArray[h][1] = 2
			}
			if BytePoker[i] == 99 {
				doubleArray[h][1] = 3
			}
			if BytePoker[i] == 100 {
				doubleArray[h][1] = 4
			}
			if BytePoker[i]==110{
				doubleArray[h][1]=5
			}
			if h < len(doubleArray) {
				h++
			}
		}
	}
	//对牌进行排序
	for n := 0; n < len(doubleArray); n++ {
		for m := 0; m < len(doubleArray); m++ {
			if doubleArray[n][0] > doubleArray[m][0] {
				temp1 := doubleArray[n][0]
				temp2 := doubleArray[n][1]
				doubleArray[n][0] = doubleArray[m][0]
				doubleArray[n][1] = doubleArray[m][1]
				doubleArray[m][0] = temp1
				doubleArray[m][1] = temp2
			}
			if doubleArray[n][0] == doubleArray[m][0] {
				if doubleArray[n][1] > doubleArray[m][1] {
					temp1 := doubleArray[n][0]
					temp2 := doubleArray[n][1]
					doubleArray[n][0] = doubleArray[m][0]
					doubleArray[n][1] = doubleArray[m][1]
					doubleArray[m][0] = temp1
					doubleArray[m][1] = temp2
				}
			}
		}

	}
	Grade,Name:=testPoker.TestPoker(doubleArray)
	return Grade,Name,doubleArray
}

//}
