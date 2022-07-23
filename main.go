package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"happy", "good", "awesome"},
		{"sad", "bad", "terrible"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods[0]))

	var mi int
	if mood != "positive" {
		mi = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[mi][n])

	// books := [...]string{

	// 	"Kafkas Revenge",
	// 	"Stay Golden",
	// 	"Everythingship",
	// 	"Kafkas Revenge 2nd Edition",
	// }
	// fmt.Printf("books : %#v\n", books)

	// var (
	// 	wBooks [winter]string
	// 	sBooks [summer]string
	// )

	// wbooks[0]

	// for i := range sBooks {
	// 	sBooks[i] = books[i+1]
	// }

	// fmt.Printf("\nwinter : %#v\n", wBooks)
	// fmt.Printf("\nsummer : %#v\n", sBooks)

	// var published [len(books)]bool

	// published[0] = true
	// published[len(books)-1] = true
	// fmt.Println("\nPublished Books")
	// for i, ok := range published {
	// 	if ok {
	// 		fmt.Printf("+%s\n", books[i])
	// 	}
	// }
}
