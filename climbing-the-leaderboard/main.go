// https://www.hackerrank.com/challenges/climbing-the-leaderboard

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
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {

  // Declare the variables
  var ai32DistinctRank []int32
  
  // Initialize the variables
  ai32DistinctRank = append(ai32DistinctRank, ranked[0])
  
  // Build the distinct rank
  for i := 1; i < len(ranked); i++ {
    if ranked[i] < ranked[i - 1] {
      ai32DistinctRank = append(ai32DistinctRank, ranked[i])
    }
  }
  
  // Declare and initialize the variables
  var ai32PlayerRank []int32
  var j int32 = int32(len(ai32DistinctRank)) - 1
  var i32CurrentRank int32
  
  // Calculate the player's rank
  for i := 0; i < len(player); i++ {
    for player[i] > ai32DistinctRank[j] {
      if j > 0 {
        j--
      } else { // j == 0
        break
      }
    }
    if player[i] < ai32DistinctRank[j] {
      i32CurrentRank = j + 2
    } else { // player[i] >= ai32DistinctRank[j]
      i32CurrentRank = j + 1
    }
    ai32PlayerRank = append(ai32PlayerRank, i32CurrentRank)
  }
  
  return ai32PlayerRank
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)

  rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var ranked []int32

  for i := 0; i < int(rankedCount); i++ {
    rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
    checkError(err)
    rankedItem := int32(rankedItemTemp)
    ranked = append(ranked, rankedItem)
  }

  playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)

  playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var player []int32

  for i := 0; i < int(playerCount); i++ {
    playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
    checkError(err)
    playerItem := int32(playerItemTemp)
    player = append(player, playerItem)
  }

  result := climbingLeaderboard(ranked, player)

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
