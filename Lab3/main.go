package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	a := 55
	b := 74
	c := 113
	fmt.Println(DoubleNod(a, b))
	fmt.Println(TripleNod(a, b, c))

	m, n := 2, 591
	fmt.Println(float64(n) / math.Log(float64(n)))

	fmt.Println()
	simpleDigits := SimpleDigits(m, n)
	fmt.Println(simpleDigits)
	fmt.Println(len(simpleDigits))

	fmt.Println()
	simpleDigitsEratosfen := sieveOfEratosthenes(n)
	fmt.Println(simpleDigitsEratosfen)
	fmt.Println(len(simpleDigitsEratosfen))
}

func DoubleNod(a, b int) int {
	if b == 0 {
		return a
	}
	return DoubleNod(b, a%b)
}

func TripleNod(a, b, c int) int {
	return DoubleNod(DoubleNod(a, b), c)
}

func SimpleDigits(m, n int) []int {
	t1 := time.Now()
	var result []int
	for i := m; i < n; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	fmt.Println(time.Since(t1).Microseconds())

	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func sieveOfEratosthenes(n int) []int {
	t1 := time.Now()
	// Инициализируем массив, где все числа от 2 до n помечены как простые
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Проходим по всем числам от 2 до sqrt(n) и помечаем все кратные как составные
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Создаем слайс из простых чисел
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	fmt.Println(time.Since(t1).Microseconds())

	return primes
}
