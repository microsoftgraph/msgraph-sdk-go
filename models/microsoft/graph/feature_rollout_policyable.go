package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FeatureRolloutPolicyable 
type FeatureRolloutPolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAppliesTo()([]DirectoryObjectable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFeature()(*StagedFeatureName)
    GetIsAppliedToOrganization()(*bool)
    GetIsEnabled()(*bool)
    SetAppliesTo(value []DirectoryObjectable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFeature(value *StagedFeatureName)()
    SetIsAppliedToOrganization(value *bool)()
    SetIsEnabled(value *bool)()
}
