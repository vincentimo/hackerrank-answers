// https://www.hackerrank.com/challenges/the-hurdle-race

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"

)
// Complete the hurdleRace function below.
func hurdleRace(k int32, height []int32) int32 {

  // Declare and initialize the variables
  var i32MaximumHeight int32 = height[0]
  var i32MagicPotionNeeded int32
  
  // Find the maximum height
  for i := 1; i < len(height); i++ {
    if height[i] > i32MaximumHeight {
      i32MaximumHeight = height[i]
    }
  }
  
  i32MagicPotionNeeded = i32MaximumHeight - k
  
  if i32MagicPotionNeeded < 0 {
    i32MagicPotionNeeded = 0
  }
  
  return i32MagicPotionNeeded
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

  heightTemp := strings.Split(readLine(reader), " ")

  var height []int32

  for i := 0; i < int(n); i++ {
    heightItemTemp, err := strconv.ParseInt(heightTemp[i], 10, 64)
    checkError(err)
    heightItem := int32(heightItemTemp)
    height = append(height, heightItem)
  }

  result := hurdleRace(k, height)

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
