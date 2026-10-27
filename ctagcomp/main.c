#include <stdio.h>
#include <stdbool.h>

#define Result(type) struct type##Result {bool ok; union {type value; int error;};}
#define Error(type, err) (Result(type)){.ok = false, .error = err}
#define Ok(type, val) (Result(type)){.ok = true, .value = val}

#define DIV_BY_ZERO 1

Result(int) divide(int a, int b) {
	if (b == 0) {
		return Error(int, DIV_BY_ZERO);
	}
	return Ok(int, a/b);
}

int main() {
	Result(int) result = divide(5,2);
	printf("%u\n", result.value);
}
