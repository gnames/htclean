package htclean_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	testOutput = "/tmp/htclean-test"
)

func TestHtclean(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Htclean Suite")
}
