// https://www.hackerrank.com/challenges/jumping-on-the-clouds

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
  
  // Declare and initialize the variables
  var i int = 0
  var i32HoopCounter int32 = 0
  
  // Do the loop
  for i < len(c) {
    if i + 2 < len(c) {
      if c[i + 2] == 0 {
        i += 2
        i32HoopCounter++
      } else if c[i + 2] == 1 {
        i += 1
        i32HoopCounter++
      }
    } else if i + 1 < len(c) {
      if c[i + 1] == 0 {
        i += 1
        i32HoopCounter++
      }
    } else if i < len(c) {
      break
    }
  }

  return i32HoopCounter
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

  cTemp := strings.Split(readLine(reader), " ")

  var c []int32

  for i := 0; i < int(n); i++ {
    cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
    checkError(err)
    cItem := int32(cItemTemp)
    c = append(c, cItem)
  }

  result := jumpingOnClouds(c)

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
