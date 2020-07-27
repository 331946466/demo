package model

import "fmt"

//判断牌型等级
func GradePoker(ShunZi string,Same string,Consecutive string,AlonePoker string,MaxPoker int)(GradeResult int,NameResult string){
	var Grade int
	var Name string
//皇家同花
	if ShunZi=="顺子"&&Same=="同花"&&MaxPoker==14{
		Grade=10
		Name="皇家同花"
	}else if ShunZi=="顺子"&&Same=="同花"{
		Grade=9
		Name="同花顺"
	}else if Consecutive=="四条"{
		Grade=8
		Name="四条"
	}else if Consecutive=="满堂彩"{
		Grade=7
		Name="满堂彩"
	}else  if Same=="同花"&&ShunZi!="顺子"{
		Grade=6
		Name="同花"
	}else  if ShunZi=="顺子"&&Same!="同花"{
		Grade=5
		Name="顺子"
	}else  if Consecutive=="三条"{
		Grade=4
		Name="三条"
	}else  if Consecutive=="两对子"{
		Grade=3
		Name="两对"
	}else if Consecutive=="对子"{
		Grade=2
		Name="对子"
	}else if AlonePoker=="散子"{
		Grade=1
		Name="散子"
	}
	fmt.Println(Name+"牌型等级为:",Grade)
	return Grade,Name
}