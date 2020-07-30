package testPoker

import (
	"fmt"
	"mygo/model"
)

//测试牌型
func TestPoker(doubleArray [][]int) (GradePoker int,NamePoker string){
	fmt.Println(doubleArray)
	ShunZi := ShunZi(doubleArray)
	Same := SameColor(doubleArray)
	Consecutive := Consecutive(doubleArray)
	AlonePoker := AlonePoker(doubleArray)
	Grade,Name :=model.GradePoker(ShunZi, Same, Consecutive, AlonePoker, doubleArray)
	return Grade,Name
}

//顺子
func ShunZi(doubleArray [][]int) (b string) {
	for i := 0; i < len(doubleArray)/2; i++ {
		if doubleArray[i][0] == doubleArray[i+1][0]+1 && doubleArray[i+1][0] == doubleArray[i+2][0]+1 && doubleArray[i+2][0] == doubleArray[i+3][0]+1 && doubleArray[i+3][0] == doubleArray[i+4][0]+1 && doubleArray[i][1] != 0 {
			b = "顺子"
		}
	}
	return b
}

//同花
func SameColor(doubleArray [][]int) (b string) {
	for i := 0; i < len(doubleArray)/2; i++ {
		var j int = 0
		temp := doubleArray[i][1]
		for k := i; k < len(doubleArray)/2; k++ {
			if temp == doubleArray[k][1] && temp != 0 {
				j++
			}
			if j >= 5 {
				b = "同花"
			}
		}
	}
	return b
}

//连对、三条、四条
func Consecutive(doubleArray [][]int) (b string) {
	for i := 0; i < len(doubleArray)/2; i++ {
		if doubleArray[i][0] == doubleArray[i+1][0] && doubleArray[i][0] != 0 {
			b = "对子"
			if doubleArray[i][0] == doubleArray[i+2][0] {
				b = "三条"
				if doubleArray[i][0] == doubleArray[i+3][0] {
					b = "四条"
					break
				}
				if doubleArray[i+3][0] == doubleArray[i+4][0] {
					b = "满堂彩"
					break
				}
			}
			if doubleArray[i+2][0] == doubleArray[i+4][0] && doubleArray[i+2][0] != 0 {
				b = "满堂彩"
				break
			}
		}
	}
	return b
}

//散子
func AlonePoker(doubleArr [][]int) (b string) {
	Same := SameColor(doubleArr)
	var j int
	for i := 0; i < len(doubleArr)/2; i++ {
		if doubleArr[i][0] != doubleArr[i+1][0] && doubleArr[i][1] != doubleArr[i+1][1] && Same != "顺子" && doubleArr[i][0] != 0 {
			j++
		}
		if j >= len(doubleArr)/2 {
			b = "散子"
		}
	}
	return b
}
