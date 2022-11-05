package model

// mock database
// var DB = []map[string]string{{"1": "laptop"}, {"2": "AC"}}

type Product struct {
	Name     string `json:"product_name"`
	ID       int    `json:"id"`
	SerialNo int    `json:"serialnumber"`
}

var DB = []Product{
	{
		Name:     "AC",
		ID:       1,
		SerialNo: 999999,
	},
	{
		Name:     "Laptop",
		ID:       2,
		SerialNo: 88888,
	},
	{
		Name:     "Table",
		ID:       3,
		SerialNo: 7777,
	},
}
