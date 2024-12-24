package main

import "net/http"

func main() {
	// map url path to a function
	http.HandleFunc("/login", login)

	// for authorizing payment
	http.HandleFunc("/customer/payment/authorize", customerPayment)
	// then we capture it
	http.HandleFunc("/customer/payment/capture", customerPaymentCapture)
	http.HandleFunc("/customer/ledger", customerLedger)
}

// responseWrite will create response to the user
func login(http.ResponseWriter, *http.Request) {

}

func customerPayment(http.ResponseWriter, *http.request) {

}
func customerPaymentCapture(http.ResponseWriter, *http.request) {

}
func customerLedger(http.ResponseWriter, *http.request) {

}
