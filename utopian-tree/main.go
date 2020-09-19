// https://www.hackerrank.com/challenges/utopian-tree

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the utopianTree function below.
func utopianTree(n int32) int32 {
  
  // Declare and initialize the variables
  var i32Height int32 = 1
  
  // Do the loop
  for i := 1; i <= int(n); i++ {
    if i % 2 == 1 {
      i32Height *= 2
    } else {
      i32Height++
    }
  }

  return i32Height
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  t := int32(tTemp)

  for tItr := 0; tItr < int(t); tItr++ {
    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    result := utopianTree(n)

    fmt.Fprintf(writer, "%d\n", result)
  }

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
