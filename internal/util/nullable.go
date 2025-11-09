package util

func GeneralNullable[T any](v T) *T {
	return &v
}

func GeneralNullableCollection[T any](args ...T) []*T {
	var res []*T
	for _, val := range args {
		res = append(res, GeneralNullable(val))
	}

	return res
}

func GeneralNonNullable[T any](v *T) T {
	if v != nil {
		return *v
	}

	var temp T
	return temp
}
