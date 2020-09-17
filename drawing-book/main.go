// https://www.hackerrank.com/challenges/drawing-book

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func min(a, b int32) int32 {
  if a < b {
    return a
  } else {
    return b
  }
}

func pageCount(n int32, p int32) int32 {

  // Declare the variables
  var i32MinimumPageTurn int32
  
  // Do the logic
  if n % 2 == 0 {
    i32MinimumPageTurn = min(p, n + 1 - p) / 2
  } else {
    i32MinimumPageTurn = min(p, n - p) / 2
  }
  
  return i32MinimumPageTurn
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  n := int32(nTemp)

  pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  p := int32(pTemp)

  result := pageCount(n, p)

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
