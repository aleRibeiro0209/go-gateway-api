package domain

import "errors"

var (
	// ErrAccountNotFound é retornado quando uma conta não é encontrada.
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicateAPIKey é retornado quando há tentativa de criar conta com API key duplicada.
	ErrDuplicateAPIKey = errors.New("API key already exists")
	// ErrInvoiceNotFound é retornado quando uma fatura não é encontrada.
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAccess é retornado quando há tentativa de acesso não autorizado a um recurso.
	ErrUnauthorizedAccess = errors.New("unauthorized access")
	// ErrInvalidAmount é retornado quando o valor do campo amount é inválido.
	ErrInvalidAmount = errors.New("invalid amount")
	// ErrInvalidStatus é retornado quando o valor do campo status é inválido.
	ErrInvalidStatus = errors.New("invalid status")
)