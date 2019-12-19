package htclean_test

import (
	"path/filepath"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnames/htclean"
)

var _ = Describe("Htclean", func() {
	Describe("NewHTclean", func() {
		It("creates an instance of HTclean", func() {
			htc, _ := NewHTclean()
			Expect(htc.JobsNum).To(Equal(runtime.NumCPU()))
			Expect(htc.ProgressNum).To(Equal(0))
		})
		It("can read options", func() {
			htc, _ := NewHTclean(initOpts()...)
			Expect(htc.JobsNum).To(Equal(4))
			Expect(htc.OutputPath).To(Equal(testOutput))
		})
	})

	Describe("Run", func() {
		It("cleans results", func() {
			htc, _ := NewHTclean(initOpts()...)
			err := htc.Run()
			Expect(err).To(BeNil())
		})
	})
})

func initOpts() []Option {
	input, err := filepath.Abs("./testdata/test-50k.csv")
	Expect(err).To(BeNil())
	opts := []Option{
		OptJobs(4),
		OptOutput(testOutput),
		OptInput(input),
	}
	return opts
}
