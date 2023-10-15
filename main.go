package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//fmt.Println("Первое задание")
	//firstTask(10)

	fmt.Println("Второе задание")
	secondTask(10, 2)

	//fmt.Println("Третье задание")
	//thirdTask()

	//fmt.Println("Четвертое задание")
	//fourthTask()

	//fmt.Println("Пятое задание")
	//fifthTask(10)

	//fmt.Println("Шествое задание")
	//circularDistribution(10, 2)
}

func firstTask(n int) {
	vector := make([]int, n)

	for i := 0; i < n; i++ {
		vector[i] = i + 1
	}

	fmt.Println("искоходный вектор: ", vector)

	t := time.Now()
	fmt.Println("Время начала", t)
	multiplier := 2

	for i := 0; i < len(vector); i++ {
		time.Sleep(time.Second)
		vector[i] *= multiplier
	}

	fmt.Println("обработанный вектор: ", vector)
	fmt.Println("Прошло времени(первое): ", time.Since(t))
}

func secondTask(n, m int) {

	vectorChan := make(chan int, n)
	results := make(chan int, n)
	multiplier := 2

	for i := 0; i < m; i++ {
		go reValue(i+1, vectorChan, results, multiplier)
	}

	fmt.Println("исходный вектор")
	for i := 0; i < n; i++ {
		value := i + 1
		fmt.Print(value, " ")
		vectorChan <- value
	}
	close(vectorChan)

	t := time.Now()
	fmt.Println("\nВремя начала", t)

	for i := 0; i < n; i++ {
		fmt.Printf("Результат %d -> значение = %d\n", i+1, <-results)
	}

	fmt.Println("Прошло времени: ", time.Since(t))
}

func reValue(id int, vectorChan <-chan int, results chan<- int, multiplier int) {
	for value := range vectorChan {
		time.Sleep(time.Second)
		fmt.Printf("Поток %d закончил\n", id)
		results <- value * multiplier
	}
}

func thirdTask() {
	nValues := []int{10, 100, 1000, 100000}
	mValues := []int{2, 3, 4, 5, 10}
	multiplier := 2

	for _, n := range nValues {

		for _, m := range mValues {
			t := time.Now()

			vectorChan := make(chan int, n)
			results := make(chan int, n)

			var wg sync.WaitGroup

			for i := 0; i < n; i++ {
				vectorChan <- i + 1
			}

			close(vectorChan)
			for i := 0; i < m; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					reValueForThird(vectorChan, results, multiplier)
				}()
			}

			wg.Wait()

			fmt.Printf("%d\t%d\t%s\n", n, m, time.Since(t))
		}
	}
}

func reValueForThird(vectorChan <-chan int, results chan<- int, multiplier int) {
	for value := range vectorChan {
		time.Sleep(time.Microsecond)
		results <- value * multiplier
	}
}

func fourthTask() {
	nValues := []int{10, 100, 1000, 100000}
	mValues := []int{2, 3, 4, 5, 10}
	multiplier := 2

	for _, n := range nValues {

		for _, m := range mValues {
			t := time.Now()

			vectorChan := make(chan int, n)
			results := make(chan float64, n)

			var wg sync.WaitGroup

			for i := 0; i < n; i++ {
				vectorChan <- i + 1
			}

			close(vectorChan)
			for i := 0; i < m; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					reValueForFourth(vectorChan, results, multiplier)
				}()
			}

			wg.Wait()

			fmt.Printf("%d\t%d\t%s\n", n, m, time.Since(t))
		}
	}
}

func reValueForFourth(vectorChan <-chan int, results chan<- float64, multiplier int) {
	for value := range vectorChan {
		time.Sleep(time.Microsecond)
		results <- math.Sin(math.Pow(float64(value*multiplier), 6))
	}
}

func fifthTask(n int) {
	vectorChan1 := make(chan int, 1)
	vectorChan2 := make(chan int, 2)
	vectorChan3 := make(chan int, 3)
	vectorChan4 := make(chan int, 4)
	results := make(chan int, n)
	multiplier := 2

	for i := 0; i < 4; i++ {
		if i == 0 {
			go reValue(i+1, vectorChan1, results, multiplier)
		} else if i == 1 {
			go reValue(i+1, vectorChan2, results, multiplier)
		} else if i == 2 {
			go reValue(i+1, vectorChan3, results, multiplier)
		} else {
			go reValue(i+1, vectorChan4, results, multiplier)
		}

	}

	for j := 0; j < 4; j++ {
		if j == 0 {
			vectorChan1 <- j + 1
		} else if j == 1 {
			vectorChan2 <- j + 1
			vectorChan2 <- j + 2
		} else if j == 2 {
			vectorChan3 <- j + 2
			vectorChan3 <- j + 3
			vectorChan3 <- j + 4
		} else {
			vectorChan4 <- j + 3
			vectorChan4 <- j + 4
			vectorChan4 <- j + 5
			vectorChan4 <- j + 6
		}
	}

	close(vectorChan1)
	close(vectorChan2)
	close(vectorChan3)
	close(vectorChan4)

	t := time.Now()
	fmt.Println("\nВремя начала", t)

	for i := 0; i < n; i++ {
		fmt.Printf("Результат %d -> значение = %d\n", i+1, <-results)
	}

	fmt.Println("Прошло времени: ", time.Since(t))
}

func circularDistribution(n, m int) {
	vector := make([]float64, n)

	for i := 0; i < n; i++ {
		vector[i] = rand.Float64()
	}

	resultVector := make([]float64, n)
	multiplier := 2.0

	var wg sync.WaitGroup

	t := time.Now()
	fmt.Println("\nВремя начала", t)

	for threadID := 0; threadID < m; threadID++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for i := threadID; i < n; i += m {
				resultVector[i] = vector[i] * multiplier
			}
		}(threadID)
	}

	wg.Wait()
	fmt.Println("Прошло времени: ", time.Since(t))
}
