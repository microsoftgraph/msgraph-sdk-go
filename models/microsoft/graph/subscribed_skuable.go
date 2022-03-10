package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubscribedSkuable 
type SubscribedSkuable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppliesTo()(*string)
    GetCapabilityStatus()(*string)
    GetConsumedUnits()(*int32)
    GetPrepaidUnits()(LicenseUnitsDetailable)
    GetServicePlans()([]ServicePlanInfoable)
    GetSkuId()(*string)
    GetSkuPartNumber()(*string)
    SetAppliesTo(value *string)()
    SetCapabilityStatus(value *string)()
    SetConsumedUnits(value *int32)()
    SetPrepaidUnits(value LicenseUnitsDetailable)()
    SetServicePlans(value []ServicePlanInfoable)()
    SetSkuId(value *string)()
    SetSkuPartNumber(value *string)()
}
