package battle

func ConvertToSlice(listInput [][]string) []Item {
	var itemsList []Item

	for _, itemInput := range listInput {
		itemsList = append(itemsList, Item{Name: itemInput[0], Score: 0})
	}

	return itemsList
}
