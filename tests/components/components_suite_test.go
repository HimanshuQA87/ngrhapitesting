package components

import (
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	file *os.File
)

func TestComponents(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Components Suite")
}

var _ = BeforeSuite(func() {

	file, _ := os.OpenFile("Components.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//err := os.Rename("/apitesting/tests/components/Components.txt", "../bd-connect-api/apitesting/logs/Components.txt")
	//if err != nil {
	//log.Fatal(err)
	//}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

})

var _ = AfterSuite(func() {
	file.Close()

})
