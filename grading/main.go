// https://www.hackerrank.com/challenges/grading

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
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {

  // Declare the variables
  var ai32RoundedGrades []int32
  var i32CurrentDifference, i32CurrentRoundedGrade int32
  
  // Do the loop
  for i := 0; i < len(grades); i++ {
    
    i32CurrentDifference = 5 - (grades[i] % 5)
    
    // Magical grading mechanism
    if i32CurrentDifference < 3 {
      i32CurrentRoundedGrade = grades[i] + i32CurrentDifference
    } else {
      i32CurrentRoundedGrade = grades[i]
    }
    
    // Failure threshold
    if i32CurrentRoundedGrade >= 40 {
      ai32RoundedGrades = append(ai32RoundedGrades, i32CurrentRoundedGrade)
    } else {
      ai32RoundedGrades = append(ai32RoundedGrades, grades[i])
    }
  }
  
  return ai32RoundedGrades
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)

  var grades []int32

  for i := 0; i < int(gradesCount); i++ {
    gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    gradesItem := int32(gradesItemTemp)
    grades = append(grades, gradesItem)
  }

  result := gradingStudents(grades)

  for i, resultItem := range result {
    fmt.Fprintf(writer, "%d", resultItem)

    if i != len(result) - 1 {
      fmt.Fprintf(writer, "\n")
    }
  }

  fmt.Fprintf(writer, "\n")

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
