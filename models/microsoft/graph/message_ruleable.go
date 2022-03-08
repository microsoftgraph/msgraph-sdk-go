package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageRuleable 
type MessageRuleable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetActions()(MessageRuleActionsable)
    GetConditions()(MessageRulePredicatesable)
    GetDisplayName()(*string)
    GetExceptions()(MessageRulePredicatesable)
    GetHasError()(*bool)
    GetIsEnabled()(*bool)
    GetIsReadOnly()(*bool)
    GetSequence()(*int32)
    SetActions(value MessageRuleActionsable)()
    SetConditions(value MessageRulePredicatesable)()
    SetDisplayName(value *string)()
    SetExceptions(value MessageRulePredicatesable)()
    SetHasError(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsReadOnly(value *bool)()
    SetSequence(value *int32)()
}
