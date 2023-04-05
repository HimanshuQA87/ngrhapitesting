package integration_test

import (
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var file *os.File

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	file, _ := os.OpenFile("Integration.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	err := os.Rename("/tests/integration/Integration.txt", "../logs")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

})

var _ = AfterSuite(func() {
	file.Close()
})
