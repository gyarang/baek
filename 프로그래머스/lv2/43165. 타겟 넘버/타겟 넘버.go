package main

type counter struct {
	numbers []int
	target  int
	count   int
}

func NewCounter(numbers []int, target int) counter {
	c := counter{
		count:   0,
		numbers: numbers,
		target:  target,
	}
	return c
}

func (c *counter) dfs(depth, sum int) {
	if depth == len(c.numbers) {
		if sum == c.target {
			c.count += 1
		}
		return
	}

	for i := 0; i < 2; i++ {
		new_sum := 0
		if i == 1 {
			new_sum = sum - c.numbers[depth]
		} else {
			new_sum = sum + c.numbers[depth]
		}
		c.dfs(depth+1, new_sum)
	}
}

func solution(numbers []int, target int) int {
	c := NewCounter(numbers, target)
	c.dfs(0, 0)
	return c.count
}