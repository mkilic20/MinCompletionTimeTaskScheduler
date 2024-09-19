package main

import (
	"fmt"
)

// Task represents a job task with a name, duration, and its dependencies.
type Task struct {
	name      string
	duration  int
	dependsOn []string
}

// Function to find the minimum time to complete all tasks.
func minimumCompletionTime(tasks []Task) (int, []string, error) {
	// Create task map to store task details for quick access.
	taskMap := make(map[string]Task)
	// Create a map to store in-degree of each task.
	inDegree := make(map[string]int)
	// Create an adjacency list to represent dependencies.
	graph := make(map[string][]string)
	// Create a map to store the earliest completion time for each task.
	earliestCompletionTime := make(map[string]int)

	// Initialize maps with task data.
	for _, task := range tasks {
		taskMap[task.name] = task
		inDegree[task.name] = 0
		earliestCompletionTime[task.name] = task.duration
	}

	// Build the graph and in-degree count.
	for _, task := range tasks {
		for _, dep := range task.dependsOn {
			if _, exists := taskMap[dep]; !exists {
				return 0, nil, fmt.Errorf("task dependency %s does not exist", dep)
			}
			graph[dep] = append(graph[dep], task.name)
			inDegree[task.name]++
		}
	}

	// Queue for processing tasks with no dependencies.
	queue := []string{}
	for name, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, name)
		}
	}

	// Order of task execution.
	order := []string{}
	// Process the queue.
	for len(queue) > 0 {
		// Get the first task in the queue.
		current := queue[0]
		queue = queue[1:]

		// Append current task to the order.
		order = append(order, current)

		// Process each dependent task.
		for _, neighbor := range graph[current] {
			// Update the earliest completion time for the dependent task.
			if earliestCompletionTime[current]+taskMap[neighbor].duration > earliestCompletionTime[neighbor] {
				earliestCompletionTime[neighbor] = earliestCompletionTime[current] + taskMap[neighbor].duration
			}

			// Decrease the in-degree of the dependent task.
			inDegree[neighbor]--

			// If in-degree becomes zero, add it to the queue.
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check for tasks that could not be processed due to circular dependencies.
	if len(order) != len(tasks) {
		return 0, nil, fmt.Errorf("circular dependency detected, cannot complete all tasks")
	}

	// Find the maximum completion time from the earliestCompletionTime map.
	maxCompletionTime := 0
	for _, time := range earliestCompletionTime {
		if time > maxCompletionTime {
			maxCompletionTime = time
		}
	}

	return maxCompletionTime, order, nil
}

func main() {
	// Define tasks as per the problem statement.
	tasks := []Task{
		{name: "A", duration: 3, dependsOn: []string{}},
		{name: "B", duration: 2, dependsOn: []string{}},
		{name: "C", duration: 4, dependsOn: []string{}},
		{name: "D", duration: 5, dependsOn: []string{"A"}},
		{name: "E", duration: 2, dependsOn: []string{"B", "C"}},
		{name: "F", duration: 3, dependsOn: []string{"D", "E"}},
	}

	// Define tasks with circular dependencies.
	// tasks := []Task{
	// 	{name: "A", duration: 3, dependsOn: []string{"B"}},
	// 	{name: "B", duration: 2, dependsOn: []string{"C"}},
	// 	{name: "C", duration: 4, dependsOn: []string{"A"}},
	// }

	// Calculate minimum completion time and task order.
	minTime, order, err := minimumCompletionTime(tasks)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Output results.
	fmt.Printf("Minimum Completion Time: %d units\n", minTime)
	fmt.Printf("Task Execution Order: %v\n", order)
}
