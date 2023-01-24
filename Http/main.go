package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	log := logWriter{}

	io.Copy(os.Stdout, resp.Body)

	io.Copy(log, resp.Body)

	fmt.Println(log)

	sq := square{10}
	printArea(sq)

	tr := triangle{5, 5}
	printArea(tr)

}

type logWriter struct{}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
