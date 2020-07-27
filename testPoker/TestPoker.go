package testPoker

import (
	"fmt"
	"mygo/model"
)
//判断牌型
func TestPoker(doubleArray [][]int, a int) {
	fmt.Println(doubleArray)
	ShunZi := ShunZi(doubleArray, a)
	Same := SameColor(doubleArray, a)
	Consecutive:=Consecutive(doubleArray,a)
	AlonePoker:=AlonePoker(doubleArray,a)
	MaxPoker:=doubleArray[0][0]
	model.GradePoker(ShunZi,Same,Consecutive,AlonePoker,MaxPoker)
}
//顺子
func ShunZi(doubleArray [][]int, a int) (b string) {
	var j int=0
	for i := 0; i < a-2; i++ {
		if doubleArray[i][0] == doubleArray[i+1][0]+1 {
			j++
			if j>=5{
				b="顺子"
			}
			}
		}

	return b
}
//同花
func SameColor(doubleArray [][]int, a int) (b string) {
	var j int = 0
	for i := 0; i < a-1; i++ {
		if doubleArray[i][1] == doubleArray[i+1][1] {
			j++
			if j >= 5 {
				b = "同花"
			}
			}
		}
	return b
}
//连对、三条、四条
func Consecutive(doubleArray [][]int, a int) (b string) {
	for i := 0; i < a-4; i++ {
		temp := doubleArray[i][0]
		if temp == doubleArray[i+1][0] {
			b="连对"
			if temp==doubleArray[i+2][0]{
				b="三条"
				if temp==doubleArray[i+3][0]{
					b="四条"
					break
				}
				if doubleArray[i+3][0]==doubleArray[i+4][0]{
					b="满堂彩"
					break
				}
			}
			if doubleArray[i+2][0]==doubleArray[i+4][0]{
				b="满堂彩"
				break
			}
		}
	}
	return b
}
//散子
func AlonePoker(doubleArr [][]int,a int)  (b string){
	Same:=SameColor(doubleArr,a)
	var  j int
for i:=0;i<a-1;i++{
	if doubleArr[i][0]!=doubleArr[i+1][0]&&doubleArr[i][1]!=doubleArr[i+1][1]&&Same!="顺子"{
		j++
	}
	if j>=5{
		b="散子"
	}
}
return b
}