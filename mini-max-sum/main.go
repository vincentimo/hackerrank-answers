package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
  
  // Declare and initialize the variables
  var i64Sum int64 = 0
  var iArrayLength int = len(arr)
  
  // Do the loop to obtain the overall sum
  for i := 0; i < iArrayLength; i++ {
    i64Sum += int64(arr[i])
  }
  
  // Declare the variables
  var i64CurrentSum, i64MinimumSum, i64MaximumSum int64
  
  // Do the loop to obtain the minimum and maximum sum
  for i := 0; i < iArrayLength; i++ {
    i64CurrentSum = i64Sum - int64(arr[i])
    
    // Initialize the minimum and maximum sum for i = 0
    if i == 0 {
      i64MinimumSum = i64CurrentSum
      i64MaximumSum = i64CurrentSum
    }
    
    // Adjust the minimum and maximum sum
    if i64CurrentSum < i64MinimumSum {
      i64MinimumSum = i64CurrentSum
    } else if i64CurrentSum > i64MaximumSum {
      i64MaximumSum = i64CurrentSum
    }
  }
    
  // Print the minimum and maximum sum
  fmt.Printf("%d %d", i64MinimumSum, i64MaximumSum)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  arrTemp := strings.Split(readLine(reader), " ")

  var arr []int32

  for i := 0; i < 5; i++ {
    arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
    checkError(err)
    arrItem := int32(arrItemTemp)
    arr = append(arr, arrItem)
  }

  miniMaxSum(arr)
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
