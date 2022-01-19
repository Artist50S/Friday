package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(6)
	days := 7
	friday := 5
	thirteen := 0
	january := n + 13
	february := january + 31
	marth := february + 28
	april := marth + 31
	may := april + 30
	jule := may + 31
	june := jule + 30
	august := june + 31
	september := august + 31
	october := september + 30
	november := october + 31
	december := november + 30
	switch n {
	case 0:
		fmt.Println("monday 1 january")
	case 1:
		fmt.Println("tuesday 1 january")
	case 2:
		fmt.Println("wednesday 1 january")
	case 3:
		fmt.Println("thursday 1 january")
	case 4:
		fmt.Println("friday 1 january")
	case 5:
		fmt.Println("saturday 1 january")
	default:
		fmt.Println("sunday 1 january")
	}

	if january%days == friday { // нужно делить на 7 и проверять остаток. Он должен быть равен 5.
		thirteen++
	}
	if february%days == friday {
		thirteen++
	}
	if marth%days == friday {
		thirteen++
	}
	if april%days == friday {
		thirteen++
	}
	if may%days == friday {
		thirteen++
	}
	if june%days == friday {
		thirteen++
	}
	if jule%days == friday {
		thirteen++
	}
	if august%days == friday {
		thirteen++
	}
	if september%days == friday {
		thirteen++
	}
	if october%days == friday {
		thirteen++
	}
	if november%days == friday {
		thirteen++
	}
	if december%days == friday {
		thirteen++
	}
	fmt.Println(thirteen, "friday 13")
}
