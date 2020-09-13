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
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
  
  // Declare and initialize the variables
  var sHour string = string(s[0:2])
  var sMinute string = string(s[3:5])
  var sSecond string = string(s[6:8])
  var sMarker string = string(s[8:10])
  
  // Do the conversion
  i64Hour, err := strconv.ParseInt(sHour, 10, 64)
  checkError(err)
  
  if (sMarker == "AM") && (i64Hour == 12) {
    sHour = "00"
  } else if (sMarker == "PM") && (i64Hour != 12) {
    sHour = strconv.FormatInt(i64Hour + 12, 10)
  }
  
  // Compose the output
  var sOutput string = sHour + ":" + sMinute + ":" + sSecond
  
  return sOutput
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer outputFile.Close()

  writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

  s := readLine(reader)

  result := timeConversion(s)

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
