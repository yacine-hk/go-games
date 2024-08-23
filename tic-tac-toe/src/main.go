
package main

/*
import (
	"fmt"
)

var inputs = [2]string{"X", "O"}
var board = [][]string{
    {"_", "_", "_"},
    {"_", "_", "_"},
    {"_", "_", "_"},
}

var wins = [][]Vertex{
    {{0,0}, {0,1}, {0,2}},
    {{1,0}, {1,1}, {1,2}},
    {{2,0}, {2,1}, {2,2}},
    {{0,0}, {1,0}, {2,0}},
    {{0,1}, {1,1}, {2,1}},
    {{0,2}, {1,2}, {2,2}},
    {{0,0}, {1,1}, {2,2}},
    {{2,0}, {1,1}, {0,2}},
}

var lines = []string{"-", "-", "-", "|", "|", "|", "\\", "/"}

type Vertex struct {
    x int
    y int
}

func compare(w, p []Vertex) bool {
    var scores = 0
    for _, i := range w {
        for _, j := range p {
            if i == j {
                scores++
                break
            }
        }
    }
    if scores == 3 {
        return true
    }
    return false
}

func winer(player []Vertex) (bool, int) {
    for i, k := range wins {
        if compare(k, player) == true {
            return true, i
        }
    }
    return false, -1
}

func displayBoard(board [][]string) {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Print(board[i][j] + " ")
        }
        fmt.Print("\n")
    }
}

func main1() {
    moves := make(map[Vertex]string)
    allmoves := make(map[string][]Vertex)
    var x, y int
    var ok bool
    var move Vertex

    displayBoard(board)
    i := 0
    for {
        if i == 9 {
            fmt.Println("DRAW!")
            return
        }
        fmt.Printf("%s move: ", inputs[i%2])
        fmt.Scan(&x, &y)

        if x > 2 || x < 0 || y > 2 || y < 0 {
            fmt.Println("Out of range")
            continue
        }

        move = Vertex{int(x), int(y)}
        if _, ok = moves[move]; ok == true {
            fmt.Println("move already played")
            continue
        }
        moves[move] = inputs[i%2]
        allmoves[inputs[i%2]] = append(allmoves[inputs[i%2]], move)
        
        board[x][y] = inputs[i%2]
        displayBoard(board)
        k, idx := winer(allmoves[inputs[i%2]]) 
        if k == true {
            fmt.Println(wins[idx][0].x)
            board[wins[idx][0].x][wins[idx][0].y] = lines[idx]
            board[wins[idx][1].x][wins[idx][1].y] = lines[idx]
            board[wins[idx][2].x][wins[idx][2].y] = lines[idx]
            fmt.Printf("%s Is winner \n", inputs[i%2])
            displayBoard(board)
            return
        }

        i++
    }
}
*/
