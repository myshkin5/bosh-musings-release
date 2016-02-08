package endtoend_tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEndToEndTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EndToEndTests Suite")
}
