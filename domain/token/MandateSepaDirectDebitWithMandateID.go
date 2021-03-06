// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// MandateSepaDirectDebitWithMandateID represents class MandateSepaDirectDebitWithMandateId
type MandateSepaDirectDebitWithMandateID struct {
	BankAccountIban            *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	CustomerContractIdentifier *string                      `json:"customerContractIdentifier,omitempty"`
	Debtor                     *Debtor                      `json:"debtor,omitempty"`
	IsRecurring                *bool                        `json:"isRecurring,omitempty"`
	MandateApproval            *MandateApproval             `json:"mandateApproval,omitempty"`
	MandateID                  *string                      `json:"mandateId,omitempty"`
	PreNotification            *string                      `json:"preNotification,omitempty"`
}

// NewMandateSepaDirectDebitWithMandateID constructs a new MandateSepaDirectDebitWithMandateID
func NewMandateSepaDirectDebitWithMandateID() *MandateSepaDirectDebitWithMandateID {
	return &MandateSepaDirectDebitWithMandateID{}
}
