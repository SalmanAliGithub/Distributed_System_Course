package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	id      int
	data    string
	retries int
}

const maxRetries = 2

// Worker function that processes tasks and handles failures with retry logic.
func worker(id int, taskChan <-chan Task, wg *sync.WaitGroup, failChan chan<- Task) {
	defer wg.Done()
	for task := range taskChan {
		fmt.Printf("Worker %d started task %d: %s\n", id, task.id, task.data)

		// Simulate random failure (30% chance of failure)
		if rand.Float32() < 0.3 {
			fmt.Printf("Worker %d failed on task %d\n", id, task.id)
			task.retries++
			if task.retries > maxRetries {
				fmt.Printf("Task %d exceeded max retries and will not be retried.\n", task.id)
			} else {
				failChan <- task // Reassign the failed task
			}
			return
		}

		// Simulate task processing time
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
		fmt.Printf("Worker %d completed task %d\n", id, task.id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Define tasks
	tasks := []Task{
		{id: 1, data: "Task 1"},
		{id: 2, data: "Task 2"},
		{id: 3, data: "Task 3"},
		{id: 4, data: "Task 4"},
		{id: 5, data: "Task 5"},
	}

	// Channels for task distribution and failure handling
	taskChan := make(chan Task, len(tasks))
	failChan := make(chan Task, len(tasks))

	// WaitGroup to ensure all workers finish their tasks
	var wg sync.WaitGroup
	numWorkers := 3

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskChan, &wg, failChan)
	}

	// Distribute tasks in round-robin fashion
	for i, task := range tasks {
		fmt.Printf("Task %d assigned to worker %d\n", task.id, (i%numWorkers)+1)
		taskChan <- task
	}
	close(taskChan)

	// Handle failed tasks with retries
	go func() {
		for failedTask := range failChan {
			fmt.Printf("Reassigning failed task %d (Retry %d)\n", failedTask.id, failedTask.retries)
			wg.Add(1)
			go worker(rand.Intn(numWorkers)+1, taskChan, &wg, failChan)
		}
	}()

	// Wait for all workers to complete
	wg.Wait()
	close(failChan)
	fmt.Println("All tasks completed.")
}
