package payment

import (
	"capstone/be/delivery/common"
	_paymentRepo "capstone/be/repository/payment"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

type PaymentController struct {
	repository _paymentRepo.Payment
}

func New(payment _paymentRepo.Payment) *PaymentController {
	return &PaymentController{repository: payment}
}

func (pc PaymentController) Charge() echo.HandlerFunc {
	return func(c echo.Context) error {
		xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRETKEY")

		availableBanks, err := virtualaccount.GetAvailableBanks()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("available va banks: %+v\n", availableBanks)

		createFixedVAData := virtualaccount.CreateFixedVAParams{
			ExternalID: "va-" + time.Now().String(),
			BankCode:   availableBanks[1].Code,
			Name:       "Michael Jackson",
		}

		resp, err := virtualaccount.CreateFixedVA(&createFixedVAData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created fixed va: %+v\n", resp)

		getFixedVAData := virtualaccount.GetFixedVAParams{
			ID: resp.ID,
		}

		resp, err = virtualaccount.GetFixedVA(&getFixedVAData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("retrieved fixed va: %+v\n", resp)

		expirationDate := time.Now().AddDate(0, 0, 1)

		updateFixedVAData := virtualaccount.UpdateFixedVAParams{
			ID:             resp.ID,
			ExpirationDate: &expirationDate,
		}

		resp, err = virtualaccount.UpdateFixedVA(&updateFixedVAData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("updated fixed va: %+v\n", resp)

		// Before running this example, create a fixed virtual account payment simulation
		// by making a POST request to
		// https://api.xendit.co/callback_virtual_accounts/external_id=<FVA external ID>/simulate_payment
		payment, err := virtualaccount.GetPayment(&virtualaccount.GetPaymentParams{
			PaymentID: "VA_fixed-1580285972",
		})

		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, common.PaymentResponse(payment))
	}
}
