package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageApprovalStageable 
type AccessPackageApprovalStageable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDurationBeforeAutomaticDenial()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetDurationBeforeEscalation()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetEscalationApprovers()([]SubjectSetable)
    GetFallbackEscalationApprovers()([]SubjectSetable)
    GetFallbackPrimaryApprovers()([]SubjectSetable)
    GetIsApproverJustificationRequired()(*bool)
    GetIsEscalationEnabled()(*bool)
    GetPrimaryApprovers()([]SubjectSetable)
    SetDurationBeforeAutomaticDenial(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetDurationBeforeEscalation(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetEscalationApprovers(value []SubjectSetable)()
    SetFallbackEscalationApprovers(value []SubjectSetable)()
    SetFallbackPrimaryApprovers(value []SubjectSetable)()
    SetIsApproverJustificationRequired(value *bool)()
    SetIsEscalationEnabled(value *bool)()
    SetPrimaryApprovers(value []SubjectSetable)()
}
