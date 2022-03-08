package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePlanInfoable 
type ServicePlanInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAppliesTo()(*string)
    GetProvisioningStatus()(*string)
    GetServicePlanId()(*string)
    GetServicePlanName()(*string)
    SetAppliesTo(value *string)()
    SetProvisioningStatus(value *string)()
    SetServicePlanId(value *string)()
    SetServicePlanName(value *string)()
}
