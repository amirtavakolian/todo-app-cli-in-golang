package validation

type CategoryValidation struct {
}

func NewCategoryValidation() CategoryValidation {
	return CategoryValidation{}
}

func (validation CategoryValidation) ValidateCategoryName(categoryName string) (bool, string) {

	if len(categoryName) <= 2 {
		return false, "Category name must be more then 2 charecters"
	}

	return true, ""
}
