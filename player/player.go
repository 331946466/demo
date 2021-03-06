package player

import (
	"bufio"
	"fmt"
	"mygo/spareArray"
	"os"
)
func Player(Player string)(GradeResult int,PokerResult string,ArrayResult [][] int,PlayerResult string){
	fmt.Println(Player+"请输入测试牌型")
	inputReader := bufio.NewReader(os.Stdin)
	Poker, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("输入错误")
	}
	BytePoker := []byte(Poker)
	//将牌存储到稀疏数组中
	Grade,PokerName,Array:=spareArray.ArrayPoker(BytePoker)
	return Grade,PokerName,Array,Player
}