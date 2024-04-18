package main

import ("fmt"
 "regexp"
  "math/big"
  "math"
  "sort")

func main() {
	nick := "markor";
	fmt.Println("Twoja silna liczba to: ", silnaLiczba(nick))
	fmt.Println("Twoja słaba liczba to: ", slabaLiczba(silnaLiczba(nick)))
}	

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(n)
	return result.Mul(result, factorial(n-1))
}

func silnaLiczba(nick string) int64{
	nickArray := make([]byte, len(nick))
    for i, char := range nick {
        nickArray[i] = byte(char)
    }
	regexStr := ""
    for _, value := range nickArray {
        regexStr += fmt.Sprintf(`\d*%d`, value)
    }
    re := regexp.MustCompile(regexStr)
    strongNumber := int64(0)
	for i := int64(1); ; i++ {
        result := factorial(i).String()
        if re.MatchString(result) {
            strongNumber = i
            break
        }
    }
    return strongNumber
}

func fibonacciCounter(n int64) (int64, map[int64]int64) {
    callCounts := make(map[int64]int64)
    var fib func(int64) int64
    fib = func(n int64) int64 {
        callCounts[n]++
        if n <= 1 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    result := fib(n)
    return result, callCounts
}

func slabaLiczba(silnaLiczba int64) int64{
    _, callCounts := fibonacciCounter(30)
    differences := []int64{}
    for i := int64(30); i >= 1; i-- {
        diff := int64(math.Abs(float64(silnaLiczba - callCounts[i])))
        differences = append(differences, diff)
    }
    sort.Slice(differences, func(i, j int) bool { return differences[i] < differences[j] })
    minDiff := differences[0]
	var res int64;
    for i := int64(30); i >= 1; i-- {
        if minDiff == int64(math.Abs(float64(silnaLiczba - callCounts[i]))) {
			res = i
        }
    }
	return res
}