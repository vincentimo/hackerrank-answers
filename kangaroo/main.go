// https://www.hackerrank.com/challenges/kangaroo

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the kangaroo function below.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

  // Declare and initialize the variable
  var sOutput string = "NO"
  
  // If the leading kangaroo is faster, then the other kangaroo won't be able to catch up.
  // Since x1 < x2, the only potentially valid condition is v1 > v2.
  if v1 > v2 {
  
    // Declare and initialize the variable
    var i32CurrentX1 int32 = x1
    var i32CurrentX2 int32 = x2
    
    // Do the loop
    for i32CurrentX1 < i32CurrentX2 {
      i32CurrentX1 += v1
      i32CurrentX2 += v2
    }
    
    // Evaluate the positions after the lagging kangaroo catches up the leading one.
    if i32CurrentX1 == i32CurrentX2 {
      sOutput = "YES"
    }
  }
  
  // Return the evaluation result
  return sOutput
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  x1V1X2V2 := strings.Split(readLine(reader), " ")

  x1Temp, err := strconv.ParseInt(x1V1X2V2[0], 10, 64)
  checkError(err)
  x1 := int32(x1Temp)

  v1Temp, err := strconv.ParseInt(x1V1X2V2[1], 10, 64)
  checkError(err)
  v1 := int32(v1Temp)

  x2Temp, err := strconv.ParseInt(x1V1X2V2[2], 10, 64)
  checkError(err)
  x2 := int32(x2Temp)

  v2Temp, err := strconv.ParseInt(x1V1X2V2[3], 10, 64)
  checkError(err)
  v2 := int32(v2Temp)

  result := kangaroo(x1, v1, x2, v2)

  fmt.Fprintf(writer, "%s\n", result)

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
