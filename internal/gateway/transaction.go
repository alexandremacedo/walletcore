package gateway

import "github.com.br/alexandremacedo/walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
