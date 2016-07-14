package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "bytes"
)

func main() {
  in := bufio.NewScanner(os.Stdin)

  in.Scan()
  testCases, _ := strconv.Atoi(in.Text())
  for i := 0; i < testCases; i++ {
    in.Scan()
    word := in.Text()

    var odds, evens bytes.Buffer
    for j := 0; j < len(word); j++ {
      if(j % 2 == 0) {
        // agg evens
        evens.WriteString(string(word[j]))
      } else {
        // agg odds
        odds.WriteString(string(word[j]))
      }
    }

    fmt.Printf("%s %s\n", evens.String(), odds.String())
  }
}
