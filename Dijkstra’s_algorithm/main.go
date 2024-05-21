package main

import (
	"fmt"
	"math"
)

type Graph map[string]map[string]int

func Dijkstra(graph Graph, start string) map[string]int {
	distances := make(map[string]int) // Создаем словарь для хранения расстояний до вершин
	visited := make(map[string]bool)  // Создаем словарь для отслеживания посещенных вершин

	// Инициализируем расстояния до всех вершин как "бесконечность"
	for vertex := range graph {
		distances[vertex] = math.MaxInt32
	}
	// Расстояние до начальной вершины устанавливаем как 0
	distances[start] = 0

	for len(visited) < len(graph) {
		// Выбираем вершину с наименьшим расстоянием из непосещенных
		minDist := math.MaxInt32
		minVertex := ""
		for vertex, dist := range distances {
			if !visited[vertex] && dist < minDist {
				minDist = dist
				minVertex = vertex
			}
		}

		// Помечаем выбранную вершину как посещенную
		visited[minVertex] = true

		// Обновляем расстояния до всех соседних вершин выбранной вершины
		for neighbor, weight := range graph[minVertex] {
			if !visited[neighbor] {
				newDist := distances[minVertex] + weight
				if newDist < distances[neighbor] {
					distances[neighbor] = newDist
				}
			}
		}
	}

	return distances
}

func main() {
	graph := Graph{
		"A": {"B": 2, "C": 4},
		"B": {"C": 1, "D": 7},
		"C": {"D": 3},
		"D": {},
	}

	start := "A"
	distances := Dijkstra(graph, start)
	fmt.Println("Shortest distances from", start, ": ", distances)
}
