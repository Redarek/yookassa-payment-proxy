package models

type YooKassaPayment struct {
	Amount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"amount"`
	PaymentMethodData struct {
		Type string `json:"type,omitempty"`
	} `json:"payment_method_data,omitempty"`
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
			VatCode  int    `json:"vat_code"`
			Quantity string `json:"quantity"`
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
	PaymentMethod struct {
		Type  string `json:"type"`
		ID    string `json:"id"`
		Saved bool   `json:"saved"`
	} `json:"payment_method"`
	CreatedAt    string `json:"created_at"`
	Confirmation struct {
		Type            string `json:"type"`
		ReturnURL       string `json:"return_url"`
		ConfirmationURL string `json:"confirmation_url"`
	} `json:"confirmation"`
	Test       bool                   `json:"test"`
	Paid       bool                   `json:"paid"`
	Refundable bool                   `json:"refundable"`
	Metadata   map[string]interface{} `json:"metadata"`
}
