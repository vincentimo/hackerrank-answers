// https://www.hackerrank.com/challenges/picking-numbers

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
  
  // Declare and initialize variables
  mi32i32ElementTypeCounter := make(map[int32]int32)
  
  // Build the map using iteration
  for i := 0; i < len(a); i++ {
    _, bExist := mi32i32ElementTypeCounter[a[i]]
    if !bExist {
      mi32i32ElementTypeCounter[a[i]] = 1
    } else {
      mi32i32ElementTypeCounter[a[i]]++
    }
  }
  
  // Declare and initialize variables
  var i32CurrentCount, i32MaximumCount int32
  
  // Find the longest subarray where the absolute difference
  // between any two elements is less than or equal to 1.
  // This means it is possible that the absolute difference
  // between any two elements is 0, i.e. the same element.
  // As such, if there are no consecutive element, then
  // count its own frequency.
  for i32Key, i32Value := range mi32i32ElementTypeCounter {
    i32NextValue, bExist := mi32i32ElementTypeCounter[i32Key + 1]
    if bExist {
      i32CurrentCount = i32Value + i32NextValue
    } else {
      i32CurrentCount = i32Value
    }
    
    if i32CurrentCount > i32MaximumCount {
      i32MaximumCount = i32CurrentCount
    }
  }
  
  return i32MaximumCount
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  n := int32(nTemp)

  aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var a []int32

  for i := 0; i < int(n); i++ {
    aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
    checkError(err)
    aItem := int32(aItemTemp)
    a = append(a, aItem)
  }

  result := pickingNumbers(a)

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
