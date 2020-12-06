// https://www.codechef.com/problems/CHFINTRO
package main

import "fmt"

func main() {
    var n, r, rank int
    _, _ = fmt.Scanf("%d %d", &n, &r)

    for i := 0; i < n; i++ {
        _, _ = fmt.Scan(&rank)

        if rank < r {
            fmt.Println("Bad boi")
        } else {
            fmt.Println("Good boi")
        }
    }
}
