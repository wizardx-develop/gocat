package concat

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func InToOut() {
	for {
		reader := bufio.NewReader(os.Stdin)
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprint(os.Stdout, result)
	}

}

func Concat(path string) [][]byte {
	var lines [][]byte

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		lines = append(lines, line)
	}

	return lines
}

func ShowEnds(lines [][]byte) [][]byte {
	var outlines [][]byte

	for _, line := range lines {
		outlines = append(outlines, bytes.ReplaceAll(line, []byte("\n"), []byte("$\n")))
	}

	return outlines
}

func ShowStrNumbers(lines [][]byte) [][]byte {
	var outlines [][]byte

	for i, line := range lines {
		s := "[" + strconv.Itoa(i+1) + "]:\t"
		line = append([]byte(s), line...)
		outlines = append(outlines, line)
	}

	return outlines
}

func ShowTabs(lines [][]byte) [][]byte {
	var outlines [][]byte

	for _, line := range lines {
		outlines = append(outlines, bytes.ReplaceAll(line, []byte("\t"), []byte("^I")))
	}

	return outlines
}

func ShowNonEmptyStrNumbers(lines [][]byte) [][]byte {
	var outlines [][]byte
	for i, line := range lines {
		if bytes.Equal(line, []byte("\n")) {
			continue
		} else {
			s := "[" + strconv.Itoa(i+1) + "]:\t"
			line = append([]byte(s), line...)
			outlines = append(outlines, line)
		}
	}
	return outlines
}

func ToString(lines [][]byte) string {
	var result string

	for _, line := range lines {
		result += string(line)
	}

	return result
}
