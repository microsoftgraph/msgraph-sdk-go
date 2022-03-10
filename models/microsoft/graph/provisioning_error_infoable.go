package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningErrorInfoable 
type ProvisioningErrorInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalDetails()(*string)
    GetErrorCategory()(*ProvisioningStatusErrorCategory)
    GetErrorCode()(*string)
    GetReason()(*string)
    GetRecommendedAction()(*string)
    SetAdditionalDetails(value *string)()
    SetErrorCategory(value *ProvisioningStatusErrorCategory)()
    SetErrorCode(value *string)()
    SetReason(value *string)()
    SetRecommendedAction(value *string)()
}
