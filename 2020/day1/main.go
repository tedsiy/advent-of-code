package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filePath := flag.String("f", "input", "get int array from `file`")
	if len(os.Args) < 2 {
		flag.Usage()
	}
	flag.Parse()

	result, err := getFileOfInts(*filePath)
	if err != nil {
		log.Fatalln(err)
	}

	x, err := getMultiply2ThatAddTo2020(result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Multiplication of the two number that equal 2020\n\t%d\n", x)

	x, err = getMultiply3ThatAddTo2020(result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Multiplication of the three number that equal 2020\n\t%d\n", x)

}

func getFileOfInts(filePath string) (result []int, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, nil
}

func getMultiply3ThatAddTo2020(stuff []int) (int, error) {

	for _, x := range stuff {
		for _, y := range stuff {
			for _, z := range stuff {
				if 2020 == (x + y + z) {
					log.Printf(fmt.Sprintf("%d, %d, %d\n", x, y, z))
					return x * y * z, nil
				}
			}
		}
	}
	return 0, errors.New("no value found")
}

func getMultiply2ThatAddTo2020(stuff []int) (int, error) {

	for _, x := range stuff {
		for _, y := range stuff {
			if 2020 == (x + y) {
				return x * y, nil
			}
		}
	}
	return 0, errors.New("no value found")
}
