// https://www.hackerrank.com/challenges/magic-square-forming

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Helper function to calculate the absolute value of an int32
func abs(n int32) int32 {
  if n >= 0 {
    return n
  } else {
    return -n
  }
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {

  // Declare and initialize the variables
  a3i32MagicSquares := [8][3][3]int32{
    {{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
    {{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
    {{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
    {{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
    {{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
    {{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
    {{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
    {{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
  }
  var i32MinimumCost, i32CurrentCost int32
  
  // Do the loop
  for i := 0; i < len(a3i32MagicSquares); i++ {
    i32CurrentCost = 0
    
    for j := 0; j < len(a3i32MagicSquares[0]); j++ {
      for k := 0; k < len(a3i32MagicSquares[0][0]); k++ {
        i32CurrentCost += abs(a3i32MagicSquares[i][j][k] - s[j][k])
      }
    }
    
    if i == 0 {
      i32MinimumCost = i32CurrentCost
    }
    
    if i32CurrentCost < i32MinimumCost {
      i32MinimumCost = i32CurrentCost
    }
  }

  return i32MinimumCost
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  var s [][]int32
  for i := 0; i < 3; i++ {
    sRowTemp := strings.Split(readLine(reader), " ")

    var sRow []int32
    for _, sRowItem := range sRowTemp {
      sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
      checkError(err)
      sItem := int32(sItemTemp)
      sRow = append(sRow, sItem)
    }

    if len(sRow) != 3 {
      panic("Bad input")
    }

    s = append(s, sRow)
  }

  result := formingMagicSquare(s)

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
