// sorter.go
// 从命令行输入排序文件、输出文件、指定排序算法
// USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>

package main

import (
	"algorithms/bubblesort" // import子目录名字
	"algorithms/quicksort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	//	"testpkg/testpkg_in"
	"time"
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

// 输入文件中读取，返回一个slice
func readInValues(infile string) (values []int, err error) {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)

	}
	return

}

func writeOutValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create file ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readInValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
		t1 := time.Now()
		switch *algorithm {
		case "bubblesort":
			MyBubbleSort.BubbleSort(values)
			fmt.Println("Write values:", values)
			writeOutValues(values, *outfile)

		case "qsort":
			quicksort.QuickSort(values)
			fmt.Println("Write values:", values)
			writeOutValues(values, *outfile)

		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("took ", t2.Sub(t1))

	} else {
		fmt.Println(err)
		return
	}

}
