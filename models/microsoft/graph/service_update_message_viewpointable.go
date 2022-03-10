package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceUpdateMessageViewpointable 
type ServiceUpdateMessageViewpointable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIsArchived()(*bool)
    GetIsFavorited()(*bool)
    GetIsRead()(*bool)
    SetIsArchived(value *bool)()
    SetIsFavorited(value *bool)()
    SetIsRead(value *bool)()
}
