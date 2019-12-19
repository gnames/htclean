package htclean

import (
	"fmt"
	"os"
	"runtime"
)

type title struct {
	id    string
	pages []page
	kinds map[string]int
}

type page struct {
	pageNum int
	names   []name
}

type name struct {
	nameVerbatim string
	name         string
	offsetStart  int
	offsetEnd    int
	odds         int
	kind         string
}

type HTclean struct {
	InputFile   string
	OutputPath  string
	JobsNum     int
	ProgressNum int
}

type Option func(h *HTclean)

// OptIntput is an absolute path to input csv file. Its has the following
// fields: timeStamp, titleID, pageNum, nameVerbatim, name, odds, nameType
func OptInput(s string) Option {
	return func(h *HTclean) {
		h.InputFile = s
	}
}

// OptOutput is an absolute path to a directory where results will be written.
// If such directory does not exist already, it will be created during
// initialization of HTindex instance.
func OptOutput(s string) Option {
	return func(h *HTclean) {
		h.OutputPath = s
	}
}

// OptJobs sets number of jobs/workers to run duing execution.
func OptJobs(i int) Option {
	return func(h *HTclean) {
		h.JobsNum = i
	}
}

// OptProgressNum sets how often to printout a line about the progress. When it
// is set to 1 report line appears after processing every title, and if it is 10
// progress is shows after every 10th title.
func OptProgressNum(i int) Option {
	return func(h *HTclean) {
		h.ProgressNum = i
	}

}

func NewHTclean(opts ...Option) (*HTclean, error) {
	htc := &HTclean{JobsNum: runtime.NumCPU()}
	for _, opt := range opts {
		opt(htc)
	}
	err := htc.setOutputDir()
	return htc, err
}

func (htc *HTclean) setOutputDir() error {
	path, err := os.Stat(htc.OutputPath)
	if os.IsNotExist(err) {
		return os.MkdirAll(htc.OutputPath, 0755)
	}
	if path.Mode().IsRegular() {
		return fmt.Errorf("'%s' is a file, not a directory", htc.OutputPath)
	}
	return nil
}

func (htc *HTclean) Run() error {
	fmt.Println(htc)
	return nil
}

// func main() {
// 	var t *title
// 	var p page
// 	f, err := os.Open("filtered.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	r := csv.NewReader(f)
//
// 	for {
// 		l, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			log.Fatal(err)
// 		}
// 		id, pgNum, verbatim, name, offsetStart, offsetEnd, odds, kind := l[1],
// 			l[2], l[3], l[4], l[5], l[6], l[7], l[8]
//
// 		if t == nil {
// 			pgNum, err := strconv.Atoi(pgNum)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			p = page{pageNum: pgNum}
// 			t = &title{id: id}
// 		}
// 		_ = p
// 		_ = verbatim
// 		_ = name
// 		_ = offsetStart
// 		_ = offsetEnd
// 		_ = odds
// 		_ = kind
// 	}
// }
