package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FeatureRolloutPolicy provides operations to manage the policyRoot singleton.
type FeatureRolloutPolicy struct {
    Entity
    // Nullable. Specifies a list of directoryObjects that feature is enabled for.
    appliesTo []DirectoryObjectable;
    // A description for this feature rollout policy.
    description *string;
    // The display name for this  feature rollout policy.
    displayName *string;
    // Possible values are: passthroughAuthentication, seamlessSso, passwordHashSync, emailAsAlternateId, unknownFutureValue.
    feature *StagedFeatureName;
    // Indicates whether this feature rollout policy should be applied to the entire organization.
    isAppliedToOrganization *bool;
    // Indicates whether the feature rollout is enabled.
    isEnabled *bool;
}
// NewFeatureRolloutPolicy instantiates a new featureRolloutPolicy and sets the default values.
func NewFeatureRolloutPolicy()(*FeatureRolloutPolicy) {
    m := &FeatureRolloutPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFeatureRolloutPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureRolloutPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFeatureRolloutPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. Nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *FeatureRolloutPolicy) GetAppliesTo()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetDescription gets the description property value. A description for this feature rollout policy.
func (m *FeatureRolloutPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for this  feature rollout policy.
func (m *FeatureRolloutPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFeature gets the feature property value. Possible values are: passthroughAuthentication, seamlessSso, passwordHashSync, emailAsAlternateId, unknownFutureValue.
func (m *FeatureRolloutPolicy) GetFeature()(*StagedFeatureName) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureRolloutPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetAppliesTo(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseStagedFeatureName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val.(*StagedFeatureName))
        }
        return nil
    }
    res["isAppliedToOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAppliedToOrganization(val)
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsAppliedToOrganization gets the isAppliedToOrganization property value. Indicates whether this feature rollout policy should be applied to the entire organization.
func (m *FeatureRolloutPolicy) GetIsAppliedToOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAppliedToOrganization
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the feature rollout is enabled.
func (m *FeatureRolloutPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
func (m *FeatureRolloutPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FeatureRolloutPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliesTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetFeature() != nil {
        cast := (*m.GetFeature()).String()
        err = writer.WriteStringValue("feature", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAppliedToOrganization", m.GetIsAppliedToOrganization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. Nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *FeatureRolloutPolicy) SetAppliesTo(value []DirectoryObjectable)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetDescription sets the description property value. A description for this feature rollout policy.
func (m *FeatureRolloutPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for this  feature rollout policy.
func (m *FeatureRolloutPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFeature sets the feature property value. Possible values are: passthroughAuthentication, seamlessSso, passwordHashSync, emailAsAlternateId, unknownFutureValue.
func (m *FeatureRolloutPolicy) SetFeature(value *StagedFeatureName)() {
    if m != nil {
        m.feature = value
    }
}
// SetIsAppliedToOrganization sets the isAppliedToOrganization property value. Indicates whether this feature rollout policy should be applied to the entire organization.
func (m *FeatureRolloutPolicy) SetIsAppliedToOrganization(value *bool)() {
    if m != nil {
        m.isAppliedToOrganization = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the feature rollout is enabled.
func (m *FeatureRolloutPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
