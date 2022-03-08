package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningStepable 
type ProvisioningStepable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDescription()(*string)
    GetDetails()(DetailsInfoable)
    GetName()(*string)
    GetProvisioningStepType()(*ProvisioningStepType)
    GetStatus()(*ProvisioningResult)
    SetDescription(value *string)()
    SetDetails(value DetailsInfoable)()
    SetName(value *string)()
    SetProvisioningStepType(value *ProvisioningStepType)()
    SetStatus(value *ProvisioningResult)()
}
