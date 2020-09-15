// https://www.hackerrank.com/challenges/divisible-sum-pairs

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
  
  // Declare and initialize the variables
  var i32Counter int32 = 0
  
  // Do the loop
  for i := 0; i < len(ar) - 1; i++ {
    for j := i + 1; j < len(ar); j++ {
      if (ar[i] + ar[j]) % k == 0 {
        i32Counter++
      }
    }
  }
  
  return i32Counter
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  nk := strings.Split(readLine(reader), " ")

  nTemp, err := strconv.ParseInt(nk[0], 10, 64)
  checkError(err)
  n := int32(nTemp)

  kTemp, err := strconv.ParseInt(nk[1], 10, 64)
  checkError(err)
  k := int32(kTemp)

  arTemp := strings.Split(readLine(reader), " ")

  var ar []int32

  for i := 0; i < int(n); i++ {
    arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
    checkError(err)
    arItem := int32(arItemTemp)
    ar = append(ar, arItem)
  }

  result := divisibleSumPairs(n, k, ar)

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
