// https://www.hackerrank.com/challenges/circular-array-rotation

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Helper function which can calculate modulo of negative number
func mod(a, b int32) int32 {
  return (a % b + b) % b
}

// Complete the circularArrayRotation function below.
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {

  // Declare and initialize the array-rotation
  var ai32QueryResult []int32
  var i32InputArrayLength int32 = int32(len(a))

  // Do the loop
  for i := 0; i < len(queries); i++ {
    ai32QueryResult = append(ai32QueryResult, a[mod(queries[i] - k, i32InputArrayLength)])
  }
  
  return ai32QueryResult
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  nkq := strings.Split(readLine(reader), " ")

  nTemp, err := strconv.ParseInt(nkq[0], 10, 64)
  checkError(err)
  n := int32(nTemp)

  kTemp, err := strconv.ParseInt(nkq[1], 10, 64)
  checkError(err)
  k := int32(kTemp)

  qTemp, err := strconv.ParseInt(nkq[2], 10, 64)
  checkError(err)
  q := int32(qTemp)

  aTemp := strings.Split(readLine(reader), " ")

  var a []int32

  for i := 0; i < int(n); i++ {
    aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
    checkError(err)
    aItem := int32(aItemTemp)
    a = append(a, aItem)
  }

  var queries []int32

  for i := 0; i < int(q); i++ {
    queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    queriesItem := int32(queriesItemTemp)
    queries = append(queries, queriesItem)
  }

  result := circularArrayRotation(a, k, queries)

  for i, resultItem := range result {
    fmt.Fprintf(writer, "%d", resultItem)

    if i != len(result) - 1 {
      fmt.Fprintf(writer, "\n")
    }
  }

  fmt.Fprintf(writer, "\n")

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
