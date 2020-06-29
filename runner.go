package htclean

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/gnames/htclean/model"
)

func (htc *HTclean) Run() error {
	chIn := make(chan [][]string)
	chOut := make(chan *model.Decider)
	var wg sync.WaitGroup
	var wgOut sync.WaitGroup

	wgOut.Add(1)
	go htc.writer(chOut, &wgOut)

	wg.Add(htc.JobsNum)
	for i := 0; i < htc.JobsNum; i++ {
		go worker(chIn, chOut, &wg)
	}

	htc.collectTitles(chIn)
	wg.Wait()
	close(chOut)
	wgOut.Wait()
	return nil
}

func (htc *HTclean) writer(ch <-chan *model.Decider, wg *sync.WaitGroup) {
	defer wg.Done()
	namePath := filepath.Join(htc.OutputPath, "names.csv")
	junkPath := filepath.Join(htc.OutputPath, "junk.csv")
	wgood, err := os.Create(namePath)
	if err != nil {
		log.Fatal(err)
	}
	wbad, err := os.Create(junkPath)
	if err != nil {
		log.Fatal(err)
	}
	g := csv.NewWriter(wgood)
	b := csv.NewWriter(wbad)
	for d := range ch {
		if d.Accept {
			for _, v := range d.Rows {
				g.Write(v)
			}
		} else {
			for _, v := range d.Rows {
				b.Write(v)
			}
		}
	}
	g.Flush()
	b.Flush()
	wgood.Close()
	wbad.Close()
}

func worker(ch <-chan [][]string, chOut chan<- *model.Decider,
	wg *sync.WaitGroup) {
	defer wg.Done()
	for rows := range ch {
		d := model.NewDecider(rows)
		d.Decide()
		d.Title = nil
		chOut <- d
	}
}

func (htc *HTclean) collectTitles(ch chan<- [][]string) {
	var t string
	var ts [][]string
	f, err := os.Open(filepath.Join(htc.WorkPath, htc.InputFile))
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
			log.Println(err)
			continue
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
