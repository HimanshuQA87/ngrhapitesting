package integration

import (
	"log"

	"github.com/HimanshuQA87/ngrhapitesting/struct_models"

	"github.com/HimanshuQA87/ngrhapitesting/utilities" // Give modulename of go.mod file along with the file from where to call the function

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var viperenv string = utilities.ViperEnvVariable("baseURL")

//var viperenv1 string = utilities.ViperEnvVariable("baseURL1")

var _ = XDescribe("Create a Cart , add Item to cart , get added item from cart", Label("Smoke"), func() {

	BeforeEach(func() {
		log.Println("----Test Started---- ")
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

	XIt("Post ", func() {
		//Post Request to create a cart
		endpoint := viperenv + "/carts"
		headers := utilities.Get_common_headers()

		resp := utilities.Fire_post_request(headers, "", endpoint)

		data, err := struct_models.UnmarshalCreateCart(resp.Body())
		if err != nil {
			log.Println(err)
		}
		Expect(data.Created).To(BeTrue(), "Cart not created - did not return true")
		Expect(resp.StatusCode()).To(Equal(201), "Status code not equal to 201")

		cartid := data.CartID // Store cart id in a variable
		log.Println(cartid)

		//**************************//
		//Post Request to add an item to cart
		endpoint1 := viperenv + "/carts/" + cartid + "/items"
		headers1 := utilities.Get_common_headers()

		body := struct_models.AddItemToCart{
			ProductId: 4643,
			Quantity:  2}
		resp1 := utilities.Fire_post_request_bodystruct(headers1, body, endpoint1)

		data1, err1 := struct_models.UnmarshalAddItemToCart(resp1.Body())
		if err1 != nil {
			log.Println(err1)
		}
		Expect(resp1.StatusCode()).To(Equal(201))
		Expect(data1.Created).To(BeTrue())

		cartitemid := data1.ItemId
		log.Println(cartitemid)
		//*************************************//
		//Get Request to verify the added item in cart
		endpoint2 := viperenv + "/carts/" + cartid + "/items"
		headers2 := utilities.Get_common_headers()

		resp2 := utilities.Fire_get_request(headers2, "", endpoint2)

		data2, err2 := struct_models.UnmarshalGetCartItems(resp2.Body())
		if err2 != nil {
			log.Println(err2)
		}

		Expect(resp2.StatusCode()).To(Equal(200))
		Expect(data2[0].Id).To(Equal(cartitemid))
		Expect(data2[0].ProductId).To(Equal(4643))
		Expect(data2[0].Quantity).To(Equal(2))
		//*****************************************//

	})

})
