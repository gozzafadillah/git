package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func main() {
	e := echo.New()
	e.GET("/data", xendit_invoice)
	e.Start(":8080")

}

func xendit_invoice(ctx echo.Context) error {

	xendit.Opt.SecretKey = "xnd_development_I0guK5bOcQB3AVQ8pYUXMtXltsVvfqsyPU2dz1RJvTDNVrsLVxqC8K0KJc3YhlZE"
	code := uuid.New()

	customer := xendit.Customer{
		GivenNames:   "John",
		Surname:      "Doe",
		Email:        "gozzafadillah@gmail.com",
		MobileNumber: "+62895631948686",
	}
	invoiceCustomer := xendit.InvoiceCustomer{
		GivenNames:   customer.GivenNames,
		Email:        customer.Email,
		MobileNumber: customer.MobileNumber,
	}
	item := xendit.InvoiceItem{
		Name:     "Paket XL 20RB",
		Price:    20000,
		Quantity: 1,
		Category: "Top Up",
	}
	fee := xendit.InvoiceFee{
		Type:  "ADMIN",
		Value: 2000,
	}

	NotificationType := []string{"whatsapp", "email", "sms"}
	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType,
		InvoiceReminder: NotificationType,
		InvoicePaid:     NotificationType,
		InvoiceExpired:  NotificationType,
	}

	data := invoice.CreateParams{
		ExternalID:                     "demo_" + code.String(),
		Amount:                         item.Price + fee.Value,
		PayerEmail:                     customer.Email,
		Description:                    "Paket Data XL",
		Items:                          []xendit.InvoiceItem{item},
		Customer:                       invoiceCustomer,
		Fees:                           []xendit.InvoiceFee{fee},
		CustomerNotificationPreference: customerNotificationPreference,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created invoice: %+v\n", resp)
	return ctx.JSON(200, map[string]interface{}{
		"message": "success",
		"payment": resp,
	})
}
