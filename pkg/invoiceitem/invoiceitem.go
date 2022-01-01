package invoiceitem

//Item contains of information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

//New returns a new item

func New(id uint, product string, value float64) Item {
	return Item{
		id:      id,
		product: product,
		value:   value,
	}
}

//Value getter of Item.Value
func (i Item) Value() float64 {
	return i.value
}
