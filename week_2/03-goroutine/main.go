package main

import (
	"fmt"
	"sync"
)

type Student struct {
	id   int
	name string
	age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Level: %d", s.id, s.name, s.age)
}

func ProcessStudent(s *Student) error {
	if s.age > 25 {
		return fmt.Errorf("Student is too old to process")
	}
	fmt.Println("Processing student:", s)
	return nil
}

func processStudents(students []*Student) {
	var wg sync.WaitGroup

	for _, s := range students {
		wg.Add(1)
		go func(s *Student) {
			defer wg.Done()
			if err := ProcessStudent(s); err != nil {
				fmt.Println(err)
			}
		}(s)
	}

	wg.Wait()
}

func main() {

	students := []*Student{
		{id: 1, name: "Alice", age: 20},
		{id: 2, name: "Bob", age: 30},
		{id: 3, name: "Charlie", age: 25},
		{id: 4, name: "David", age: 28},
		{id: 5, name: "Eve", age: 22},
		{id: 6, name: "Frank", age: 27},
	}
	processStudents(students)
}

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

// func worker(id int, message string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d: %s\n", id, message)
// }

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

// var wg sync.WaitGroup
// for i := 1; i <= 5; i++ {
// 	wg.Add(1)
// 	go worker(i, fmt.Sprintf("Hello from worker %d", i), &wg)
// }
// wg.Wait()
// fmt.Println("Barcha workerlar tugadi")
