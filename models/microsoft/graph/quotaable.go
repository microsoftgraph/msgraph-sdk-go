package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Quotaable 
type Quotaable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeleted()(*int32)
    GetRemaining()(*int32)
    GetState()(*string)
    GetStoragePlanInformation()(StoragePlanInformationable)
    GetTotal()(*int32)
    GetUsed()(*int32)
    SetDeleted(value *int32)()
    SetRemaining(value *int32)()
    SetState(value *string)()
    SetStoragePlanInformation(value StoragePlanInformationable)()
    SetTotal(value *int32)()
    SetUsed(value *int32)()
}
