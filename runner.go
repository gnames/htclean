package htclean

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/gnames/htclean/model"
)

func (htc *HTclean) Run() error {
	chIn := make(chan [][]string)
	var wg sync.WaitGroup
	// wg.Add(htc.JobsNum)
	wg.Add(1)
	// for i := 0; i < htc.JobsNum; i++ {
	for i := 0; i < 1; i++ {
		go worker(chIn, &wg)
	}
	htc.collectTitles(chIn)
	wg.Wait()
	return nil
}

func worker(ch <-chan [][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for rows := range ch {
		d := model.NewDecider(rows)
		d.Decide()
		fmt.Println(d.Accept)
		fmt.Println("----------------------")
	}
}

func (htc *HTclean) collectTitles(ch chan<- [][]string) {
	var t string
	var ts [][]string
	f, err := os.Open(htc.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	for {
		l, err := r.Read()
		if err == io.EOF {
			if len(ts) > 0 {
				ch <- ts
			}
			break
		} else if err != nil {
			log.Fatal(err)
		}
		id := l[model.TitleF.Int()]

		if t == "" {
			t = id
		} else if t != id {
			ch <- ts
			t = id
			ts = nil
		}
		ts = append(ts, l)
	}
	close(ch)
}
