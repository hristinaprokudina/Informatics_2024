package lab9

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskManager struct {
	Tasks []Task
}

type TaskManagerInterface interface {
	AddTask(description string)
	ShowTasks()
	CompleteTask(index int)
	DeleteTask(index int)
	SearchTask(keyword string)
	SaveToFile(filename string)
	LoadFromFile(filename string)
}

func (tm *TaskManager) AddTask(description string) {
	tm.Tasks = append(tm.Tasks, Task{Description: description, Completed: false})
}

func (tm *TaskManager) ShowTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("Нет задач для отображения.")
		return
	}
	for i, task := range tm.Tasks {
		status := "Не выполнена"
		if task.Completed {
			status = "Выполнена"
		}
		fmt.Printf("%d: %s [%s]\n", i+1, task.Description, status)
	}
}

func (tm *TaskManager) CompleteTask(index int) {
	if index < 0 || index >= len(tm.Tasks) {
		fmt.Println("Ошибка: индекс задачи вне диапазона.")
		return
	}
	tm.Tasks[index].Completed = true
}

func (tm *TaskManager) DeleteTask(index int) {
	if index < 0 || index >= len(tm.Tasks) {
		fmt.Println("Ошибка: индекс задачи вне диапазона.")
		return
	}
	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
}

func (tm *TaskManager) SearchTask(keyword string) {
	found := false
	for _, task := range tm.Tasks {
		if strings.Contains(task.Description, keyword) {
			status := "Не выполнена"
			if task.Completed {
				status = "Выполнена"
			}
			fmt.Printf("Найдена задача: %s [%s]\n", task.Description, status)
			found = true
		}
	}
	if !found {
		fmt.Println("Задачи с таким ключевым словом не найдены.")
	}
}

func (tm *TaskManager) SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Ошибка при сохранении файла: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tm.Tasks); err != nil {
		fmt.Printf("Ошибка при записи в файл: %v\n", err)
	}
}

func (tm *TaskManager) LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tm.Tasks); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}
}

func Runlab9() {
	taskManager := TaskManager{}
	taskManager.LoadFromFile("tasks.json")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Выйти")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Введите описание задачи:")
			scanner.Scan()
			description := scanner.Text()
			taskManager.AddTask(description)
		case "2":
			taskManager.ShowTasks()
		case "3":
			fmt.Println("Введите номер задачи для отметки как выполненной:")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err == nil {
				taskManager.CompleteTask(index - 1) // Индексация с 0
			} else {
				fmt.Println("Ошибка: некорректный ввод.")
			}
		case "4":
			fmt.Println("Введите номер задачи для удаления:")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err == nil {
				taskManager.DeleteTask(index - 1) // Индексация с 0
			} else {
				fmt.Println("Ошибка: некорректный ввод.")
			}
		case "5":
			fmt.Println("Введите ключевое слово для поиска:")
			scanner.Scan()
			keyword := scanner.Text()
			taskManager.SearchTask(keyword)
		case "6":
			taskManager.SaveToFile("tasks.json")
			fmt.Println("Задачи сохранены. Выход из программы.")
			return
		default:
			fmt.Println("Ошибка: некорректный выбор.")
		}
	}
}