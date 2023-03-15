package models

type Config struct {
	PaymentDetails PaymentDetails `json:"payment_details"`
	Output         Output         `json:"output"`
}

type Output struct {
	FolderPath string `json:"folder_path"`
	FileType   string `json:"file_type"`
}

type PaymentDetails struct {
	Company Company `json:"company"`
	Issuer  Issuer  `json:"issuer"`
	Works   []*Work `json:"works"`
}

type Company struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	Representative string `json:"representative"`
	Email          string `json:"email"`
	Website        string `json:"website"`
}

type Issuer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}

type Work struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Currency    string  `json:"currency"`
}
