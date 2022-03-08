package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityUserFlowAttributeAssignmentable 
type IdentityUserFlowAttributeAssignmentable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDisplayName()(*string)
    GetIsOptional()(*bool)
    GetRequiresVerification()(*bool)
    GetUserAttribute()(IdentityUserFlowAttributeable)
    GetUserAttributeValues()([]UserAttributeValuesItemable)
    GetUserInputType()(*IdentityUserFlowAttributeInputType)
    SetDisplayName(value *string)()
    SetIsOptional(value *bool)()
    SetRequiresVerification(value *bool)()
    SetUserAttribute(value IdentityUserFlowAttributeable)()
    SetUserAttributeValues(value []UserAttributeValuesItemable)()
    SetUserInputType(value *IdentityUserFlowAttributeInputType)()
}
