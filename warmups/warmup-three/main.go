package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func repeatedString(s string, n int64) int64 {

    var originalStrLength int64 = int64(len(s))
    var count int64 = 0
    
    // Single Letter Edge Cases
    if (originalStrLength == 1) {
        if (s[0] == 'a') {
            return n
        } else {
            return 0
        }
    }

    // Get num times n go into OGStrLen
    quotient := n / originalStrLength
    
    // Get remainer of n going in to OGStrLen
    remainder := n % originalStrLength
    
    // Get number of times 'a' exists in s
    var numAlphas int64 = 0
    
    for _, val := range(s) {
        if (val == 'a') {
            numAlphas++
        }
    }
    
    // Do the math portion //
    
    // Get the base number excluding our remainder
    count = numAlphas * quotient
    
    // If we don't have a remainder bail
    if (remainder == 0) {
        return count
    }
    
    // Loop through s until we hit our remainer num and tack on any
    // extra a's
    for index, val := range(s) {
        
        if (val == 'a') {
            count++
        }
        
        index++
        
        if (int64(index) == remainder) {
            return count
        }
    }
    
    return count
}

// Brute Force: O(n) way ^
// Will time out on HackerRank

// index := 0
// for ( n > 0 ) {
        
//     if (s[index] == 'a') {
//         count++
//     } 
    
//     n--
    
//     if (index == originalStrLength - 1) {
//         index = 0
//     } else {
//         index++
//     } 
// }

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

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
