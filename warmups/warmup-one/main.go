package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/****************************************************/
/**** Function I completed to solve the question ****/
/****************************************************/
func sockMerchant(n int32, ar []int32) int32 {

    var totalPairs int32 = 0
    var sockIndex int32
    
    // Create a Map to use to easily keep track of
    // sock counts.
    sockMap := make(map[int32]int32)
    
    for sockIndex = 0; sockIndex < n; sockIndex++ {
        // Get the current sock value
        currentSockValue := ar[sockIndex]
        
        // If the curent sock value is already in the map,
        // just increment that value, otherwise create a
        // new map index with a value of 1
        if mapValue, exists := sockMap[currentSockValue]; exists {            
            sockMap[currentSockValue] = mapValue + 1
        } else {
            sockMap[currentSockValue] = 1
        }
    }
    
    // Iterate over our map and count the number of pairs
    // in each index.
    // Note int math will round down decimals to the nearest
    // whole number. ex 21 / 2 = 10, or 1 / 2 = 0
    for key := range sockMap {
        totalPairs = totalPairs + (sockMap[key] / 2)
    }    
    
    return totalPairs
}
/****************************************************/
/****************************************************/
/****************************************************/

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := sockMerchant(n, ar)

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