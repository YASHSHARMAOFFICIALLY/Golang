// package main

// import "fmt"

// func main() {
// 	fmt.Println("main start")
// 	fmt.Println("main end")
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go func() {
// 		fmt.Println("background task running")
// 	}()

// 	fmt.Println("main does its work")
// 	time.Sleep(100 * time.Millisecond)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func notifyShipped(orderID int) {
// 	fmt.Println("Sending shipping notification for order", orderID)
// }

// func main() {
// 	go notifyShipped(101)
// 	go notifyShipped(102)
// 	go notifyShipped(103)

// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("main done")
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func notifyCustomer(orderID int) {
// 	time.Sleep(50 * time.Millisecond)
// 	fmt.Println("notified order", orderID)
// }

// func main() {
// 	start := time.Now()

// 	for i := 1; i <= 5; i++ {
// 		go notifyCustomer(i)
// 	}
// 	time.Sleep(100 * time.Millisecond)

//		fmt.Println("total:", time.Since(start))
//	}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	orders := []int{101, 102, 103}

// 	for i := 0; i < len(orders); i++ {
// 		go func() {
// 			fmt.Println("processing order", orders[i])
// 		}()
// 	}

// 	time.Sleep(100 * time.Millisecond)
// }

package main

import "fmt"

func main() {
	go func() {
		fmt.Println("this might never print")
	}()
}
