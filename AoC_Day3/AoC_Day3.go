package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	readFile, err := os.Open("Day3_input.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	fileScan := bufio.NewScanner(readFile)
	fileScan.Split(bufio.ScanLines)

	sum := 0
	var fileLines []string
	for fileScan.Scan() {
		fileLines = append(fileLines, fileScan.Text())
	}

	for i := 0; i < len(fileLines); i += 3 {
		for y := 0; y < len(fileLines[i]); y++ {
			if strings.Contains(fileLines[i+1], string(fileLines[i][y])) && strings.Contains(fileLines[i+2], string(fileLines[i][y])) {
				sum += strings.Index(letters, string(fileLines[i][y])) + 1
				break
			}
		}
	}

	fmt.Println(sum)
}

/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	file, err := os.Open("Day3_input.txt")
	if err != nil {
		panic(err)
	}
	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	sum := 0
	for fileScan.Scan() {
		mid := len(fileScan.Text()) / 2
		fir := fileScan.Text()[0:mid]
		sec := fileScan.Text()[mid:]

		for i := 0; i < len(fir); i++ {
			if strings.Contains(sec, string(fir[i])) {
				sum += strings.Index(letters, string(fir[i])) + 1
				break
			}
		}
	}

	file.Close()
	fmt.Println(sum)
}
*/
