package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExtensionPropertyable 
type ExtensionPropertyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    DirectoryObjectable
    GetAppDisplayName()(*string)
    GetDataType()(*string)
    GetIsSyncedFromOnPremises()(*bool)
    GetName()(*string)
    GetTargetObjects()([]string)
    SetAppDisplayName(value *string)()
    SetDataType(value *string)()
    SetIsSyncedFromOnPremises(value *bool)()
    SetName(value *string)()
    SetTargetObjects(value []string)()
}
