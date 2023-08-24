package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Утилита sort
//Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.
//
//Реализовать поддержку утилитой следующих ключей:
//
//-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
//-n — сортировать по числовому значению
//-r — сортировать в обратном порядке
//-u — не выводить повторяющиеся строки

func main() {
	// Get flags
	column := flag.Int("k", 0, "column index for sorting")
	numeric := flag.Bool("n", false, "sort by numeric value")
	reverse := flag.Bool("r", false, "sort in reverse order")
	unique := flag.Bool("u", false, "do not print duplicate lines")

	flag.Parse()

	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Slice for our lines
	lines := []string{}

	// Append lines to slice
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, strings.TrimSpace(line))
	}

	// Sort function
	sortFunc := func(i, j int) bool {
		a := lines[i]
		b := lines[j]

		// Columns
		fieldsA := strings.Fields(a)
		fieldsB := strings.Fields(b)

		// if flag k required
		if *column > 0 {
			// Check the column
			if *column > len(fieldsA) || *column > len(fieldsB) {
				log.Fatalf("Column %d is not available in some lines", *column)
			}

			a = fieldsA[*column-1]
			b = fieldsB[*column-1]
		}

		// sort by digit
		if *numeric {
			aFloat, err := strconv.ParseFloat(a, 64)
			if err != nil {
				log.Fatalf("Unable to parse digit value in line: %s", lines[i])
			}
			bFloat, err := strconv.ParseFloat(b, 64)
			if err != nil {
				log.Fatalf("Unable to parse digit value in line: %s", lines[j])
			}

			return aFloat < bFloat
		}

		return a < b
	}

	sort.Slice(lines, sortFunc)

	if *unique {
		uniqueLines := []string{}
		uniqueSet := make(map[string]bool)

		for _, line := range lines {
			if !uniqueSet[line] {
				uniqueLines = append(uniqueLines, line)
				uniqueSet[line] = true
			}
		}

		lines = uniqueLines
	}

	// if flag r required
	if *reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	// Output sorted
	for _, line := range lines {
		fmt.Println(line)
	}
}
