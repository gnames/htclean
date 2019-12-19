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
			hti, _ := NewHTclean(initOpts()...)
			Expect(hti.JobsNum).To(Equal(4))
			Expect(hti.OutputPath).To(Equal(testOutput))
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
