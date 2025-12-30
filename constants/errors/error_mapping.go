package errors

import "slices"

func ErrMapping(err error) bool {
	allErrors := append(GeneralErrors[:], UserErrors[:]...)
	return slices.Contains(allErrors, err)
}
