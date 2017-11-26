/*******************************************************************************
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package product
// represent the product sold by this site
type Item struct {
	Name string
	Color string
	Quantity string
}

// service offered
type ItemOperation interface {
	Buy(quantity int) bool
	Sell(quantity int) bool
	Query() int
}
// default implementation...
func (it *Item) Buy(number int) bool {
	return false
}
// default implementation..
func (it *Item) Sell(number int) bool {
	return false
}

func (it *Item) Query() int {
	return 0
}