package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

// ExternalConnectionable 
type ExternalConnectionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfiguration()(Configurationable)
    GetDescription()(*string)
    GetGroups()([]ExternalGroupable)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetSchema()(Schemaable)
    GetState()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionState)
    SetConfiguration(value Configurationable)()
    SetDescription(value *string)()
    SetGroups(value []ExternalGroupable)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetSchema(value Schemaable)()
    SetState(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionState)()
}
