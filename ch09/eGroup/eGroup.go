package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	g := new(errgroup.Group)

	for _, url := range os.Args[1:] {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Println(url, "is OK.")
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Everything went fine!")
}
