// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Helper function to calculate the absolute value of an integer
func abs(n int32) int32 {
  if n >= 0 {
    return n
  } else {
    return -n
  }
}

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {

  // Declare and initialize the variables
  var sCurrentNumber, sCurrentReversedNumber string
  var i32CurrentReversedNumber int32
  var i32BeautifulCounter int32 = 0
  
  // Do the loop
  for p := i; p <= j; p++ {
    sCurrentNumber = strconv.FormatInt(int64(p), 10)
    
    sCurrentReversedNumber = ""
    for q := len(sCurrentNumber) - 1; q >= 0; q-- {
      sCurrentReversedNumber += string(sCurrentNumber[q])
    }
    
    i64CurrentReversedNumber, err := strconv.ParseInt(sCurrentReversedNumber, 10, 64)
    checkError(err)
    i32CurrentReversedNumber = int32(i64CurrentReversedNumber)
    
    if abs(p - i32CurrentReversedNumber) % k == 0 {
      i32BeautifulCounter++
    }
  }
  
  return i32BeautifulCounter
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  ijk := strings.Split(readLine(reader), " ")

  iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
  checkError(err)
  i := int32(iTemp)

  jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
  checkError(err)
  j := int32(jTemp)

  kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
  checkError(err)
  k := int32(kTemp)

  result := beautifulDays(i, j, k)

  fmt.Fprintf(writer, "%d\n", result)

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
