package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func getDescription() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/description")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func getJobList() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/jobs")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

type Cache struct {
	DescString string
	JobList    string
	IsFilled   bool
}

func (c *Cache) aggregate() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	if !c.IsFilled {
		go func() {
			c.DescString = getDescription()
			wg.Done()
		}()

		go func() {
			c.JobList = getJobList()
			wg.Done()
		}()
		wg.Wait()

		c.IsFilled = true
	}
}

func (c Cache) testFetch() {
	start1 := time.Now()
	fmt.Println("dari calculate =", start1)

	c.aggregate()

	fmt.Println("took ", time.Since(start1))

	start2 := time.Now()
	fmt.Println("dari calculate =", start2)
	c.aggregate()
	fmt.Println("took ", time.Since(start2))
}

func main() {
	var testNum = 2
	c := Cache{}

	for i := 0; i < testNum; i++ {
		fmt.Printf("Test %d:\n", i+1)
		c.testFetch()
		fmt.Println()
	}
}
