package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/config"
	"ecommerce/web/helpers"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

type PaymentHandler struct {
	PaymentUsecase usecasesinterface.PaymentUsecaseInterface
	OrderUsecase   usecasesinterface.OrderUsecaseInterface
	CouponUsecase  usecasesinterface.CouponUsecaseInterface
	Cfg            config.Config
}

func NewPaymentHandler(usecase usecasesinterface.PaymentUsecaseInterface, orderUsecase usecasesinterface.OrderUsecaseInterface, couponusecase usecasesinterface.CouponUsecaseInterface, cfg config.Config) *PaymentHandler {
	return &PaymentHandler{
		PaymentUsecase: usecase,
		OrderUsecase:   orderUsecase,
		CouponUsecase:  couponusecase,
		Cfg:            cfg,
	}
}

func (payment *PaymentHandler) MakePayment(c *gin.Context) {

	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	//get the status of order
	if err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "cant find order id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	orderdata, err := payment.OrderUsecase.GetOrderByID(uint(orderId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "cant find data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if orderdata.RecievedPayment {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "already payed",
			Data:       nil,
			Errors:     nil,
		})
		return
	}

	client := razorpay.NewClient(payment.Cfg.RAZORPAYID, payment.Cfg.RAZORPAYSECRET)

	data := map[string]interface{}{
		"amount":   orderdata.Price * 100,
		"currency": "INR",
		"receipt":  "test_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "cant process order right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	value := body["id"]
	razorpayID := value.(string)

	c.HTML(200, "payment.html", gin.H{
		"total_price": orderdata.Price,
		"total":       orderdata.Price,
		"orderData":   orderId,
		"orderid":     razorpayID,
		"amount":      orderdata.Price,
		"email":       orderdata.Email,
	})

}

func (cr *PaymentHandler) PaymentSuccess(c *gin.Context) {

	paymentRef := c.Query("payment_ref")
	fmt.Println("paymentRef from query :", paymentRef)

	idStr := c.Query("order_id")
	fmt.Print("order id from query _:", idStr)

	idStr = strings.ReplaceAll(idStr, " ", "")

	orderID, err := strconv.Atoi(idStr)
	fmt.Println("_converted order  id from query :", orderID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "Can't find orderId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID := c.Query("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "Can't find UserId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	t := c.Query("total")
	fmt.Println("total from query :", t)
	total, err := strconv.ParseInt(t, 10, 64)
	fmt.Println("total from query converted:", total)

	if err != nil {
		//	handle err
		fmt.Println("failed to fetch order id")
	}

	if err := cr.OrderUsecase.ToggleReceivedPayment(uint(orderID)); err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 500,
			Message:    "payment coouldnt be processed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	paymentVerifier := helperstructs.PaymentReq{
		Email:         userID,
		OrderID:       uint(orderID),
		PaymentRef:    paymentRef,
		TotalPrice:    int(total),
		PaymentStatus: "SUCCESS",
	}

	fmt.Println("payment verifier in handler : ", paymentVerifier)
	//paymentVerifier.
	err = cr.PaymentUsecase.Makepayment(paymentVerifier)
	if err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "faild to update payment",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	coupons, err := cr.CouponUsecase.ListofCouponsAvailableForThisOrder(paymentVerifier.TotalPrice)

	if err != nil {
		fmt.Println("here occured a coupon error....", err)
	}

	i := helpers.SelectRandomintBetweenRange(0, (len(coupons) - 1))

	if err := cr.CouponUsecase.CreditUserWithCoupon(paymentVerifier.Email, coupons[i]); err != nil {
		fmt.Println("here also occured a coupon error ....", err)
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "payment updated",
		Data:       nil,
		Errors:     nil,
	})

}
