// https://www.hackerrank.com/challenges/day-of-the-programmer

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Check Julian leap year
func isJulianLeap(year int32) bool {
  return year % 4 == 0
}

// Check Gregorian leap year
func isGregorianLeap(year int32) bool {
  return (year % 400 == 0) || ((year % 4 == 0) && (year % 100 != 0))
}

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
  
  // Declare the variable
  var sDdMm string
  
  // Check based on the type of calendar
  if year < 1918 {
    if isJulianLeap(year) {
      sDdMm = "12.09."
    } else {
      sDdMm = "13.09."
    }
  } else if year == 1918 {
    sDdMm = "26.09."
  } else {
    if isGregorianLeap(year) {
      sDdMm = "12.09."
    } else {
      sDdMm = "13.09."
    }
  }
  
  return sDdMm + strconv.FormatInt(int64(year), 10)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  year := int32(yearTemp)

  result := dayOfProgrammer(year)

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
