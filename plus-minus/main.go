package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
  
  // Declare and initialize the variables
  var iPositiveCount int = 0
  var iNegativeCount int = 0
  var iZeroCount int = 0
  var iArrayLength int = len(arr)
  
  // Do the loop
  for i := 0; i < iArrayLength; i++ {
    if arr[i] > 0 {
      iPositiveCount++
    } else if arr[i] < 0 {
      iNegativeCount++
    } else {
      iZeroCount++
    }
  }
  
  // Evaluate the ratio
  var f64PositiveRatio float64 = float64(iPositiveCount) / float64(iArrayLength)
  var f64NegativeRatio float64 = float64(iNegativeCount) / float64(iArrayLength)
  var f64ZeroRatio float64 = float64(iZeroCount) / float64(iArrayLength)
  
  // Print the ratio
  fmt.Printf("%f\n", f64PositiveRatio)
  fmt.Printf("%f\n", f64NegativeRatio)
  fmt.Printf("%f\n", f64ZeroRatio)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

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

  plusMinus(arr)
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
