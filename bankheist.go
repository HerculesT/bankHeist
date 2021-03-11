package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		heistSuccess(true) //give bool to decide the console color
	} else {
		isHeistOn = false
		heistFail(false)
	}

	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		vaultSuccess(true)
	} else if isHeistOn {
		isHeistOn = false
		vaultFail(false)
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {

		switch leftSafely {
		case 0:
			isHeistOn = false
			leaveFail(false)
		case 1:
			isHeistOn = false
			leaveFail(false)
		case 2:
			isHeistOn = false
			leaveFail(false)
		default:
			leaveSuccess(true)
		}
	}
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("The Stolen amount on the last nights robbery is: ", amtStolen)
	}
	// fmt.Println(isHeistOn)
}

func robberyResult(fileName string, x bool) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[0;32m"

	//reset the seed time to get new results
	//without it you'll always get the same result
	rand.Seed(time.Now().Unix())
	success, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(success)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		//     // Append line to result.
		result = append(result, line)
	}
	n := rand.Int() % len(result)
	sHeist := result[n]
	//check if the boolean comes true
	if x != false {
		fmt.Print((colorGreen), sHeist, "\n", (colorReset))
	} else {
		fmt.Print((colorRed), sHeist, "\n", (colorReset))
	}
}

func heistSuccess(x bool) {
	robberyResult("src/heistSuccess.txt", true)
}
func heistFail(x bool) {
	robberyResult("src/heistFail.txt", false)
}
func vaultSuccess(x bool) {
	robberyResult("src/vaultsuccess.txt", true)
}
func vaultFail(x bool) {
	robberyResult("src/vaultFail.txt", false)
}
func leaveSuccess(x bool) {
	robberyResult("src/leaveSuccess.txt", true)
}
func leaveFail(x bool) {
	robberyResult("src/leaveFail.txt", false)
}
