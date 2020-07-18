package main

//https://practice.geeksforgeeks.org/problems/subarray-with-given-sum

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Testcase holds a testcase
type Testcase struct {
	N   int
	S   int
	Arr []int
}

func main() {
	tcs, err := getInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, tc := range tcs {
		printSequence(tc)
	}
}

func getInput() ([]Testcase, error) {
	tc := make([]Testcase, 0)

	var t int
	_, err := fmt.Scanf("%d", &t)
	if err != nil {
		return nil, err
	}

	for i := 0; i < t; i++ {
		var n int
		var s int

		_, err := fmt.Scanf("%d %d", &n, &s)
		if err != nil {
			return nil, err
		}

		arr := make([]int, 0)
		ct := Testcase{N: n, S: s, Arr: arr}

		r := bufio.NewReader(os.Stdin)
		ar, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}

		parts := strings.Split(strings.Trim(ar, "\n"), " ")

		if len(parts) != n {
			return nil, errors.New("invalid numbr of array items provided")
		}

		for _, n := range parts {
			d, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			ct.Arr = append(ct.Arr, d)
		}

		tc = append(tc, ct)
	}
	return tc, nil
}

func printSequence(t Testcase) {
	var i int
	var j int

	var s int

	for i < t.N && j < t.N {

		if s == t.S {
			fmt.Printf("%d %d\n", i+1, j+1)
			return
		}

		s += t.Arr[j]

		if s < t.S {
			j++
			continue
		} else if s > t.S {
			s -= t.Arr[i]
			i++
			for j > i && s > t.S {
				s -= t.Arr[j]
				j--
			}
		}
	}
	fmt.Printf("-1\n")
}
