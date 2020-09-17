// https://www.hackerrank.com/challenges/sock-merchant

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {

  // Declare and initialize the variables 
  mi32i32SockTypeCounter := make(map[int32]int32)
  var i32SockPairCounter int32 = 0
  
  // Count the number of socks for each type
  // Also count the number of sock pairs
  for i := 0; i < len(ar); i++ {
    _, bExist := mi32i32SockTypeCounter[ar[i]]
    if !bExist {
      mi32i32SockTypeCounter[ar[i]] = 1
    } else {
      mi32i32SockTypeCounter[ar[i]]++
    }
    
    i32Value, _ := mi32i32SockTypeCounter[ar[i]]
    if i32Value == 2 {
      i32SockPairCounter++
      mi32i32SockTypeCounter[ar[i]] = 0
    }
  }
  
  return i32SockPairCounter
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

  arTemp := strings.Split(readLine(reader), " ")

  var ar []int32

  for i := 0; i < int(n); i++ {
    arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
    checkError(err)
    arItem := int32(arItemTemp)
    ar = append(ar, arItem)
  }

  result := sockMerchant(n, ar)

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
