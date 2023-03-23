package main

import (
	"fmt"
	"time"
)

func main() {

	var customerName string

	//creating food data structure using struct

	type Item struct {
		itemId       int
		itemQuantity int
		itemPrice    int
		itemName     string
	}
	//adding elements to into the menulist
	menuList := []Item{
		{1, 1, 10, "IDLY"}, {2, 1, 10, "VADAI"}, {3, 1, 12, "DOSA"}, {4, 1, 15, "POORI"}, {5, 1, 25, "PONGAL"}, {6, 1, 25, "KICHADI"}, {7, 1, 15, "CHAPPATHI"}, {8, 1, 15, "PAROTTA"}, {9, 1, 30, "PULAO"},
		{10, 1, 30, "ROAST"}, {11, 1, 35, "GHEE ROAST"}, {12, 1, 35, "ONION ROAST"}, {13, 1, 25, "UTHAAPPAM"}, {14, 1, 30, "GHEE RICE"}, {15, 1, 20, "ADAI"}, {16, 1, 25, "HONEY ADAI"}, {17, 1, 15, "KESARI"}, {18, 1, 12, "SAMBAR-VADAI"}, {19, 1, 15, "PAYASAM"},
		{20, 1, 15, "FILTER COFFEE"}, {21, 1, 15, "TEA"}, {22, 1, 12, "MILK"}, {23, 1, 30, "FRUIT-SALAD"}, {24, 1, 25, "ROSE MILK"}, {25, 1, 25, "BADAM GHIR"},
	}
	//greeting message

	fmt.Printf("\n******************************************\n")
	fmt.Println("\nWelcome to Raajbhavan Restaurent")
	fmt.Printf("\n******************************************\n\n")
	fmt.Printf("Enter Your Name: ")
	fmt.Scan(&customerName)

	//getting customer data

	fmt.Printf("\n\nWelcome you Dear Customer %v to our RaajBhavan restaurant\n", customerName)
	fmt.Printf("\nWe are here to serve you !!\n\n\n")
	//Displaying the menulist
	fmt.Printf("Kindly look at our catalogue:\n")
	fmt.Printf("\n\n*************************************************************")
	fmt.Printf("\n                         RAAJBHAVAN                         ")
	fmt.Printf("\n                         **********                         \n")

	fmt.Print("\n Id         Quantity           Price         Item Name ")
	fmt.Printf("\n*************************************************************\n")

	for _, item := range menuList {
		if item.itemId < 10 {
			fmt.Printf("  %v            %v                 %v            %v\n", item.itemId, item.itemQuantity, item.itemPrice, item.itemName)
		} else {
			fmt.Printf(" %v            %v                 %v            %v\n", item.itemId, item.itemQuantity, item.itemPrice, item.itemName)

		}
	}
	fmt.Printf("\n*************************************************************\n")
	//asking user to select and order food items

	var userSelection int = 1
	var selectedItemId int
	var itemQuantitySelected int
	var totalAmount int = 0
	var orderConfirmation int
	var payment int = 0
	orderList := make([]Item, 0)
	for userSelection == 1 {

		fmt.Printf("\n\nKindly Enter the Item ID : ")
		fmt.Scan(&selectedItemId)
		fmt.Printf("\n\nEnter the quantity of the item you ordered: ")
		fmt.Scan(&itemQuantitySelected)
		var newOrder = Item{
			itemId:       selectedItemId,
			itemQuantity: itemQuantitySelected,
			itemPrice:    menuList[selectedItemId-1].itemPrice,
			itemName:     menuList[selectedItemId-1].itemName,
		}
		//adding the new order to the orderlist slice
		orderList = append(orderList, newOrder)

		fmt.Printf("\nWant to continue ordering press 1 ")
		fmt.Printf("\nWant to complete your ordering, press 2 ")
		fmt.Printf("\nTo cancel ordering press 3 :\n\n ")
		fmt.Scan(&userSelection)
	}
	if userSelection == 2 {
		fmt.Printf("Your orders are : ")

		//displaying the orders to the customer
		fmt.Printf("\n\n*************************************************************")
		fmt.Printf("\n                         RAAJBHAVAN                         ")
		fmt.Printf("\n                         **********                         \n")

		fmt.Print("\n Id         Quantity           Price         Item Name ")
		fmt.Printf("\n*************************************************************\n")

		for _, item := range orderList {
			//calculating the total amount

			totalAmount = totalAmount + item.itemPrice*item.itemQuantity

			if item.itemId < 10 {
				fmt.Printf("  %v            %v                 %v            %v \n", item.itemId, item.itemQuantity, item.itemPrice, item.itemName)
			} else {
				fmt.Printf(" %v            %v                 %v            %v \n", item.itemId, item.itemQuantity, item.itemPrice, item.itemName)
			}
		}
		fmt.Printf("\n Total Amount:  %v Rs/.\n", totalAmount)
		fmt.Printf("**************************************************************\n")

		//confirming the ordered items

		var orderConfirmation int
		fmt.Printf("\n\nTo confirm your order, press 1 \n")
		fmt.Printf("To cancel the order, press 0 : \n")
		fmt.Scan(&orderConfirmation)

		if orderConfirmation == 1 {

			//displaying payment options
			var paymentMethod int
			fmt.Printf("\nSelect your payment option \n")
			fmt.Printf("\nPress 1 for online Payment")
			fmt.Printf("\nPress 2 for Credit or Debit card Payment: ")
			fmt.Printf("\nPress 3 for Hard cash Payment ")
			fmt.Printf("\nPress 4 for cancel the order : \n")
			fmt.Scan(&paymentMethod)

			var cancelPayment int = 1

			switch paymentMethod {

			case 1:
				var selection int
				fmt.Printf("\n\nTo scan QR code press 1  ")
				fmt.Printf("\nTo cancel payment press 0 : \n")
				fmt.Scan(&selection)

				if selection == 1 {
					fmt.Printf("\n\nKindly scan our QR code placed on your table\n")
					time.Sleep(5 * time.Second)
					fmt.Printf("\n\nProcessing Your payment....\n")
					time.Sleep(10 * time.Second)
					fmt.Printf("\nYour payment has completed successfully!\n")
					payment = 1
				}

				if selection == 0 {

					cancelPayment = 0
				}

			case 2:
				var selection int
				fmt.Printf("\n\nTo provide card details press 1  ")
				fmt.Printf("\nTo cancel payment press 0 : \n")
				fmt.Scan(&selection)

				for selection == 1 && payment != 1 {

					var cardNumber, cardPin string
					fmt.Printf("\n\nKindly Enter your Credit or Debit card number :")
					fmt.Scan(&cardNumber)
					fmt.Printf("\nEnter Your Pin :")
					fmt.Scan(&cardPin)

					if len(cardNumber) == 16 && len(cardPin) == 4 {
						fmt.Printf("\n\nProcessing Your card.....\n")
						time.Sleep(10 * time.Second)
						fmt.Printf("\nYour payment has completed successfully!\n\n")
						payment = 1
					} else {
						fmt.Printf("Enter a valid card number and try again !!")
						payment = 0
					}
				}
				if selection == 0 {
					fmt.Printf("\n\nYour Payment is cancelled\n")
					cancelPayment = 0
				}

			case 3:

				var selection int
				for selection != 1 && cancelPayment != 0 {
					fmt.Printf("\n\nYou have selected the Hard cash payment\n\n")
					fmt.Printf("Press 1 and our team will receive the payment from you \n")
					fmt.Printf("Press 0 to cancel order\n ")
					fmt.Scan(&selection)

					if selection != 1 && selection != 0 {
						fmt.Printf("\nKindly select the correct option!!!\n")
					}

					if selection == 1 {
						fmt.Printf("\n\nOur team will receive payment shortly!!...\n")
						payment = 1
					} else {

						fmt.Printf("\nYour order is cancelled!\n")
						cancelPayment = 0

					}

				}

			}
			//payment menthod has completed

			if payment == 1 {

				//order completion message

				fmt.Printf("\n\nThanks for making order on RaajBhavan !!\n")
				fmt.Printf("Your order will be served soon!!....\n\n")
			}

		}
	}
	if userSelection == 3 || (orderConfirmation == 0 || payment == 0) {
		fmt.Printf("\n\nThanks for Visiting RaajBhavan !!!\n")
	}

}
