package example

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/trungtvq/craw/client/craw"
	"github.com/trungtvq/go-utils/stringutils"
)

func crawDa() {
	fmt.Println(craw.Post())
	//var arr [][]string
	var result []string
	var wg sync.WaitGroup
	c := make(chan []string, 10)
	for i := 0; i < 12; i++ {
		go func() {
			for j := 0; j < 50; j++ {
				c <- strings.Split(craw.Post(), "\n")
			}
			wg.Done()
		}()
		wg.Add(1)
	}

	go func() {
		for {
			a := <-c
			for l := 0; l < len(a); l++ {
				result = append(result, a[l])
			}
		}
	}()
	wg.Wait()
	fmt.Println(len(result))
	close(c)
	res := stringutils.DeleteReduction(result)
	fmt.Println(len(res))

	var str string
	for i := 0; i < len(res); i++ {
		str = str + result[i] + "\n"
	}

	f, err := os.Create("./ok")
	fmt.Println(err)

	defer f.Close()
	d2 := []byte(str)
	f.Write(d2)

}
