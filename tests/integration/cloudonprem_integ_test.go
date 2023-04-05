package integration

import (
	"io/ioutil"
	"log"

	"github.com/HimanshuQA87/ngrhapitesting/struct_models"
	"github.com/HimanshuQA87/ngrhapitesting/utilities"

	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var body interface{}
var downloadbody interface{}
var viperenvcloud string = utilities.ViperEnvVariable("baseCloudOnPrem")

var _ = Describe("Post Request Integration: Cloud to On-Prem", Label("Integration"), func() {

	BeforeEach(func() {
		log.Println("----Test Started-- ")
		log.Printf("Running test : %s\n", CurrentSpecReport().FullText())
	})

	AfterEach(func() {
		Fail := CurrentSpecReport().Failure
		if Fail.Message == "" {
			log.Println("Test Passed")
		} else {
			log.Println("Test Failed")
			log.Println(Fail.Message)

		}
		log.Println("----Test Ended----")
		log.Println("")
	})

	It("Can send a file to S3 bucket using load balancer API and receive the same file", Label("Integration"), func() {

		/* 		filename, _ := utilities.RandomNameandEmail()
		   		endpoint := viperenvcloud
		   		headers := utilities.Get_common_headers()
		   		encodedstring := utilities.ConvertFiletoBase64("Test.pdf")
		   		body := struct_models.CloudPrem{FileName: filename + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

		   		resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

		   		data, err := struct_models.UnmarshalCloudPrem(resp.Body())
		   		if err != nil {
		   			log.Println(err)
		   		} else {
		   			log.Println("Response Deserialized successfully")
		   		}
		   		Expect(resp.StatusCode()).To(Equal(200), "Status code not equal to 200")
		   		Expect(resp.Time().Seconds()).To(BeNumerically("<", 5), "Response Time is more than 5 seconds")
		   		Expect(filename+".pdf").To(Equal(data.FileName), "File name not as expected") */

		resp, data, _, _, headers := utilities.SendFileToBucket("baseCloudOnPrem", "Test.pdf", "deploymentidsqe")
		Expect(resp.StatusCode()).To(Equal(200), "Status code not equal to 200")
		//download
		/* 	downloadendpoint := viperenvcloud + "/download"
		downloadbody := struct_models.CloudPrem{FileName: filename + ".pdf", Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

		downloadresp := utilities.Fire_post_request_bodystruct(headers, downloadbody, downloadendpoint)

		downloaddata, downloaderr := struct_models.UnmarshalCloudPrem(downloadresp.Body())

		if downloaderr != nil {
			log.Println(downloaderr)
		} else {
			log.Println("Response Deserialized successfully")
		} */

		downloadresp, downloaddata, _ := utilities.DownloadFileOnPrem("baseCloudOnPrem", data.FileName, "deploymentidsqe", headers)

		Expect(downloadresp.StatusCode()).To(Equal(202), "Status code not equal to 200")
		Expect(downloaddata.FileName).To(Equal(data.FileName), "File name not as expected after download")
		Expect(downloaddata.Error).To(Equal(""), "Error is not empty")
		Expect(downloaddata.Body).ShouldNot(BeEmpty(), "Body is empty")

	})

	XIt("Can send a file to S3 bucket using load balancer API and receive Access Denied on passing incorrect file name", func() {

		filename, _ := utilities.RandomNameandEmail()
		endpoint := viperenvcloud
		headers := utilities.Get_common_headers()
		encodedstring := utilities.ConvertFiletoBase64("Testin.pdf")
		body := struct_models.CloudPrem{FileName: filename + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

		resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

		data, err := struct_models.UnmarshalCloudPrem(resp.Body())
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Response Deserialized successfully")
		}
		Expect(resp.StatusCode()).To(Equal(200))
		Expect(filename + ".pdf").To(Equal(data.FileName))

		//download
		downloadendpoint := viperenvcloud + "/download"
		downloadbody := struct_models.CloudPrem{FileName: filename + "abc" + ".pdf", Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

		downloadresp := utilities.Fire_post_request_bodystruct(headers, downloadbody, downloadendpoint)

		downloaddata, downloaderr := struct_models.UnmarshalCloudPrem(downloadresp.Body())

		if downloaderr != nil {
			log.Println(downloaderr)
		} else {
			log.Println("Response Deserialized successfully")
		}

		//To check wrong input - AccessDenied in error
		Expect(downloadresp.StatusCode()).To(Equal(200))
		Expect(downloaddata.FileName).Should(BeEmpty())
		Expect(downloaddata.Error).Should(ContainSubstring("AccessDenied"))
		Expect(downloaddata.Body).Should(BeEmpty())

	})

	XIt("Can send multiple pdf files to S3 bucket using load balancer API and validate all filenames through download endpoint", func() {

		filename, _ := ioutil.ReadDir("../../data/files")

		for _, file := range filename {

			filename1, _ := utilities.RandomNameandEmail()
			endpoint := viperenvcloud
			headers := utilities.Get_common_headers()
			encodedstring := utilities.ConvertFiletoBase64(file.Name())
			body := struct_models.CloudPrem{FileName: filename1 + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

			resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

			data, err := struct_models.UnmarshalCloudPrem(resp.Body())
			if err != nil {
				log.Println(err)
			}
			Expect(resp.StatusCode()).To(Equal(200))
			Expect(filename1 + ".pdf").To(Equal(data.FileName))

			//download
			downloadendpoint := viperenvcloud + "/download"
			downloadbody := struct_models.CloudPrem{FileName: filename1 + ".pdf", Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

			downloadresp := utilities.Fire_post_request_bodystruct(headers, downloadbody, downloadendpoint)

			downloaddata, downloaderr := struct_models.UnmarshalCloudPrem(downloadresp.Body())

			if downloaderr != nil {
				log.Println(downloaderr)
			}

			Expect(downloadresp.StatusCode()).To(Equal(200))
			Expect(downloaddata.FileName).To(Equal(filename1 + ".pdf"))
			Expect(downloaddata.Error).To(Equal(""))
			Expect(downloaddata.Body).ShouldNot(BeEmpty())

		}

	})

	XIt("Can send files to two different S3 buckets and receive file names corresponsing to respective deploymentid", func() {

		filename, _ := ioutil.ReadDir("../../data/files")

		for _, file := range filename {

			filename1, _ := utilities.RandomNameandEmail()
			endpoint := viperenvcloud
			headers := utilities.Get_common_headers()
			encodedstring := utilities.ConvertFiletoBase64(file.Name())
			if file.Name() == "Test2.pdf" || file.Name() == "Testin.pdf" {
				body = struct_models.CloudPrem{FileName: filename1 + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe2")}

			} else {
				body = struct_models.CloudPrem{FileName: filename1 + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}
			}

			resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

			data, err := struct_models.UnmarshalCloudPrem(resp.Body())
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Response Deserialized successfully")
			}
			Expect(resp.StatusCode()).To(Equal(200))
			Expect(filename1 + ".pdf").To(Equal(data.FileName))

			//download
			downloadendpoint := viperenvcloud + "/download"

			if file.Name() == "Test2.pdf" || file.Name() == "Testin.pdf" {
				downloadbody = struct_models.CloudPrem{FileName: filename1 + ".pdf", Deploymentid: utilities.ViperEnvVariable("deploymentidsqe2")}
			} else {
				downloadbody = struct_models.CloudPrem{FileName: filename1 + ".pdf", Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}
			}

			downloadresp := utilities.Fire_post_request_bodystruct(headers, downloadbody, downloadendpoint)

			downloaddata, downloaderr := struct_models.UnmarshalCloudPrem(downloadresp.Body())

			if downloaderr != nil {
				log.Println(downloaderr)
			} else {
				log.Println("Response Deserialized successfully")
			}

			Expect(downloadresp.StatusCode()).To(Equal(200))
			Expect(downloaddata.FileName).To(Equal(filename1 + ".pdf"))
			Expect(downloaddata.Error).To(Equal(""))
			Expect(downloaddata.Body).ShouldNot(BeEmpty())

		}

	})

	XIt("Can send different file formats to S3 bucket and download the different format files", func() {

		filename, _ := ioutil.ReadDir("../../data/allformatfiles")

		for _, file := range filename {

			fileextention := filepath.Ext(file.Name())
			filename1, _ := utilities.RandomNameandEmail()
			endpoint := viperenvcloud
			headers := utilities.Get_common_headers()
			encodedstring := utilities.ConvertAllFilestoBase64(file.Name())
			body := struct_models.CloudPrem{FileName: filename1 + fileextention, Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

			resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

			data, err := struct_models.UnmarshalCloudPrem(resp.Body())
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Response deserialized successfully for filename : " + file.Name())
			}
			Expect(resp.StatusCode()).To(Equal(200))
			Expect(filename1 + fileextention).To(Equal(data.FileName))
			Expect(data.Filelocation).To(ContainSubstring(filename1 + fileextention))

			//download
			downloadendpoint := viperenvcloud + "/download"
			downloadbody := struct_models.CloudPrem{FileName: filename1 + fileextention, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

			downloadresp := utilities.Fire_post_request_bodystruct(headers, downloadbody, downloadendpoint)

			downloaddata, downloaderr := struct_models.UnmarshalCloudPrem(downloadresp.Body())

			if downloaderr != nil {
				log.Println(downloaderr)
			} else {
				log.Println("Download response deserialized successfully for filename : " + file.Name())
			}

			Expect(downloadresp.StatusCode()).To(Equal(200))
			Expect(downloaddata.FileName).To(Equal(filename1 + fileextention))
			Expect(downloaddata.Error).To(Equal(""))
			Expect(downloaddata.Body).ShouldNot(BeEmpty())

		}

	})

})
