package main

import (
	"fmt"
	"sort"
	"strings"
)

func clone(list []int) []int {
	cloned := make([]int, len(list))
	copy(cloned, list)
	return cloned
}

func deepClone(list[][]int) [][]int {
	cloned := make([][]int, len(list))
	for i, el := range list {
		cloned[i] = clone(el)
	}
	return cloned
}

func findSolution(upToNow [][]int, goals []int, n int, limit int) [][][]int {
	if limit <= 0 {
		return nil
	}
	if n < 0 {
		return nil
	}
	solved := true
	for _, goal := range goals {
		if goal != 0 {
			solved = false
			break
		}
	}
	if solved {
		return [][][]int{upToNow}
	}

	solutions := [][][]int{}
	for i, _ := range goals {
		if n > goals[i] {
			continue
		}
		newUpToNow := deepClone(upToNow)
		newGoals := clone(goals)

		newUpToNow[i] = append(newUpToNow[i], n)
		newGoals[i] -= n

		newSolutions := findSolution(newUpToNow, newGoals, n-1, limit)
		solutions = append(solutions, newSolutions...)
		limit -= len(newSolutions)

		if len(upToNow[i]) == 0 {
			break
		}
	}
	return solutions
}

func printSol(sols [][]int, n int) {
	for _, sol := range sols {
		line := []rune(strings.Repeat(" ", 2*n))
		for _, num := range sol {
			line[2*(num-1)] = '█'
			line[2*(num-1)+1] = '█'
		}
		fmt.Print(string(line))
		fmt.Print(" = ")
		for _, num := range sol {
			fmt.Print(num)
			fmt.Print(", ")
		}
		fmt.Println()
	}
}

func printSols(sols [][][]int, n int) {
	for _, sol := range sols {
		fmt.Println(strings.Repeat("-", 2*n))
		printSol(sol, n)
		fmt.Println(strings.Repeat("-", 2*n))
		//printSol(sortedByBeggining(sol), n)
		//fmt.Println(strings.Repeat("-", 2*n))
		fmt.Println()
	}
}

func splitSum(n, k, limit int) [][][]int {
	if n*(n+1) % (2*k) != 0 {
		return nil
	}

	target :=  n*(n+1)/2/k

	goals := make([]int, k)
	for i := 0; i < k; i++ {
		goals[i] = target
	}
	upToNow := make([][]int, k)
	solutions := findSolution(upToNow, goals, n, limit)
	if len(solutions) == 0 {
		fmt.Println("EXCEPTION", n, k)
	}
	return solutions
}

func unevenSplitSum(n, k, target, limit int) [][][]int {
	sum := n*(n+1)/2
	lastGroup := sum - target*(k-1)

	goals := make([]int, k)
	for i := 0; i < k-1; i++ {
		goals[i] = target
	}

	goals[k-1] = lastGroup
	upToNow := make([][]int, k)
	solutions := findSolution(upToNow, goals, n, limit)
	if len(solutions) == 0 {
		fmt.Println("EXCEPTION", n, k)
	}
	return solutions
}

func sortedByBeggining(sol [][]int) [][]int {
	cloned := deepClone(sol)
	sort.Slice(cloned, func(i, j int) bool {
		return cloned[i][len(cloned[i])-1] < cloned[j][len(cloned[j])-1]
	})
	return cloned
}

func isUnevenSimpleSolution(sol [][]int) bool {
	for _, group := range sol {
		if len(group) != 2 {
			return false
		}
		if group[0] == group[1] + 2 {
			return true
		}
	}
	return false
}

func hasUnevenSimpleSolution(n, k int) (bool, int) {
	sols := splitSum(n, k, 1)
	if len(sols) == 0 {
		return false, 0
	}
	return isUnevenSimpleSolution(sols[0]), sols[0][0][1]-1
}

func showSols(n, k, limit int) {
	printSols(splitSum(n, k, limit), n)
}

func showUnevenSols(n, k, target, limit int) {
	printSols(unevenSplitSum(n, k, target, limit), n)
}
