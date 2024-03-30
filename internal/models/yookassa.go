package models

type YooKassaPayment struct {
	Amount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Confirmation struct {
		Type      string `json:"type,omitempty"`
		ReturnURL string `json:"return_url,omitempty"`
	} `json:"confirmation"`
	Receipt struct {
		Customer struct {
			FullName string `json:"full_name,omitempty"`
			Email    string `json:"email,omitempty"`
			Phone    string `json:"phone,omitempty"`
		} `json:"customer,omitempty"`
		Items []struct {
			Description string `json:"description"`
			Amount      struct {
				Value    string `json:"value"`
				Currency string `json:"currency"`
			} `json:"amount"`
			VatCode        string `json:"vat_code"`
			Quantity       string `json:"quantity"`
			PaymentMode    string `json:"payment_mode,omitempty"`
			PaymentSubject string `json:"payment_subject,omitempty"`
		} `json:"items"`
		Email string `json:"email,omitempty"`
		Phone string `json:"phone,omitempty"`
	} `json:"receipt,omitempty"`
	Description string `json:"description,omitempty"`
}

type YooKassaPaymentResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Amount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Description string `json:"description"`
	Recipient   struct {
		AccountID string `json:"account_id"`
		GatewayID string `json:"gateway_id"`
	} `json:"recipient"`
	CreatedAt    string `json:"created_at"`
	Confirmation struct {
		Type              string `json:"type"`
		ConfirmationToken string `json:"confirmation_token"`
	} `json:"confirmation"`
	Test       bool                   `json:"test"`
	Paid       bool                   `json:"paid"`
	Refundable bool                   `json:"refundable"`
	Metadata   map[string]interface{} `json:"metadata"`
}
