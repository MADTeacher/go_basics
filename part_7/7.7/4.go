package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Количество философов и вилок
const numPhilosophers = 5

// Философ представляет участника задачи
type Philosopher struct {
	id           int
	leftFork     *sync.Mutex
	rightFork    *sync.Mutex
	mealsEaten   int
	thinkingTime time.Duration
	eatingTime   time.Duration
	wg           *sync.WaitGroup
}

// Функция приема пищи философом
// Пытается взять две вилки и затем есть
func (p *Philosopher) eat() {
	defer p.wg.Done()

	// Количество приемов пищи для каждого философа
	for i := 0; i < 3; i++ {
		// Думаем некоторое время перед тем как пытаться взять вилки
		fmt.Printf("Philosopher %d is thinking... 🤔\n", p.id)
		time.Sleep(p.thinkingTime)

		// Решение проблемы взаимной блокировки:
		// Нечетные философы сначала берут левую вилку, четные - правую
		if p.id%2 == 0 {
			// Четные философы: сначала правая, потом левая вилка
			p.rightFork.Lock()
			fmt.Printf("Philosopher %d picked up right fork 🍴\n", p.id)
			// Небольшая задержка для наглядности
			time.Sleep(10 * time.Millisecond)

			p.leftFork.Lock()
			fmt.Printf(
				"Philosopher %d picked up left fork and started eating\n",
				p.id,
			)
		} else {
			// Нечетные философы: сначала левая, потом правая вилка
			p.leftFork.Lock()
			fmt.Printf("Philosopher %d picked up left fork 🍴😊\n", p.id)
			time.Sleep(10 * time.Millisecond)

			p.rightFork.Lock()
			fmt.Printf(
				"Philosopher %d picked up right fork and started eating 🍴😊\n",
				p.id,
			)
		}

		// Прием пищи
		fmt.Printf("Philosopher %d is eating... 🍽️😋\n", p.id)
		time.Sleep(p.eatingTime)
		p.mealsEaten++

		// Кладем вилки обратно на стол
		p.leftFork.Unlock()
		p.rightFork.Unlock()
		fmt.Printf(
			"Philosopher %d put down forks and finished eating (%d times) ✓\n",
			p.id, p.mealsEaten,
		)
	}

	fmt.Printf("Philosopher %d finished dining and left the table 👋\n", p.id)
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Создаем группу ожидания для синхронизации горутин
	var wg sync.WaitGroup

	// Создаем вилки (мьютексы)
	forks := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &sync.Mutex{}
	}

	// Создаем философов
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		// Каждый философ имеет доступ к вилке слева и справа
		leftForkIndex := i
		rightForkIndex := (i + 1) % numPhilosophers
		thinkingTime := time.Duration(rnd.Intn(500)+100) * time.Millisecond
		eatingTime := time.Duration(rnd.Intn(300)+200) * time.Millisecond
		// Создаем философа с случайным временем размышления и приема пищи
		philosophers[i] = &Philosopher{
			id:           i,
			leftFork:     forks[leftForkIndex],
			rightFork:    forks[rightForkIndex],
			mealsEaten:   0,
			thinkingTime: thinkingTime,
			eatingTime:   eatingTime,
			wg:           &wg,
		}

		// Добавляем философа в группу ожидания
		wg.Add(1)
	}

	// Запускаем философов в отдельных горутинах
	fmt.Println("Philosophers are sitting at the table... 🪑🪑🪑")
	for i := 0; i < numPhilosophers; i++ {
		go philosophers[i].eat()
	}

	// Ожидаем завершения трапезы всех философов
	wg.Wait()
	fmt.Println("All philosophers have finished dining! 🎉")

	// Выводим статистику по приемам пищи
	for i, p := range philosophers {
		fmt.Printf("Philosopher %d ate %d times 🍽️\n", i, p.mealsEaten)
	}
}
