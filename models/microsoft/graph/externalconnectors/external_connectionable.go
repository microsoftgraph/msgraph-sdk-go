package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// ExternalConnectionable 
type ExternalConnectionable interface {
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfiguration()(Configurationable)
    GetDescription()(*string)
    GetGroups()([]ExternalGroupable)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetSchema()(Schemaable)
    GetState()(*ConnectionState)
    SetConfiguration(value Configurationable)()
    SetDescription(value *string)()
    SetGroups(value []ExternalGroupable)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetSchema(value Schemaable)()
    SetState(value *ConnectionState)()
}
