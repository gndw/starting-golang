package constants

const (
	TransactionTypeIn  TransactionType = 1
	TransactionTypeOut TransactionType = 2
)

const (
	TransactionStatusPending   TransactionStatus = 1
	TransactionStatusProcessed TransactionStatus = 2
	TransactionStatusFailed    TransactionStatus = 3
)

const (
	BankTransferStatusSuccess BankTransferStatus = "SUCCESS"
	BankTransferStatusFailed  BankTransferStatus = "FAILED"
)
