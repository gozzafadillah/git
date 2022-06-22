package main

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/retailoutlet"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_I0guK5bOcQB3AVQ8pYUXMtXltsVvfqsyPU2dz1RJvTDNVrsLVxqC8K0KJc3YhlZE"

	data := retailoutlet.CreateFixedPaymentCodeParams{
		ExternalID:       "retailoutlet-external-id",
		RetailOutletName: xendit.RetailOutletNameAlfamart,
		Name:             "Michael Jackson",
		ExpectedAmount:   200000,
	}

	resp, err := retailoutlet.CreateFixedPaymentCode(&data)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("created retail outlet fixed payment code: %+v\n", resp)

}
