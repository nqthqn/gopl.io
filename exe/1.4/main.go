// Modifed so it printe filenames out too
// Test it out: go build main.go && ./main foo bar baz
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Build it
	counts := make(map[string]Dup)
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			dup, present := counts[input.Text()]
			if !present {
				dup = Dup{
					names: make(map[string]struct{}),
				}
			}
			dup.count++
			dup.names[f.Name()] = struct{}{}
			counts[input.Text()] = dup
		}
		f.Close()
	}

	for line, dup := range counts {
		if dup.count > 1 {
			fmt.Println(dup.String(), line)
		}
	}
}

// Dup ...
type Dup struct {
	names map[string]struct{} // Set of file names the line occures in
	count int                 // Number of times the line was duplicated
}

func (d *Dup) String() string {
	var names []string
	for name := range d.names {
		names = append(names, name)
	}
	return fmt.Sprintf("%d %v", d.count, names)
}
