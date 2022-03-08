package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppDiagnosticStatusable 
type ManagedAppDiagnosticStatusable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetMitigationInstruction()(*string)
    GetState()(*string)
    GetValidationName()(*string)
    SetMitigationInstruction(value *string)()
    SetState(value *string)()
    SetValidationName(value *string)()
}
