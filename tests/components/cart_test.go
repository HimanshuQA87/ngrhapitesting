package components

import (
	"fmt"
	"log"

	"github.com/HimanshuQA87/ngrhapitesting/struct_models"

	"github.com/HimanshuQA87/ngrhapitesting/utilities" // Give modulename of go.mod file along with the file from where to call the function

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/HimanshuQA87/ngrhapitesting/database"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/* InfoLogger.Println("This is some info")
WarningLogger.Println("This is probably important")
ErrorLogger.Println("Something went wrong") */

var _ = XDescribe("DB ", Label("Smoke"), func() {

	XIt("Connect DB and run select Query", func() {
		database.Connectdatabase()

	})

})

var _ = XDescribe("EC2StartStopFunctionhandler", Label("Smoke"), func() {

	BeforeEach(func() {
		log.Println("----Test Started----")
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

	XIt("aa", func() {
		// create a new AWS session

		creds := credentials.NewEnvCredentials()

		// Retrieve the credentials value
		credValue, err := creds.Get()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(credValue)
		sess, err1 := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		})
		if err1 != nil {
			fmt.Println(err1)
		}

		/* sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})) */

		// create a new Lambda client
		svc := lambda.New(sess)

		// specify the name of your Lambda function
		//functionName := "api-rh-ec2-start-function"
		functionName := "sqe-rh-api-automation"

		// specify the instance ID of the EC2 instance you want to start/stop
		instanceId := "i-078bf6ba3c665c6e8"

		// specify whether you want to start or stop the instance
		action := "start"

		// create the input payload for the Lambda function
		input := fmt.Sprintf(`{"instance_id": "%s", "action": "%s"}`, instanceId, action)

		// create the input object for the Invoke API
		invokeInput := &lambda.InvokeInput{
			FunctionName: aws.String(functionName),
			Payload:      []byte(input),
		}

		// invoke the Lambda function
		result, err := svc.Invoke(invokeInput)
		if err != nil {
			panic(err)
		}

		// print the response from the Lambda function
		fmt.Println(string(result.Payload))
	})

})

var _ = XDescribe("Add order with and without authentication", Label("Smoke"), func() {

	XIt("Post Request - Create a New order with authentication 201", func() {

		cName, _ := utilities.RandomNameandEmail()
		endpoint := viperenv + "/orders"
		token := utilities.GetAuthorizationToken()
		cart := utilities.GetCartId()

		//Post Request to add an item to cart
		endpoint1 := viperenv + "/carts/" + cart + "/items"
		headers1 := utilities.Get_common_headers()

		body1 := struct_models.AddItemToCart{
			ProductId: 4643,
			Quantity:  2}
		resp1 := utilities.Fire_post_request_bodystruct(headers1, body1, endpoint1)

		data1, err1 := struct_models.UnmarshalAddItemToCart(resp1.Body())
		if err1 != nil {
			log.Println(err1)
		}
		Expect(resp1.StatusCode()).To(Equal(201))
		Expect(data1.Created).To(BeTrue())

		cartitemid := data1.ItemId
		log.Println(cartitemid)

		headers := utilities.Get_common_headers()
		headers["Authorization"] = token
		body := struct_models.CreateNewOrder{CartID: cart, CustomerName: cName}
		resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

		data, err := struct_models.UnmarshalCreateNewOrder(resp.Body())

		if err != nil {
			log.Println(err)
		}
		Expect(resp.StatusCode()).To(Equal(201))
		Expect(data.Created).To(BeTrue())

	})

	XIt("Post Request - Pass wrong authentication token 401", func() {

		cName, _ := utilities.RandomNameandEmail()
		endpoint := viperenv + "/orders"
		token := utilities.GetAuthorizationToken()
		cart := utilities.GetCartId()

		//Post Request to add an item to cart
		endpoint1 := viperenv + "/carts/" + cart + "/items"
		headers1 := utilities.Get_common_headers()

		body1 := struct_models.AddItemToCart{
			ProductId: 4641,
			Quantity:  5}
		resp1 := utilities.Fire_post_request_bodystruct(headers1, body1, endpoint1)

		data1, err1 := struct_models.UnmarshalAddItemToCart(resp1.Body())
		if err1 != nil {
			fmt.Println(err1)
		}
		Expect(resp1.StatusCode()).To(Equal(201))
		Expect(data1.Created).To(BeTrue())

		cartitemid := data1.ItemId
		log.Println(cartitemid)

		headers := utilities.Get_common_headers()
		headers["Authorization"] = token + "abc"
		body := struct_models.CreateNewOrder{CartID: cart, CustomerName: cName}
		resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

		data, err := struct_models.UnmarshalCreateNewOrder(resp.Body())

		if err != nil {
			fmt.Println(err)
		}
		Expect(resp.StatusCode()).To(Equal(401))
		Expect(data.Error).To(Equal("Invalid bearer token."))

		content, _ := utilities.ReadFile("unauthorized.txt")
		isPayloadCorrect, _ := utilities.JSONBytesEqual([]byte(content), resp.Body())
		Expect(isPayloadCorrect).To(BeTrue(), "Payload should be equal")

	})

})
