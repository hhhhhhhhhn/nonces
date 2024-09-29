function possible(n: number, k: number): boolean {
	return n*(n+1) % (2*k) == 0 && target(n, k) >= n
}

function target(n: number, k: number): number {
	return n*(n+1)/2/k
}

function chain(n: number, k: number) {
	while(true) {
		console.log(n, k, target(n, k))
		if (!possible(n, k)) {
			console.log("IMPOSSIBLE")
			break
		}
		else if (k == 1) {
			console.log("CASO BASE")
			break
		}
		else if (n == 2*k || n == 2*k-1) {
			console.log("CASO BASE")
			break
		}
//		else if (k % 2 == 0 && (n == 3*k || n == 3*k-1)) {
//			console.log("BASE CASE (EVEN)")
//			break
//		}
		else if (n >= 4*k-1) {
			n = n - 2*k
			console.log("AUMENTO")
		}
		else if (n > 2*k) {
			let t = target(n, k)
			if (t % 2 == 0) {
				let newN = t - (n+1)
				let newK = 2*k - 2*n + t - 1
				n = newN
				k = newK
				console.log("AMPLIFICACION EMPAREJADA")
			} else {
				let newN = t - (n+1)
				let newK = k - n + (t+1)/2 - 1
				n = newN
				k = newK
				console.log("AMPLIFICACIÓN DIRECTA")
			}
		} else {
			console.log("CANNOT REDUCE")
			break
		}
	}
}

function main() {
//	for (let n = 0; n < 100; n++){
//		for (let k = 0; k < 4*n; k++){
//			if (possible(n, k)) chain(n, k)
//		}
//	}
	chain(2024, 690)
}
main()
