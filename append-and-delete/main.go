// https://www.hackerrank.com/challenges/append-and-delete

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Helper function to calculate the mininum of 2 integers
func min(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
  
  // Declare and initialize the variables
  var sIsOperationPossible string = "No"
  
  // Do the logic
  if int(k) >= len(s) + len(t) {
    sIsOperationPossible = "Yes"
  } else {
  
    // Declare and initialize the variables
    var iMinimumLength int = min(len(s), len(t))
    var iCommonCharacterCounter int = 0
    var bLoopCondition bool = true
    
    // Do the loop
    for bLoopCondition {
      bLoopCondition = false
      if iCommonCharacterCounter < iMinimumLength {
        if s[iCommonCharacterCounter] == t[iCommonCharacterCounter] {
          iCommonCharacterCounter++
          bLoopCondition = true
        }
      }
    }
    
    // Effective length difference is the length difference
    // between s and t after the common characters are removed
    var iEffectiveLengthDifference int = len(s) + len(t) - 2 * iCommonCharacterCounter
    
    // Do the logic
    if (int(k) >= iEffectiveLengthDifference) && (int(k) % 2 == iEffectiveLengthDifference % 2) {
      sIsOperationPossible = "Yes"
    }
  }
  
  return sIsOperationPossible
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  s := readLine(reader)

  t := readLine(reader)

  kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  k := int32(kTemp)

  result := appendAndDelete(s, t, k)

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
