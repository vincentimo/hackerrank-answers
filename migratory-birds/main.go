// https://www.hackerrank.com/challenges/migratory-birds

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {

  // Declare and initialize the variables
  aiBirdCounter := [6]int{0, 0, 0, 0, 0, 0}
  
  // Calculate the count for each bird
  for i := 0; i < len(arr); i++ {
    aiBirdCounter[int(arr[i])]++
  }
  
  // Declare and initialize the variables
  var iIndexWithMaximumValue int = 1
  var iMaximumValue int = aiBirdCounter[iIndexWithMaximumValue]
  
  // Find the bird with the maximum count
  for i := 2; i < len(aiBirdCounter); i++ {
    if aiBirdCounter[i] > iMaximumValue {
      iIndexWithMaximumValue, iMaximumValue = i, aiBirdCounter[i]  
    }
  }
  
  return int32(iIndexWithMaximumValue)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)

  arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var arr []int32

  for i := 0; i < int(arrCount); i++ {
    arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
    checkError(err)
    arrItem := int32(arrItemTemp)
    arr = append(arr, arrItem)
  }

  result := migratoryBirds(arr)

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
