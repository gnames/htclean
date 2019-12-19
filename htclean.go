package htclean

import (
	"fmt"
	"os"
	"runtime"
)

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
