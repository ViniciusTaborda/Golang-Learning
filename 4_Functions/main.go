package main

import "fmt"

func sum(n1 int8, n2 int8) (int8, error) {
	return n1 + n2, nil
}

func math_calcs(n1, n2 int8) (int8, int8, int8, int8) {
	return (n1 + n2), (n1 - n2), (n1 * n2), (n1 / n2)
}

func main() {
	sum_result, err := sum(10, 20)
	if err == nil {
		fmt.Println(sum_result)
	}

	var var_as_func = func(text string) string {
		fmt.Println("Anon func." + text)
		return text
	}
	fmt.Println(var_as_func(""))

	result_sum, _, result_mult, result_div := math_calcs(10, 10)
	fmt.Println(result_sum, result_mult, result_div)

}
