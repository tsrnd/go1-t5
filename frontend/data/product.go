package data

type Product struct {
	Model
	Name          string  `schema:"name"`
	Quantity      int     `schema:"quantity"`
	Price         float64 `schema:"price"`
	CategoryID    uint    `schema:"category_id"`
	Description   string  `schema:"description"`
	Category      *Category
	OrderProducts []OrderProduct //has many order products
	Images        []Image        //has many image
}

type Image struct {
	Model
	Name      string   `schema:"name"`
	URL       string   `schema:"url"`
	ProductId uint     `schema:"product_id"`
	Product   *Product //belong to Product
}

const IMG_BASE_URL = "uploads/images"

// Create a new product, save product info into the database
func (product *Product) Create() (err error) {
	return
}

// Delete product from database
func (product *Product) Delete() (err error) {
	return
}

// Update product information in the database
func (product *Product) Update() (err error) {
	return
}

// Delete all product from database
func ProductDeleteAll() (err error) {
	return
}

// Get all product in the database and returns it
func Products() (products []Product, err error) {
	return
}
