package utils_test

import (
	"bosh-musings/utils"

	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PID", func() {
	var (
		name string
	)

	BeforeEach(func() {
		file, err := ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())

		name = file.Name()

		err = file.Close()
		Expect(err).NotTo(HaveOccurred())

		err = os.Remove(name)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		_ = os.Remove(name)
	})

	It("returns an error if given an empty name to write", func() {
		err := utils.WritePIDFile("")
		Expect(err.Error()).To(Equal("Can't write to empty file name"))
	})

	It("returns an error if it can't create a file with the given name", func() {
		err := utils.WritePIDFile("/not-there/pid-file")
		Expect(err.Error()).To(MatchRegexp(".*no such file or directory"))
	})

	It("writes the process id to a file", func() {
		err := utils.WritePIDFile(name)
		Expect(err).NotTo(HaveOccurred())

		_, err = os.Stat(name)
		Expect(err).NotTo(HaveOccurred())

		file, err := os.Open(name)
		Expect(err).NotTo(HaveOccurred())

		buf, err := ioutil.ReadAll(file)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(buf)).To(MatchRegexp(`[\d]+`))
	})
})
