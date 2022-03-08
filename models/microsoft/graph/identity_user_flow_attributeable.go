package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityUserFlowAttributeable 
type IdentityUserFlowAttributeable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDataType()(*IdentityUserFlowAttributeDataType)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetUserFlowAttributeType()(*IdentityUserFlowAttributeType)
    SetDataType(value *IdentityUserFlowAttributeDataType)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetUserFlowAttributeType(value *IdentityUserFlowAttributeType)()
}
