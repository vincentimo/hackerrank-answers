// https://www.hackerrank.com/challenges/strange-advertising

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the viralAdvertising function below.
func viralAdvertising(n int32) int32 {

  // Declare and initialize the variables
  var i32CurrentShared int32 = 5
  var i32CurrentLiked int32
  var i32CumulativeLiked int32 = 0
  
  for i := 1; i <= int(n); i++ {
    if i > 1 {
      i32CurrentShared = i32CurrentLiked * 3
    }
    i32CurrentLiked = i32CurrentShared / 2
    i32CumulativeLiked += i32CurrentLiked
  }
  
  return i32CumulativeLiked
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

  result := viralAdvertising(n)

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
