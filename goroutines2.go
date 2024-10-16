package main

import (
	"fmt"
	"time"
)

// чтобы убедиться, что все горутины будут выполнены
// до завершения основной функции, мы усыпляем процесс,
// чтобы дать возможность выполниться другим процессам.

func display(message string) {
	//for {
	//	fmt.Println(message)
	//	time.Sleep(1 * time.Second)
	//}
	fmt.Println(message)

}

// При запуске нескольких горутин вместе они взаимодействуют
//друг с другом одновременно. Порядок выполнения горутины
//случаен, поэтому мы можем не получить ожидаемый результат.

func main() {
	//go display("Process 1")
	//display("Process 2")

	go display("Print 1")
	go display("Print 2")
	go display("Print 3")

	time.Sleep(1 * time.Second)
}

// Преимущества горутин:
// * С помощью Goroutines достигается параллелизм в программировании Go. Он помогает двум или более независимым функциям работать вместе.
// * Горутины можно использовать для запуска фоновых операций в программе.
// * Они общаются по частным каналам, поэтому общение между ними безопаснее.
// * С помощью горутин мы можем разделить одну задачу на несколько сегментов для лучшего выполнения.
