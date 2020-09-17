// https://www.hackerrank.com/challenges/counting-valleys

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
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
  
  // Declare and initialize the variables
  var i32CurrentPosition int32 = 0
  var i32PreviousPosition int32
  var i32ValleyCounter int32 = 0
  
  // Do the loop
  for i := 0; i < len(path); i++ {
    i32PreviousPosition = i32CurrentPosition
  
    if path[i] == 'D' {
      i32CurrentPosition--
    } else if path[i] == 'U' {
      i32CurrentPosition++
    }
    
    // Valley walked through if the previous position is in a valley
    // and the current position is no longer in a valley
    if (i32PreviousPosition < 0) && (i32CurrentPosition >= 0) {
      i32ValleyCounter++
    }
  }
  
  return i32ValleyCounter
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  steps := int32(stepsTemp)

  path := readLine(reader)

  result := countingValleys(steps, path)

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
