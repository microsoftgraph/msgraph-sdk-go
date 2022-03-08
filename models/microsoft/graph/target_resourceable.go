package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TargetResourceable 
type TargetResourceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDisplayName()(*string)
    GetGroupType()(*GroupType)
    GetId()(*string)
    GetModifiedProperties()([]ModifiedPropertyable)
    GetType()(*string)
    GetUserPrincipalName()(*string)
    SetDisplayName(value *string)()
    SetGroupType(value *GroupType)()
    SetId(value *string)()
    SetModifiedProperties(value []ModifiedPropertyable)()
    SetType(value *string)()
    SetUserPrincipalName(value *string)()
}
