package main

import (
	"fmt"
	"math"
)

type Graph map[string]map[string]int

func BellmanFord(graph Graph, start string) (map[string]int, bool) {
	distances := make(map[string]int) // Создаем словарь для хранения расстояний до вершин

	// Инициализируем расстояния до всех вершин как "бесконечность"
	for vertex := range graph {
		distances[vertex] = math.MaxInt32
	}
	// Расстояние до начальной вершины устанавливаем как 0
	distances[start] = 0

	//Обновление растояния
	for range graph {
		for src, edges := range graph {
			for dst, weight := range edges {
				if distances[src] != math.MaxInt32 && distances[src]+weight < distances[dst] {
					distances[dst] = distances[src] + weight
				}
			}
		}
	}

	// Проверка наличия отрицательных циклов
	for src, edges := range graph {
		for dst, weight := range edges {
			if distances[src] != math.MaxInt32 && distances[src]+weight < distances[dst] {
				return distances, true // Обнаружен отрицательный цикл
			}
		}
	}

	return distances, false // Отрицательных циклов нет
}

func main() {
	graph := Graph{
		"A": {"B": 2, "C": 4},
		"B": {"C": 1, "D": 7},
		"C": {"D": -3},
		"D": {"A": 10}, // В этом примере добавим отрицательный цикл
	}

	start := "A"
	distances, hasNegativeCycle := BellmanFord(graph, start)
	if hasNegativeCycle {
		fmt.Println("Граф содержит отрицательный цикл")
	} else {
		fmt.Println("Shortest distances from", start, ": ", distances)
	}
}
