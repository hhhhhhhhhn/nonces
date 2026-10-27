// Where each bipartition has an "in" group, and an "out" group
function bipartitions(elements, inGroupAmount) {
	if (inGroupAmount > elements.length || inGroupAmount < 0) {
		return []
	}
	if (elements.length == 0) {
		return [[[], []]]
	}
	let head = elements[0]
	let tail = elements.slice(1)
	let totalBipartitions = []

	let tailBipartitionsTakingHead = bipartitions(tail, inGroupAmount-1)
	for (let tailBipartition of tailBipartitionsTakingHead) {
		let [inGroup, outGroup] = tailBipartition
		totalBipartitions.push([[head, ...inGroup], outGroup])
	}
	let tailBipartitionsWithoutTakingHead = bipartitions(tail, inGroupAmount)
	for (let tailBipartition of tailBipartitionsWithoutTakingHead) {
		let [inGroup, outGroup] = tailBipartition
		totalBipartitions.push([inGroup, [head, ...outGroup]])
	}
	return totalBipartitions
}

function handHas24Solution(hand) {
	if (hand.length == 1) {
		return Math.abs(hand[0] - 24) < 0.01
	}
	let handBipartitions = bipartitions(hand, 2)

	for (let [operated, others] of handBipartitions) {
		if (
			handHas24Solution([operated[0] + operated[1], ...others]) ||
			handHas24Solution([operated[0] - operated[1], ...others]) ||
			handHas24Solution([operated[0] * operated[1], ...others]) ||
			handHas24Solution([operated[0] / operated[1], ...others]) ||
			handHas24Solution([operated[1] - operated[0], ...others]) ||
			handHas24Solution([operated[1] / operated[0], ...others])
		) {
			return true
		}
	}
	return false
}

function generate24Problem() {
	while(true) {
		let hand = Array(100).fill(0).map(() => Math.floor(Math.random()*12)+1)
		if (handHas24Solution(hand)) {
			return hand
		}
	}
}

function arrayEquals(a, b) {
	return a.length == b.length && a.every((_, i) => a[i] == b[i])
}

function is24SolutionHorrible(input, problem) {
	try {
		input = input.replace(/[^\(\)*-/+0-9]/g, "")
		let numbers = input.split(/[^0-9]+/).filter(Boolean).map(Number)
		return Math.abs(eval(input)-24) < 0.01 && arrayEquals(numbers.toSorted(), problem.toSorted())
	} catch (e) {
		console.error(e)
		return false
	}

}

console.log(is24SolutionHorrible("12 + 11 + 8 - 7", [12,11,7,8]))
