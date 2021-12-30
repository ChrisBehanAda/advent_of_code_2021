package day16

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInput("./day16/input.txt")
	binary := hexToBinary(input)
	versions, values, _ := packetVersions(binary, false)
	verSum := sum(versions)
	fmt.Printf("Sum of versions is: %v\n", verSum)
	fmt.Printf("Values: %v\n", values)
}

func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(data), "\n")
	return input
}

func hexToBinary(hex string) string {
	var binary = map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	var sb strings.Builder
	for _, r := range hex {
		b := binary[string(r)]
		sb.WriteString(b)
	}
	return sb.String()
}

func packetVersions(input string, parseOne bool) ([]int, []int, string) {
	versions := []int{}
	values := []int{}
	for len(input) > 0 {
		currentVer := getVersion(input)
		currentLabel := getLabel(input)
		versions = append(versions, int(currentVer))
		input = input[6:]
		isLiteral := currentLabel == 4
		if isLiteral {
			val, litLen := parseLiteral(input)
			values = append(values, val)
			input = input[litLen:]
		} else {
			lengthType := input[0]
			input = input[1:]
			var lengthBits int
			var subPacketVals []int
			if lengthType == '0' {
				lengthBits = 15
				length, _ := strconv.ParseInt(input[:lengthBits], 2, 64)
				input = input[lengthBits:]
				subPacketVers, vals, _ := packetVersions(input[:length], false)
				versions = append(versions, subPacketVers...)
				input = input[length:]
				subPacketVals = vals
			} else {
				lengthBits = 11
				length, _ := strconv.ParseInt(input[:lengthBits], 2, 64)
				input = input[lengthBits:]
				for i := 0; i < int(length); i++ {
					subPacketVers, vals, i := packetVersions(input, true)
					input = i
					versions = append(versions, subPacketVers...)
					subPacketVals = append(subPacketVals, vals...)
				}
			}
			if currentLabel == 0 {
				sums := packetSums(subPacketVals)
				values = append(values, sums)
			} else if currentLabel == 1 {
				prods := packetProducts(subPacketVals)
				values = append(values, prods)
			} else if currentLabel == 2 {
				min := packetMin(subPacketVals)
				values = append(values, min)
			} else if currentLabel == 3 {
				max := packetMax(subPacketVals)
				values = append(values, max)
			} else if currentLabel == 5 {
				gt := gtPacket(subPacketVals)
				values = append(values, gt)
			} else if currentLabel == 6 {
				lt := ltPacket(subPacketVals)
				values = append(values, lt)
			} else if currentLabel == 7 {
				eq := eqPacket(subPacketVals)
				values = append(values, eq)
			}

		}

		if isZero(input) || parseOne {
			break
		}

	}
	return versions, values, input
}

func packetSums(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func packetProducts(nums []int) int {
	product := 1
	for _, n := range nums {
		product *= n
	}
	return product
}

func packetMin(nums []int) int {
	small := math.MaxInt
	for _, n := range nums {
		if n < small {
			small = n
		}
	}
	return small
}

func packetMax(nums []int) int {
	max := math.MinInt
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func gtPacket(nums []int) int {
	if nums[0] > nums[1] {
		return 1
	}
	return 0
}

func ltPacket(nums []int) int {
	if nums[0] < nums[1] {
		return 1
	}
	return 0
}

func eqPacket(nums []int) int {
	if nums[0] == nums[1] {
		return 1
	}
	return 0
}

func parseLiteral(bits string) (int, int) {
	numString := ""
	length := 0
	for {
		val := bits[:5]
		if val[0] == '0' {
			num := bits[1:5]
			numString += string(num)
			bits = bits[5:]
			length += 5
			break
		} else {
			num := bits[1:5]
			numString += string(num)
			bits = bits[5:]
			length += 5
		}
	}
	num, _ := strconv.ParseInt(numString, 2, 64)
	return int(num), length
}

func isZero(input string) bool {
	num, _ := strconv.ParseInt(input, 2, 64)
	return num == 0
}

func getVersion(input string) int64 {
	s := input[:3]
	ver, _ := strconv.ParseInt(s, 2, 64)
	return ver
}

func getLabel(input string) int64 {
	s := input[3:6]
	ver, _ := strconv.ParseInt(s, 2, 64)
	return ver
}

func removeLeading0s(input string) string {
	for {
		if len(input) > 0 && input[0] == '0' {
			input = input[1:]
		} else {
			break
		}
	}
	return input
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
