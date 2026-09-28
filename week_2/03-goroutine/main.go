package main

import (
	"fmt"
	"sync"
)

// func sayHello() {
// 	fmt.Println("Hello ")
// }

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // funksiya tugaganda counter -1
// 	fmt.Printf("Worker %d ishni boshladi\n", id)
// 	// ish simulyatsiyasi...
// 	fmt.Printf("Worker %d ishni tugatdi\n", id)
// }
//

func worker(id int, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: %s\n", id, message)
}

func main() {
	// go sayHello()           // yangi goroutine ishga tushdi
	// time.Sleep(time.Second) // main tugab ketmasligi uchun kutamiz

	// sync.WaitGroup — bu counter (hisoblagich) bo'lib, "nechta goroutine tugashini kutyapman" degan ma'lumotni saqlaydi. Uchta metodi bor:
	// sync.WaitGroup — bu counter (hisoblagich) bo'lib, "nechta goroutine tugashini kutyapman" degan ma'lumotni saqlaydi. Uchta metodi bor:

	// Add(n) — counter'ga n qo'shadi (nechta goroutine kutilayotganini bildiradi)
	// Done() — counter'ni 1ga kamaytiradi (bitta goroutine tugadi degani; Add(-1) bilan bir xil)
	// Wait() — counter 0 bo'lguncha bloklaydi (kutadi)
	//
	// var wg sync.WaitGroup

	// for i := 1; i <= 5; i++ {
	// 	wg.Add(1)         // har bir goroutine uchun +1
	// 	go worker(i, &wg) // pointer orqali uzatiladi!
	// }

	// wg.Wait() // barcha 5 ta worker tugashini kutadi
	// fmt.Println("Barcha workerlar tugadi")
	//

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, fmt.Sprintf("Hello from worker %d", i), &wg)
	}
	wg.Wait()
	fmt.Println("Barcha workerlar tugadi")

}
