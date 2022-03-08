package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResourceOperationable 
type ResourceOperationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetActionName()(*string)
    GetDescription()(*string)
    GetResourceName()(*string)
    SetActionName(value *string)()
    SetDescription(value *string)()
    SetResourceName(value *string)()
}
