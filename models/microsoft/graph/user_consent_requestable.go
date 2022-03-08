package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserConsentRequestable 
type UserConsentRequestable interface {
    Requestable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApproval()(Approvalable)
    GetReason()(*string)
    SetApproval(value Approvalable)()
    SetReason(value *string)()
}
