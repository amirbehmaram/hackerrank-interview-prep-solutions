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
    
    var uniqueValleys int32 = 0
    
    /*
    * Elevation is how we will track whether we are above or below
    * sea level. Will help us differentiate between unique valleys.
    */
    elevation := 0
    
    // Iterate over each path item and get the specific character
    // so we can easily keep track of whether we enter a new valley or not.
    for _ , character := range path {
    
        if character == 'U' {
           elevation++
            
        } else if character == 'D' {
           elevation--
           
           /**
           * If we dip back below sea level we know we
           * have entered a new valley
           */
           if elevation == -1 {
               uniqueValleys++
           } 
        }
    }
    
    return uniqueValleys
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
