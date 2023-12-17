package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	uints := map[string]int{}
	uints["quarter_of_a_dozen"] = 3
	uints["half_of_a_dozen"] = 6
	uints["dozen"] = 12
	uints["small_gross"] = 120
	uints["gross"] = 144
	uints["great_gross"] = 1728
	return uints
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	uint_v, ok := units[unit]

	if !ok {
		return false
	}

	v, ok := bill[item]
	if ok {
		bill[item] = v + uint_v
	} else {
		bill[item] = uint_v
	}

	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}
	if bill[item] >= value {
		bill[item] -= value
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
	return false

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, ok := bill[item]
	return v, ok
}
