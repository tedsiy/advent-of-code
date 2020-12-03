package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(?m)(?P<min>\d+)-(?P<max>\d+) (?P<letter>\w): (?P<word>\w+)`)

func main() {
	valid := 0
	invalid := 0

	filePath := "data/input"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		results, err := checkPassword2(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if results {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Println(fmt.Sprintf("%d valid and %d invalid", valid, invalid))

}

func checkPassword2(line string) (bool, error) {
	res := re.FindAllStringSubmatch(line, -1)[0]
	// res1 := strings.Count(res[4], res[3])
	min, err := strconv.Atoi(res[1])
	if err != nil {
		return false, err
	}
	max, err := strconv.Atoi(res[2])
	if err != nil {
		return false, err
	}
	r := []rune(res[4])    // password
	c := []rune(res[3])[0] // charater
	v1 := c == r[min-1]
	v2 := c == r[max-1]
	if (v1 || v2) && !(v1 && v2) {
		// fmt.Printf("%t, %d, %d, %s, %s\n", true, min, max, res[4], res[3])
		return true, nil
	}

	// fmt.Printf("%t, %d, %d, %s, %s\n", false, min, max, res[4], res[3])
	return false, nil
}

func part1() {
	valid := 0
	invalid := 0

	filePath := "data/input"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		results, err := checkPassword(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if results {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Println(fmt.Sprintf("%d valid and %d invalid", valid, invalid))

}

func checkPassword(line string) (bool, error) {
	res := re.FindAllStringSubmatch(line, -1)[0]
	res1 := strings.Count(res[4], res[3])
	min, err := strconv.Atoi(res[1])
	if err != nil {
		return false, err
	}
	max, err := strconv.Atoi(res[2])
	if err != nil {
		return false, err
	}
	if min <= res1 && max >= res1 {
		return true, nil
	}
	return false, nil
}
