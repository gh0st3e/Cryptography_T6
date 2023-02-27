package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	FIO = "Leonov Denis Igorevich"
)

func main() {
	engAlphabet := "abcdefghijklmnopqrstuvwxyz"
	rusAlphabet := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	digitsAlphabet := "01"

	engText := "This is a sample text document containing Latin characters! Hello, world!"
	rusText := "Это образец текстового документа. Привет, мир!"
	digitsText := generateDigitsSequence()

	engAlphabetEntropy := entropy(engAlphabet, engAlphabet, "engAlphabet")
	rusAlphabetEntropy := entropy(rusAlphabet, rusAlphabet, "rusAlphabet")
	digitsAlphabetEntropy := entropy(digitsAlphabet, digitsAlphabet, "digitsAlphabet")
	fmt.Printf("English Alphabet Entropy = %f\n", engAlphabetEntropy)
	fmt.Printf("Russian Alphabet Entropy = %f\n", rusAlphabetEntropy)
	fmt.Printf("Digits Alphabet Entropy = %f\n", digitsAlphabetEntropy)

	engEntropy := entropy(engAlphabet, engText, "eng")
	rusEntropy := entropy(rusAlphabet, rusText, "rus")
	digitsEntropy := entropy(digitsAlphabet, digitsText, "digits")

	fmt.Printf("English Text Entropy = %f\n", engEntropy)
	fmt.Printf("Russian Text Entropy = %f\n", rusEntropy)
	fmt.Printf("Digits Text Entropy = %f\n", digitsEntropy)

	countBits(FIO, rusEntropy)
	digitsCountBits(FIO, digitsEntropy)
	digitsCountBitsWithErrors(FIO, 0.1)
	digitsCountBitsWithErrors(FIO, 0.5)
	digitsCountBitsWithErrors(FIO, 1)
}

func entropy(alphabet, text, fileName string) float64 {
	frequency := make(map[rune]int)

	for _, char := range alphabet {
		frequency[char] = 0
		for _, textChar := range text {
			if textChar == char {
				frequency[char]++
			}
		}
	}

	textLen := len(text)
	probability := make(map[rune]float64)
	for char, count := range frequency {
		probability[char] = float64(count) / float64(textLen)
	}

	saveToExcel(fileName, probability)

	entropy := 0.0
	for _, p := range probability {
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}

	return entropy
}

func generateDigitsSequence() string {
	rand.Seed(time.Now().UnixNano())

	sequence := ""
	for i := 0; i < 500; i++ {
		sequence += fmt.Sprintf("%d", rand.Intn(2))
	}

	return sequence
}

func saveToExcel(fileName string, probability map[rune]float64) {
	var index int = 2
	var cell string

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "Letter")
	f.SetCellValue("Sheet1", "B1", "Probability")

	for char, prob := range probability {
		cell = fmt.Sprintf("A%d", index)
		f.SetCellValue("Sheet1", cell, string(char))
		cell = fmt.Sprintf("B%d", index)
		f.SetCellValue("Sheet1", cell, prob)
		index++
	}

	err := f.SaveAs(fmt.Sprintf("%s.xlsx", fileName))
	if err != nil {
		println(err.Error())
		return
	}
}

func countBits(name string, entropy float64) {
	var lenName float64 = 0
	for _, v := range name {
		if v != ' ' {
			lenName++
		}
	}
	fmt.Printf("Count of bytes = %f\n", lenName*entropy)
}

func digitsCountBits(name string, entropy float64) {
	var lenName int = 0
	for _, v := range name {
		if v != ' ' {
			temp := fmt.Sprintf("%08b", v)
			lenName += len(temp)
		}
	}
	fmt.Printf("Digits count of bytes = %f\n", float64(lenName)*entropy)
}

func digitsCountBitsWithErrors(name string, errorProb float64) {
	fEntropy := -errorProb*math.Log2(errorProb) - (1-errorProb)*math.Log2(1-errorProb)
	sEntropy := 1 - fEntropy
	var lenName int = 0
	for _, v := range name {
		if v != ' ' {
			temp := fmt.Sprintf("%08b", v)
			lenName += len(temp)
		}
	}
	fmt.Printf("Count of bits with errorProb = %f is %f\n", errorProb, sEntropy*float64(lenName))

}
