// https://www.hackerrank.com/challenges/designer-pdf-viewer

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

// Complete the designerPdfViewer function below.
func designerPdfViewer(h []int32, word string) int32 {

  // Declare and initialize the variables
  var i32CurrentHeight int32
  var i32MaximumHeight int32 = h[word[0] - 'a']

  // Do the loop
  for i := 1; i < len(word); i++ {
    i32CurrentHeight = h[word[i] - 'a']
    if i32CurrentHeight > i32MaximumHeight {
      i32MaximumHeight = i32CurrentHeight
    }
  }
  
  return i32MaximumHeight * int32(len(word))
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  hTemp := strings.Split(readLine(reader), " ")

  var h []int32

  for i := 0; i < 26; i++ {
    hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
    checkError(err)
    hItem := int32(hItemTemp)
    h = append(h, hItem)
  }

  word := readLine(reader)

  result := designerPdfViewer(h, word)

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
