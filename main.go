package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var VERSION string = "0.1.0"

func ProcessFile(file string, query func(string) int) int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		fmt.Println(file)
		return 0
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	c := 0
	for s.Scan() {
		c += query(s.Text())
	}
	return c
}

func ProcessFiles(paths []string, query func(string) int) int {
	ch := make(chan int)
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)
		go func(p string) {
			ch <- ProcessFile(p, query)
			wg.Done()
		}(path)
	}
	t := 0
	go func() {
		for c := range ch {
			t += c
		}
	}()
	wg.Wait()
	return t
}

func GetFiles(root string) []string {
	paths := []string{}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if info.Name()[0:1] == "." {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	return paths
}

func main() {
	// Set the number of processors to the number of CPUs.
	runtime.GOMAXPROCS(runtime.NumCPU())

	var matches = flag.Int("o", 0, "output the matches found, use -1 for all")
	var number = flag.Bool("c", false, "output the count of the number of matches found")
	var timetaken = flag.Bool("t", false, "output the time taken in seconds")
	var version = flag.Bool("version", false, "output the version information")

	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		return
	}

	if len(flag.Arg(1)) == 0 {
		fmt.Println("Usage: [options] filename|directory \"query\"")
		fmt.Println("  ss -c ./path \"a b\"")
		fmt.Println("  ss -c ./path \"a b NOT y z\"")
		fmt.Println("  ss -c ./path \"a OR b\"")
		return
	}

	path, _ := filepath.Abs(flag.Arg(0))

	start := time.Now()
	files := GetFiles(path)

	t := ProcessFiles(files, Query(flag.Arg(1), os.Stdout, *matches))

	if *timetaken {
		fmt.Println(time.Since(start))
	}
	if *number {
		fmt.Println(t)
	}
}
