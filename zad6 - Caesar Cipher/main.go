package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	output := ""
	for _, c := range s {
		letter := c
		changed := true
		var offset int32
		if c >= 'A' && c <= 'Z' {
			offset = 'A'
		} else if c >= 'a' && c <= 'z' {
			offset = 'a'
		} else {
			changed = false
		}
		if changed {
			letter = ((int32(c) - offset) + k) % ('Z' - 'A')
			output += string(offset + letter)
		} else {
			output += string(c)
		}
	}
	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	_ = nTemp

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
