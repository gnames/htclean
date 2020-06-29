package htclean

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type HTclean struct {
	WorkPath     string
	InputFile    string
	KeyValPath   string
	OutputPath   string
	JobsNum      int
	ProgressNum  int
	LangFile     string
	LangTitleIdx int
	LangPageIdx  int
	LangIdx      int
}

type Option func(h *HTclean)

// OptWorkPath is an absolute path to a directory that contains all existing and
// future input and output for htclean.
func OptWorkPath(s string) Option {
	return func(h *HTclean) {
		h.WorkPath = s
	}
}

// OptIntputFile is a name of a file located in WorkPath that contains
// name-finding data to use for htclean. Its has the following fields:
// timeStamp, titleID, pageNum, nameVerbatim, name, odds, nameType
func OptInputFile(s string) Option {
	return func(h *HTclean) {
		h.InputFile = s
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

// OptLangFile is a name of a file with the language data. This file should
// be a valid csv file and should contain TitleID, PageID and Language values.
func OptLangFile(s string) Option {
	return func(h *HTclean) {
		h.LangFile = s
	}
}

// OptLangTitleIdx is a field index in LangFile that contains TitleID values.
func OptLangTitleIdx(i int) Option {
	return func(h *HTclean) {
		h.LangTitleIdx = i
	}
}

// OptLangPageIdx is a field index in LangFile that contains PageID values.
func OptLangPageIdx(i int) Option {
	return func(h *HTclean) {
		h.LangPageIdx = i
	}
}

// OptLangIdx is a field index in LangFile that contains language values.
func OptLangIdx(i int) Option {
	return func(h *HTclean) {
		h.LangIdx = i
	}
}

func NewHTclean(opts ...Option) (*HTclean, error) {
	htc := &HTclean{
		OutputPath: "output",
		KeyValPath: "kv",
		JobsNum:    runtime.NumCPU(),
	}
	for _, opt := range opts {
		opt(htc)
	}
	err := htc.setOutputDir()
	return htc, err
}

func (htc *HTclean) setOutputDir() error {
	output := filepath.Join(htc.WorkPath, htc.OutputPath)
	path, err := os.Stat(output)
	if os.IsNotExist(err) {
		return os.MkdirAll(output, 0755)
	}
	if path.Mode().IsRegular() {
		return fmt.Errorf("'%s' is a file, not a directory", htc.OutputPath)
	}
	return nil
}
