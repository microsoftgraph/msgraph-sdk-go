package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InferenceClassificationOverrideable 
type InferenceClassificationOverrideable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClassifyAs()(*InferenceClassificationType)
    GetSenderEmailAddress()(EmailAddressable)
    SetClassifyAs(value *InferenceClassificationType)()
    SetSenderEmailAddress(value EmailAddressable)()
}
