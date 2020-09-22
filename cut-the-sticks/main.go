package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) []int32 {

  // Declare and initialize the variables
  var ai32RemainingSticksCounter []int32
  var ai32CurrentSticks []int32 = arr
  
  // Do the loop
  for len(ai32CurrentSticks) > 0 {
    ai32RemainingSticksCounter = append(ai32RemainingSticksCounter, int32(len(ai32CurrentSticks)))
    
    // Find the minimum positive element in the array
    var i32MinimumElement int32 = ai32CurrentSticks[0]
    for i := 1; i < len(ai32CurrentSticks); i++ {
      if ai32CurrentSticks[i] < i32MinimumElement {
        i32MinimumElement = ai32CurrentSticks[i]
      }
    }
    
    // Reduce sticks
    var ai32TemporarySticks []int32
    var i32ReducedStickLength int32
    for i := 0; i < len(ai32CurrentSticks); i++ {
      i32ReducedStickLength = ai32CurrentSticks[i] - i32MinimumElement
      if i32ReducedStickLength > 0 {
        ai32TemporarySticks = append(ai32TemporarySticks, i32ReducedStickLength)
      }
    }
    ai32CurrentSticks = ai32TemporarySticks
  }
  
  return ai32RemainingSticksCounter
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

  arrTemp := strings.Split(readLine(reader), " ")

  var arr []int32

  for i := 0; i < int(n); i++ {
    arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
    checkError(err)
    arrItem := int32(arrItemTemp)
    arr = append(arr, arrItem)
  }

  result := cutTheSticks(arr)

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
