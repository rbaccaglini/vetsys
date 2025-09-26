package domain

import "errors"

// ErrCustomerNotFound is returned when the client does not exist in the repository.
var ErrCustomerNotFound = errors.New("customer not found")
