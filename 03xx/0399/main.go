package main

import "fmt"

/* 0ms */
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	edge := make(map[string]map[string]float64)
	unionMap := make(map[string]string)
	for i := 0; i < len(values); i++ {
		a, b := equations[i][0], equations[i][1]
		if edge[a] == nil {
			edge[a] = make(map[string]float64)
			unionMap[a] = a
		}
		if edge[b] == nil {
			edge[b] = make(map[string]float64)
			unionMap[b] = b
		}
		x, y := float64(1), float64(1)
		for a != unionMap[a] {
			x *= edge[a][unionMap[a]]
			a = unionMap[a]
		}
		for b != unionMap[b] {
			y *= edge[b][unionMap[b]]
			b = unionMap[b]
		}
		if a != b {
			unionMap[a] = b
			edge[a][b] = y * values[i] / x
		}
	}
	res := make([]float64, 0)
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		if edge[a] == nil || edge[b] == nil {
			res = append(res, -1)
			continue
		}
		x, y := float64(1), float64(1)
		for a != unionMap[a] {
			x *= edge[a][unionMap[a]]
			a = unionMap[a]
		}
		for b != unionMap[b] {
			y *= edge[b][unionMap[b]]
			b = unionMap[b]
		}
		if a != b {
			res = append(res, -1)
			continue
		}
		res = append(res, x/y)
	}
	return res
}

func main() {
	var equations [][]string
	var values []float64
	var queries [][]string

	equations = [][]string{{"a", "b"}, {"b", "c"}}
	values = []float64{2, 3}
	queries = [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(equations)
	fmt.Println(values)
	fmt.Println(queries)
	fmt.Println(calcEquation(equations, values, queries))
}
