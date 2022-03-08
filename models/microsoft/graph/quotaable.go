package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Quotaable 
type Quotaable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDeleted()(*int64)
    GetRemaining()(*int64)
    GetState()(*string)
    GetStoragePlanInformation()(StoragePlanInformationable)
    GetTotal()(*int64)
    GetUsed()(*int64)
    SetDeleted(value *int64)()
    SetRemaining(value *int64)()
    SetState(value *string)()
    SetStoragePlanInformation(value StoragePlanInformationable)()
    SetTotal(value *int64)()
    SetUsed(value *int64)()
}
