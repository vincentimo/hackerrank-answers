// https://www.hackerrank.com/challenges/between-two-sets

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
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

// Calculate GCD with Euclidean algorithm
func gcd(a, b int32) int32 {
  for a % b != 0 {
    a, b = b, (a % b)
  }
  return b
}

// Calculate GCD of an array
func gcdArray(ai32Array []int32) int32 {
  var i32Output int32 = ai32Array[0]
  for i := 1; i < len(ai32Array); i++ {
    i32Output = gcd(i32Output, ai32Array[i])
  }
  return i32Output
}

// Calculate LCM with Euclidean algorithm
func lcm(a, b int32) int32 {
  return a * b / gcd(a, b)
}

// Calculate LCM of an array
func lcmArray(ai32Array []int32) int32 {
  var i32Output int32 = ai32Array[0]
  for i := 1; i < len(ai32Array); i++ {
    i32Output = lcm(i32Output, ai32Array[i])
  }
  return i32Output
}

// Write your code here
func getTotalX(a []int32, b []int32) int32 {

  // Declare and initialize the variables
  var i32MinimumCandidate int32 = lcmArray(a)
  var i32MaximumCandidate int32 = gcdArray(b)
  var i32CurrentCandidate int32 = i32MinimumCandidate
  var i32Counter int32 = 0
  
  // The candidates are defined as the multiples of the minimum candidate
  // A candidate is valid if it is the factor of the maximum candidate
  for i32CurrentCandidate <= i32MaximumCandidate {
    if i32MaximumCandidate % i32CurrentCandidate == 0 {
      i32Counter++
    }
    i32CurrentCandidate += i32MinimumCandidate
  }
  
  return i32Counter
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
  checkError(err)
  n := int32(nTemp)

  mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
  checkError(err)
  m := int32(mTemp)

  arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var arr []int32

  for i := 0; i < int(n); i++ {
    arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
    checkError(err)
    arrItem := int32(arrItemTemp)
    arr = append(arr, arrItem)
  }

  brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var brr []int32

  for i := 0; i < int(m); i++ {
    brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
    checkError(err)
    brrItem := int32(brrItemTemp)
    brr = append(brr, brrItem)
  }

  total := getTotalX(arr, brr)

  fmt.Fprintf(writer, "%d\n", total)

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
