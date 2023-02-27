package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
	engAlphabet = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	str := "Hello my name is Denis Leonov"
	data := []byte(str)
	res := encode64(data)
	fmt.Println(res)

	engEntropy := entropy(engAlphabet, str, "")
	base64Entropy := entropy(base64Table, res, "")

	engEntropy2 := math.Log2(float64(len(engAlphabet)))
	base64Entropy2 := math.Log2(float64(len(base64Table)))

	engIzb := (engEntropy2 - engEntropy) / engEntropy2 * 100
	base64Izb := (base64Entropy2 - base64Entropy) / base64Entropy2 * 100

	fmt.Printf("english alphabet izb: %f\nbase64 alphabet izb: %f\n", engIzb, base64Izb)

	str1 := "Leonov"
	str2 := "Denis"

	resXor, resDoubleXor := XorOperation([]byte(str1), []byte(str2))
	fmt.Println(resXor)
	fmt.Println(resDoubleXor)
}

func encode64(data []byte) string {
	var index = 0
	var binarySequnce = ""
	output := make([]string, 0)
	var result = ""
	for index < len(data) {
		for j := index; j < index+3; j++ {
			if j >= len(data) {
				continue
			} else {
				binarySequnce += fmt.Sprintf("%08b", data[j])
				//fmt.Println(fmt.Sprintf("%08b", data[j]))
			}
		}
		//fmt.Printf("Binary Sequence: %s\n", binarySequnce)
		//fmt.Printf("Length of Binary Sequence: %d\n", len(binarySequnce))

		switch len(binarySequnce) {
		case 24:
			for k := 0; k < 4; k++ {
				temp := binarySequnce[6*k : 6*k+6]
				output = append(output, temp)
			}
			break
		case 16:
			output = append(output, binarySequnce[0:6])
			output = append(output, binarySequnce[6:12])
			output = append(output, binarySequnce[12:16]+"00")
			output = append(output, "=")
			break
		case 8:
			output = append(output, binarySequnce[0:6])
			output = append(output, binarySequnce[6:8]+"0000")
			output = append(output, "=")
			output = append(output, "=")
			break
		}
		//fmt.Println(output)

		binarySequnce = ""
		index += 3
	}

	for _, v := range output {
		if v == "=" {
			result += "="
		} else {
			num, _ := strconv.ParseUint(v, 2, 64)
			result += string(base64Table[num])
		}
	}
	return result
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

	entropy := 0.0
	for _, p := range probability {
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}

	return entropy
}

func XorOperation(str1 []byte, str2 []byte) (string, string) {
	var binaryStr1 = ""
	var binaryStr2 = ""

	for i := 0; i < len(str1); i++ {
		binaryStr1 += fmt.Sprintf("%08b", str1[i])
	}
	for i := 0; i < len(str2); i++ {
		binaryStr2 += fmt.Sprintf("%08b", str2[i])
	}
	//fmt.Println(binaryStr1)
	//fmt.Println(binaryStr2)

	if len(binaryStr1) > len(binaryStr2) {
		for len(binaryStr2) != len(binaryStr1) {
			binaryStr2 += "0"
		}
	} else {
		for len(binaryStr1) != len(binaryStr2) {
			binaryStr1 += "0"
		}
	}

	//fmt.Println(binaryStr1)
	//fmt.Println(binaryStr2)

	result := ""
	for i := 0; i < len(binaryStr1); i++ {
		switch {
		case string(binaryStr1[i]) == "0" && string(binaryStr2[i]) == "0":
			result += "0"
			break
		case string(binaryStr1[i]) == "0" && string(binaryStr2[i]) == "1":
			result += "1"
			break
		case string(binaryStr1[i]) == "1" && string(binaryStr2[i]) == "0":
			result += "1"
			break
		case string(binaryStr1[i]) == "1" && string(binaryStr2[i]) == "1":
			result += "0"
			break
		}
	}

	doubleXorResult := ""
	for i := 0; i < len(result); i++ {
		switch {
		case string(result[i]) == "0" && string(binaryStr2[i]) == "0":
			doubleXorResult += "0"
			break
		case string(result[i]) == "0" && string(binaryStr2[i]) == "1":
			doubleXorResult += "1"
			break
		case string(result[i]) == "1" && string(binaryStr2[i]) == "0":
			doubleXorResult += "1"
			break
		case string(result[i]) == "1" && string(binaryStr2[i]) == "1":
			doubleXorResult += "0"
			break
		}
	}

	fmt.Println(result)
	fmt.Println(doubleXorResult)

	var utf8Runes []rune
	for i := 0; i < len(result); i += 8 {
		binaryChunk := result[i : i+8]
		byteVal, err := strconv.ParseUint(binaryChunk, 2, 8)
		if err != nil {
			return "", ""
		}
		utf8Runes = append(utf8Runes, rune(byteVal))
	}

	var utf8RunesForDoubleXor []rune
	for i := 0; i < len(result); i += 8 {
		binaryChunk := doubleXorResult[i : i+8]
		byteVal, err := strconv.ParseUint(binaryChunk, 2, 8)
		if err != nil {
			return "", ""
		}
		utf8RunesForDoubleXor = append(utf8RunesForDoubleXor, rune(byteVal))
	}
	return string(utf8Runes), string(utf8RunesForDoubleXor)
}
