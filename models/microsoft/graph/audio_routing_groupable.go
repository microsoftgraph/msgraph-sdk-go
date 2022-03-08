package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AudioRoutingGroupable 
type AudioRoutingGroupable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetReceivers()([]string)
    GetRoutingMode()(*RoutingMode)
    GetSources()([]string)
    SetReceivers(value []string)()
    SetRoutingMode(value *RoutingMode)()
    SetSources(value []string)()
}
