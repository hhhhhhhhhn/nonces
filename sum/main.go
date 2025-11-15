package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"image"
	"image/png"
	"image/color"
)


type Key struct {
	goals [MAX_K]int
	max int
}

var memoTable = map[Key]bool{}

func sumPossible(goals [MAX_K]int, max int) bool {
	if max < 0 {
		return false
	}
	key := Key{goals, max}
	isReady := true
	sum := 0
	for i := 0; i < MAX_K; i++ {
		if goals[i] < 0 {
			return false
		}
		if goals[i] != 0 {
			isReady = false
			sum += goals[i]
		}
	}
	if isReady {
		return true
	}
	if sum > 10 {
		if result, ok := memoTable[key]; ok {
			return result
		}
	}
	result := false
	for i := 0; i < MAX_K; i++ {
		if goals[i] <= max-1 {
			continue
		}
		var newGoals [MAX_K]int
		copy(newGoals[:], goals[:])
		newGoals[i] -= max
		if sumPossible(newGoals, max-1) {
			result = true
			break
		}
	}
	memoTable[key] = result
	return result
}

func canSplitIntoNGroups(max int, k int) bool {
	sum := max*(max+1)/2

	if sum % k != 0 {
		return false
	}

	target := sum / k

	if target < max {
		return false
	}


//	// Special case, can be split into groups of n+1
//	// e.g. 1 2 3 4 -> (1+4), (2+3)
//	if max % 2 == 0 && target % (max+1) == 0 {
//		return true
//	}
//
//	// Other special case, can be split into groups of n
//	// e.g. 1 2 3 4 5 -> (1+4), (2+3), 5
//	if max % 2 == 1 && target % max == 0 {
//		return true
//	}
//
//	// Unproved
//	if max % 4 == 0 && 2*target % max == 0 {
//		return true
//	}
	if proof(max, k) {
		return true
	}

	goals := [MAX_K]int{}
	for i := 0; i < k; i++ {
		goals[i] = target
	}

	goals[0] -= max
	result := sumPossible(goals, max-1)
	if result == false {
		fmt.Println("=============== FAILED BY SEARCH ===========")
	}
	return result
}

var cpuProfFile *os.File
func toggleCpuProf() {
	if cpuProfFile == nil {
		cpuProfFile, _ = os.Create("cpuprof")
		pprof.StartCPUProfile(cpuProfFile)
	} else {
		pprof.StopCPUProfile()
		cpuProfFile.Close()
		cpuProfFile = nil
	}
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func proof(n, k int) bool {
	target := n*(n+1)/2/k
	heuristic := n*(n+1)/2 % k == 0 && k < n && target >= n
	proof := k % 2 == 1 && (n % k == 0 || (n+1) % k == 0)
	proof = proof || (k%2 == 0 && ((n)%(2*k) == 0 || (n+1)%(2*k) == 0))
	return proof && heuristic
}

func heuristicf(n, k int) bool {
	target := n*(n+1)/2/k
	return n*(n+1)/2 % k == 0 && k < n && target >= n
}

func fromSimplerCase(n, k int) (bool, int) {
	target := n*(n+1)/2/k
	if !heuristicf(n, k) {
		return false, 0
	}
	for lower := 1; lower < k; lower++ {
		newN := n - 2*lower
		newK := k - lower
		if !heuristicf(newN, newK) {
			continue
		}
		newTarget := newN*(newN+1)/2/newK
		if newTarget == target {
			return true, lower
		}
	}
	return false, 0
}

func unevenFixable(n, k int) (bool, int, int) {
	if n*(n+1) % (2*k) != 0 {
		return false, 0, 0
	}
	target := n*(n+1)/2/k
	if target % 2 != 0 {
		return false, 0, 0
	}

	newN := (n-2*k)*(n+1)/2/k
	newHalfK := k - (n - newN + 1)/2

	newK := 2*newHalfK + 1

	return heuristicf(newN, newK), newN, newK
}

const SIZE = 90
func generate_theoric_image() {
	image := image.NewRGBA64(image.Rect(0, 0, SIZE, SIZE))

	for r := 0; r < SIZE; r++ {
		for c := 0; c < SIZE; c++ {
			k := r+1
			n := c+1
			target := n*(n+1)/2/k
			heuristic := n*(n+1)/2 % k == 0 && k < n && target >= n
			proof := k % 2 == 1 && (n % k == 0 || (n+1) % k == 0)
			proof = proof || (k%2 == 0 && ((n)%(2*k) == 0 || (n+1)%(2*k) == 0))

			simplifyable, factor := fromSimplerCase(n, k)
			var hasUneven bool
			var unevenN int

			//proof = proof || simplifyable

			//if heuristic && !proof && !simplifyable {
			//	hasUneven, unevenN = hasUnevenSimpleSolution(n, k)
			//}

			red := uint8(0)
			blue := uint8(0)
			green := uint8(0)

			hasUneven, unevenN, unevenK := unevenFixable(n, k)


			if heuristic {
				red = 255
			}
			if proof {
				blue = 255
			}
			if simplifyable || (n >= 4*k && heuristic) || hasUneven {
				green = 255
			}

			if heuristic && !proof {
				m := n % (2*k) + 2*k
				if heuristicf(m, k) && m != n{
					fmt.Print("Repeat ", n, k)
				} else {
					fmt.Print(n, k)
				}
				if simplifyable {
					newK := k - factor
					newN := n - 2*factor
					fmt.Print(" <- Even ", newN, newK, " (of ", target, ")")
				} else if hasUneven {
					fmt.Print(" <- Uneven ", unevenN, unevenK, " (of ", target, ")")
				}
				fmt.Println()
			}

			image.Set(c, r, color.RGBA{red, green, blue, 255})
		}
	}

	f, _ := os.Create("theoric.png")
	png.Encode(f, image)
	f.Close()
}

func cap(x int) uint8 {
	if x > 255 {
		return 255
	}
	return uint8(x)
}

func possible(n int, k int) bool{
	return n*(n+1) % (2*k) == 0 && n*(n+1)/2/k >= n
}

const CHAIN_SIZE int = 150
func chainImageMain() {
	image := image.NewRGBA64(image.Rect(0, 0, CHAIN_SIZE/2, CHAIN_SIZE))

	for r := 0; r < CHAIN_SIZE; r++ {
		for c := 0; c < CHAIN_SIZE/2; c++ {
			k := c+1
			n := r+1

			target := n*(n+1)/2/k

			if (!possible(n, k)) {
				image.Set(c, r, color.RGBA{0, 0, 0, 255})
			} else if n == 2*k || n == 2*k-1 {
				image.Set(c, r, color.RGBA{255, 255, 255, 255})
			} else if n >= 4*k - 1 {
				image.Set(c, r, color.RGBA{255, 255, 0, 255})
			} else if target % 2 == 0 {
				image.Set(c, r, color.RGBA{0, 255, 0, 255})
			} else {
				image.Set(c, r, color.RGBA{0, 0, 255, 255})
			}

		}
	}

	f, _ := os.Create("chain.png")
	png.Encode(f, image)
	f.Close()
}

func chainLengthImageMain() {
	image := image.NewRGBA64(image.Rect(0, 0, CHAIN_SIZE/2, CHAIN_SIZE))

	for r := 0; r < CHAIN_SIZE; r++ {
		for c := 0; c < CHAIN_SIZE/2; c++ {
			k := c+1
			n := r+1

			if (!possible(n, k)) {
				image.Set(c, r, color.RGBA{0, 0, 0, 255})
			} else {
				length := chain(n, k)
				if length == 1 {
					image.Set(c, r, color.RGBA{100, 0, 0, 255})
				} else if length == 2 {
					image.Set(c, r, color.RGBA{255, 0, 0, 255})
				} else if length == 3 {
					image.Set(c, r, color.RGBA{255, 100, 0, 255})
				} else if length == 4{
					image.Set(c, r, color.RGBA{255, 255, 100, 255})
				} else if length == 5 {
					image.Set(c, r, color.RGBA{255, 255, 255, 255})
				} else {
					image.Set(c, r, color.RGBA{255, 0, 255, 255})
				}
			}

		}
	}

	f, _ := os.Create("chainlength.png")
	png.Encode(f, image)
	f.Close()
}

func chain(n int, k int) int {
	if (!possible(n, k)) {
		fmt.Println("CHAIN IMPOSSIBLE")
		os.Exit(1)
	}
	if (n == 2*k || n == 2*k-1) {
		return 1
	}
	if (n >= 4*k-1) {
		return 0 + chain(n-2*k, k)
	}
	t := n*(n+1)/2/k
	if t % 2 == 0 {
		newN := t - (n+1)
		newK := 2*k - 2*n + t - 1
		return 1 + chain(newN, newK)
	} else {
		newN := t - (n+1)
		newK := k - n + (t+1)/2 - 1
		return 1 + chain(newN, newK)
	}
}
const MAX_K = 10
func imageMain() {
	image := image.NewRGBA64(image.Rect(0, 0, MAX_K, MAX_K))
	toggleCpuProf()

	unexpected := 0

	slow := [][3]int{}

	for n := 1; n <= MAX_K; n++ {
		for k := 1; k <= MAX_K; k++ {
			can := canSplitIntoNGroups(n, k)
			target := n*(n+1)/2/k
			fmt.Printf("n=%v, k=%v: %v\n", n, k, can)
			fmt.Println(len(memoTable), "entries")
			fmt.Println("Target:", target)
			heuristic := n*(n+1)/2 % k == 0 && (k < n || n == 1) && target >= n
			if can {
				pixel := len(memoTable)
				image.Set(k-1, n-1, color.RGBA{cap(pixel), cap(pixel/100), cap(pixel/10000), 255})
			} else {
				image.Set(k-1, n-1, color.Black)
			}

			if (!can && heuristic) || (can && !heuristic) {
				fmt.Println("============== UNEXPECTED =============")
				image.Set(k-1, n-1, color.RGBA{255, 0, 255, 255})
				unexpected++
			}

			if len(memoTable) > 10000 {
				slow = append(slow, [3]int{n, k, len(memoTable)})
			}

			fmt.Println()
			memoTable = map[Key]bool{}
		}
	}
	fmt.Println("Unexpected:", unexpected)
	toggleCpuProf()

	fmt.Println(slow)

	f, _ := os.Create("image.png")
	png.Encode(f, image)
	f.Close()

	generate_theoric_image()
}

func main() {
	//chainImageMain()
	//chainLengthImageMain()
	// fmt.Println(unevenFixable(15, 6))
	//fmt.Println(hasUnevenSimpleSolution(15, 9))
	// imageMain()
	n := 2024
	k := 690
	showSols(n, k, 1)
	//_, unevenN := hasUnevenSimpleSolution(n, k)
	//newK := k - ((n - unevenN)-1)/2
	//showUnevenSols(unevenN, newK, n*(n+1)/2/k, 1)
}
