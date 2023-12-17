package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layer_time int) int {
	if layer_time == 0 {
		return len(layers) * 2
	}
	return layer_time * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var grams int = 0
	var liters float64 = 0.0
	for _, v := range layers {
		if v == "noodles" {
			grams += 50
		}
		if v == "sauce" {

			liters += 0.2
		}
	}

	return grams, liters
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friends_list []string, myList []string) {
	myList[len(myList) - 1] = friends_list[len(friends_list)-1]
	//myList = append(myList[len(friends_list)-1], friends_list[len(friends_list)-1])
	//return myList
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var new_quantities []float64
	for _, v := range quantities {
		new_quantities = append(new_quantities, v*float64(portions)/2)
	}
	return new_quantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
