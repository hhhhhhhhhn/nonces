function pickN<T>(array: T[], n: number): T[] {
	return array.sort(() => Math.random() - 0.5)
		.sort(() => Math.random() - 0.5)
		.slice(0, n)
}

function average(array: number[]): number {
	return array.reduce((a, b) => a + b) / array.length
}

function nintiethPercentile(array: number[]): number {
	let sorted = [...array].sort()
	return sorted[Math.floor(sorted.length * 0.9)]
}

const BLANK = " "

type Board = {
	numbers: (number|typeof BLANK)[][]
	used: boolean[][]
}

function newBoard(): Board {
	let board: Board = {
		numbers: [[], [], []],
		used: [[], [], []],
	}

	for (let i = 0; i < 9; i++) {
		let numbers = pickN([0, 1, 2, 3, 4, 5, 6, 7, 8, 9], 3)

		board.numbers[0].push(i*10 + numbers[0])
		board.numbers[1].push(i*10 + numbers[1])
		board.numbers[2].push(i*10 + numbers[2])

		board.used[0].push(false)
		board.used[1].push(false)
		board.used[2].push(false)
	}

	for (let i = 0; i < board.numbers.length; i++) {
		let blankIndices = pickN([0, 1, 2, 3, 4, 5, 6, 7, 8], 4)

		board.numbers[i][blankIndices[0]] = BLANK
		board.numbers[i][blankIndices[1]] = BLANK
		board.numbers[i][blankIndices[2]] = BLANK
		board.numbers[i][blankIndices[3]] = BLANK

		board.used[i][blankIndices[0]] = true
		board.used[i][blankIndices[1]] = true
		board.used[i][blankIndices[2]] = true
		board.used[i][blankIndices[3]] = true
	}

	assert(board.numbers.length == 3)
	assert(board.used.length == 3)
	assert(board.numbers[0].length == 9)
	assert(board.used[0].length == board.used[1].length && board.used[0].length == board.used[2].length)
	assert(board.numbers[0].length == board.numbers[1].length && board.numbers[0].length == board.numbers[2].length)
	assert(board.numbers[0].length == board.used[0].length)

	return board
}

function boardAddNumber(board: Board, number: number) {
	for (let i = 0; i < board.numbers.length; i++) {
		let j = board.numbers[i].findIndex(n => n === number)

		if (j != -1) {
			board.used[i][j] = true
		}
	}
}

function boardHasPick(board: Board): boolean {
	for (let i = 0; i < board.used.length; i++) {
		for (let j = 0; j < board.used[0].length; j++) {
			if (board.used[i][j] && board.numbers[i][j] != BLANK) {
				return true
			}
		}
	}
	return false
}

function boardHasLine(board: Board): boolean {
	for (let i = 0; i < board.used.length; i++) {
		if (board.used[i].every(e => e)) {
			return true
		}
	}
	return false
}

function boardHasLota(board: Board): boolean {
	return board.used.flat().every(e => e)
}

function assert(condition: boolean) {
	if (!condition) {
		throw new Error("Assertion failed")
	}
}

function printBoard(board: Board) {
	for (let i = 0; i < board.numbers.length; i++) {
		let row = board.numbers[i].map(n => String(n).padStart(3, " ")).join(" ")
		console.log(row)
		let used = board.used[i].map(e => e ? " ## " : " __ ").join("")
		console.log(used)
		console.log()
	}
}

function randomNumbers(): number[] {
	return Array(90).fill(0).map((_, i: number) => i)
		.sort(() => Math.random() - 0.5)
		.sort(() => Math.random() - 0.5)
}

type Game = {
	numbers: number[]
	boards: Board[]
	isPlayed: boolean
	firstPickTurn: number
	firstPickBoards: number
	firstLineTurn: number
	firstLineBoards: number
	firstLotaTurn: number
	firstLotaBoards: number
}

function newGame(boards: number): Game {
	let game: Game = {
		isPlayed: false,
		numbers: randomNumbers(),
		boards: Array(boards).fill(0).map(() => newBoard()),
		firstPickTurn: -1,
		firstLineTurn: -1,
		firstLotaTurn: -1,
		firstPickBoards: -1,
		firstLineBoards: -1,
		firstLotaBoards: -1,
	}
	return game
}

function playGame(game: Game): void {
	if (game.isPlayed) {
		return
	}
	game.isPlayed = true
	
	let turn = 1
	for (let number of game.numbers) {
		game.boards.forEach(b => boardAddNumber(b, number))

		if (game.firstPickTurn == -1) {
			let amount = game.boards.filter(b => boardHasPick(b)).length
			if (amount > 0) {
				game.firstPickTurn = turn
				game.firstPickBoards = amount
			}
		}
		else if (game.firstLineTurn == -1) {
			let amount = game.boards.filter(b => boardHasLine(b)).length
			if (amount > 0) {
				game.firstLineTurn = turn
				game.firstLineBoards = amount
			}
		}
		else if (game.firstLotaTurn == -1) {
			let amount = game.boards.filter(b => boardHasLota(b)).length
			if (amount > 0) {
				game.firstLotaTurn = turn
				game.firstLotaBoards = amount
			}
		}
		else {
			return
		}

		turn++
	}
}


function printStatistics(boards: number) {
	let firstPickTurns: number[] = []
	let firstPickBoards: number[] = []
	let firstLineTurns: number[] = []
	let firstLineBoards: number[] = []
	let firstLotaTurns: number[] = []
	let firstLotaBoards: number[] = []

	for (let i = 0; i < RUNS; i++) {
		let game = newGame(boards)
		playGame(game)
		firstPickTurns.push(game.firstPickTurn)
		firstPickBoards.push(game.firstPickBoards)
		firstLineTurns.push(game.firstLineTurn)
		firstLineBoards.push(game.firstLineBoards)
		firstLotaTurns.push(game.firstLotaTurn)
		firstLotaBoards.push(game.firstLotaBoards)
	}

	let line = [
		boards,
		average(firstPickTurns),
		nintiethPercentile(firstPickTurns),
		average(firstPickBoards),
		average(firstLineTurns),
		nintiethPercentile(firstLineTurns),
		average(firstLineBoards),
		average(firstLotaTurns),
		nintiethPercentile(firstLotaTurns),
		average(firstLotaBoards),
		firstLineBoards.filter(e => e > 1).length / RUNS,
		firstLotaBoards.filter(e => e > 1).length / RUNS,
	].map(e => String(e)).join(",")

	console.log(line)
}

function printHeader() {
	let line = [
		"Boards",
		"First Pick Turns",
		"First Pick 90th",
		"First Pick Boards",
		"First Line Turns",
		"First Line 90th",
		"First Line Boards",
		"First Lota Turns",
		"First Lota 90th",
		"First Lota Boards",
		"First Line % >1 Board",
		"First Lota % >1 Board",
	].join(",")
	console.log(line)
}
const RUNS = 5000

printHeader()
printStatistics(1)
printStatistics(2)
printStatistics(5)
for (let i = 10; i <= 250; i += 20) {
	printStatistics(i)
}

// let board = newBoard()
// printBoard(board)
// 
// let numbers = randomNumbers()
// for (let i = 0; i < 80; i++) {
// 	console.log(`================ ${numbers[i]} ===============`)
// 	boardAddNumber(board, numbers[i])
// 	printBoard(board)
// 	console.log(`Pick: ${boardHasPick(board)}, Line: ${boardHasLine(board)}, Lota: ${boardHasLota(board)}`)
// 	console.log()
// 	console.log()
// }
