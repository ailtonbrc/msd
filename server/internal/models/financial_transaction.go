package models

import (
	"time"

	"gorm.io/gorm"
)

// FinancialTransaction representa uma transação financeira
type FinancialTransaction struct {
	gorm.Model

	TransactionType string         `gorm:"size:20;not null" json:"transaction_type"` // 'receita', 'despesa'
	Amount          float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Description     string         `json:"description"`
	ReferenceID     *uint          `json:"reference_id"`                  // ID da venda, compra, etc.
	ReferenceType   string         `gorm:"size:20" json:"reference_type"` // 'venda', 'compra', 'despesa', etc.
	TransactionDate time.Time      `json:"transaction_date"`
	PaymentMethodID *uint          `json:"payment_method_id"`
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"payment_method,omitempty"`
	Status          string         `gorm:"size:20;not null" json:"status"` // 'pendente', 'pago', 'cancelado'
	DueDate         *time.Time     `json:"due_date"`
	PaymentDate     *time.Time     `json:"payment_date"`
	CreatedByID     *uint          `gorm:"column:created_by" json:"created_by"`
	CreatedBy       *User          `gorm:"foreignKey:CreatedByID" json:"created_by_user,omitempty"`
}

// TableName especifica o nome da tabela
func (FinancialTransaction) TableName() string {
	return "financial_transactions"
}
