package common

type Address struct {
	ZIPCode      string `bson:"zip_code"`
	City         string `bson:"city"`
	Street       string `bson:"street"`
	Number       string `bson:"number"`
	Neighborhood string `bson:"neighborhood"`
}

type Phone struct {
	Phone   string `bson:"phone"`
	Primary bool   `bson:"primary"`
}

type Email struct {
	Email   string `bson:"email"`
	Primary bool   `bson:"primary"`
}
