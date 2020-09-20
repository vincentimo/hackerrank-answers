// https://www.hackerrank.com/challenges/find-digits

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the findDigits function below.
func findDigits(n int32) int32 {

  // Declare and initialize the variables
  var i32RemainingNumber int32 = n
  var i32DivisorCounter int32 = 0
  var i32CurrentDigit int32
  
  // Do the loop
  for i32RemainingNumber != 0 {
    
    i32CurrentDigit = i32RemainingNumber % 10
    
    if i32CurrentDigit != 0 {
      if n % i32CurrentDigit == 0 {
        i32DivisorCounter++
      }
    }
    
    i32RemainingNumber /= 10
  }
  
  return i32DivisorCounter
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  t := int32(tTemp)

  for tItr := 0; tItr < int(t); tItr++ {
    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    result := findDigits(n)

    fmt.Fprintf(writer, "%d\n", result)
  }

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
