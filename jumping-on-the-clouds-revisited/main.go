// https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited

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
func jumpingOnClouds(c []int32, k int32) int32 {

  // Declare and initialize the variables
  var n int = len(c)
  var iCurrentPosition int = 0
  var i32RemainingEnergy int32 = 100
  var bLoopCondition bool = true
  
  // Do the loop
  for bLoopCondition {
    iCurrentPosition = (iCurrentPosition + int(k)) % n
    i32RemainingEnergy -= 2 * c[iCurrentPosition] + 1
    if iCurrentPosition == 0 {
      bLoopCondition = false
    }
  }
  
  return i32RemainingEnergy
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

  cTemp := strings.Split(readLine(reader), " ")

  var c []int32

  for i := 0; i < int(n); i++ {
    cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
    checkError(err)
    cItem := int32(cItemTemp)
    c = append(c, cItem)
  }

  result := jumpingOnClouds(c, k)

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
