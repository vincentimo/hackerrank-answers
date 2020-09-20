// https://www.hackerrank.com/challenges/extra-long-factorials

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
  "math/big"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {

  // Declare and initialize the variables
  var bigIntMultiplicationResult = new(big.Int)
  
  // Do the thing
  bigIntMultiplicationResult.MulRange(1, int64(n))

  // Print the result
  fmt.Printf("%d\n", bigIntMultiplicationResult)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  n := int32(nTemp)

  extraLongFactorials(n)
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
