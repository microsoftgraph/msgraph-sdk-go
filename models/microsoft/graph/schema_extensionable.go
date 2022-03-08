package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SchemaExtensionable 
type SchemaExtensionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDescription()(*string)
    GetOwner()(*string)
    GetProperties()([]ExtensionSchemaPropertyable)
    GetStatus()(*string)
    GetTargetTypes()([]string)
    SetDescription(value *string)()
    SetOwner(value *string)()
    SetProperties(value []ExtensionSchemaPropertyable)()
    SetStatus(value *string)()
    SetTargetTypes(value []string)()
}
