package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

const (
	iterationsNum = 6
	goroutinesNum = 6
)

func doWork(wg *sync.WaitGroup, th int) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(th, j))
		runtime.Gosched()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	runtime.GOMAXPROCS(1)
	wg.Add(goroutinesNum)
	for i := 0; i < goroutinesNum; i++ {
		go doWork(wg, i)
	}
	wg.Wait()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), " ",
		strings.Repeat(" ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat(" ", j))
}

/*
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup по завершению функции
	for i := 0; i < 3; i++ {
		fmt.Println(message, "Iteration:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup // Создаем WaitGroup
	runtime.GOMAXPROCS(6)
	wg.Add(3) // Увеличиваем счетчик на 3, так как будем запускать 3 горутины

	go printMessage("Goroutine 1", &wg)
	go printMessage("Goroutine 2", &wg)
	go printMessage("Goroutine 3", &wg)

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("All goroutines finished")
}
*/

/*
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	var wg sync.WaitGroup

	// Запуск нескольких горутин для обработки HTTP-запросов
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			http.HandleFunc(fmt.Sprintf("/hello%d", id), handler)
		}(i)
	}

	go func() {
		// Запускаем HTTP-сервер
		http.ListenAndServe(":8080", nil)
	}()

	wg.Wait() // Ждем завершения всех горутин
}
*/

/*
Теоретические основы горутин

	1.	Горутины:
	•	Это легковесные потоки, управляемые планировщиком Go.
	•	Они гораздо легче системных потоков, что позволяет запускать тысячи горутин одновременно без значительных затрат памяти и процессорного времени.
	2.	Планировщик Go:
	•	Планировщик Go реализует модель совместного использования ресурсов (M:N), где M — это количество горутин, а N — количество системных потоков.
	•	Планировщик автоматически распределяет горутины между потоками операционной системы для эффективного использования ресурсов.
	3.	Синхронизация:
	•	При использовании горутин важно обеспечивать правильную синхронизацию при доступе к общим данным.
	•	Использование примитивов из пакета sync (например, WaitGroup, Mutex, Once) позволяет предотвратить гонки данных и обеспечить корректное выполнение программ.

Горутины в Go позволяют легко реализовывать параллелизм и асинхронные задачи, что делает их незаменимым инструментом для высокопроизводительных программ.
*/

/*
Основные моменты использования sync.WaitGroup

	1.	Добавление задач: Метод Add(n int) увеличивает счетчик ожидающих горутин на n. Это означает, что вы ожидаете, что будет запущено n горутин.
	2.	Сигнал о завершении: Каждая горутина, когда завершает свою работу, должна вызвать метод Done(). Это уменьшает счетчик ожидающих горутин на 1.
	3.	Ожидание завершения: Метод Wait() блокирует выполнение текущей горутины до тех пор, пока счетчик не станет равным нулю. Это значит, что все добавленные горутины завершили свою работу.

Пример использования

Вот простой пример кода, который демонстрирует использование sync.WaitGroup с wg.Add, wg.Done и wg.Wait:

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшает счетчик на 1, когда работа завершена
	fmt.Printf("Worker %d: starting\n", id)
	time.Sleep(2 * time.Second) // Имитация работы
	fmt.Printf("Worker %d: done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Добавляем одну горутину
		go worker(i, &wg) // Запускаем горутину
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("All workers completed.")
}

бъяснение кода

	1.	Объявление WaitGroup: Мы объявляем переменную wg типа sync.WaitGroup.
	2.	Запуск горутин:
	•	В цикле for мы добавляем 1 к счетчику с помощью wg.Add(1) перед запуском каждой горутины.
	•	Каждая горутина выполняет функцию worker, которая имитирует работу с помощью time.Sleep.
	3.	Сигнал о завершении: Внутри функции worker мы используем defer wg.Done(), чтобы гарантировать, что счетчик будет уменьшен при выходе из функции.
	4.	Ожидание завершения: Метод wg.Wait() блокирует выполнение главной горутины до тех пор, пока все запущенные горутины не завершатся.

Заключение

Использование sync.WaitGroup и методов Add, Done и Wait позволяет легко управлять параллельными задачами и гарантировать, что все они завершены перед выполнением следующего кода. Это особенно полезно в многопоточных приложениях, где необходимо контролировать порядок выполнения и синхронизацию горутин.

*/
