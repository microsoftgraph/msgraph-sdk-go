package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintTaskDefinitionable 
type PrintTaskDefinitionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCreatedBy()(AppIdentityable)
    GetDisplayName()(*string)
    GetTasks()([]PrintTaskable)
    SetCreatedBy(value AppIdentityable)()
    SetDisplayName(value *string)()
    SetTasks(value []PrintTaskable)()
}
