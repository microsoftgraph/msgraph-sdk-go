package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintTaskable 
type PrintTaskable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDefinition()(PrintTaskDefinitionable)
    GetParentUrl()(*string)
    GetStatus()(PrintTaskStatusable)
    GetTrigger()(PrintTaskTriggerable)
    SetDefinition(value PrintTaskDefinitionable)()
    SetParentUrl(value *string)()
    SetStatus(value PrintTaskStatusable)()
    SetTrigger(value PrintTaskTriggerable)()
}
