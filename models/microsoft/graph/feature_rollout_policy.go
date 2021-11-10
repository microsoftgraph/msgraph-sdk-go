package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type FeatureRolloutPolicy struct {
    Entity
    // Nullable. Specifies a list of directoryObjects that feature is enabled for.
    appliesTo []DirectoryObject;
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
// Instantiates a new featureRolloutPolicy and sets the default values.
func NewFeatureRolloutPolicy()(*FeatureRolloutPolicy) {
    m := &FeatureRolloutPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appliesTo property value. Nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *FeatureRolloutPolicy) GetAppliesTo()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// Gets the description property value. A description for this feature rollout policy.
func (m *FeatureRolloutPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for this  feature rollout policy.
func (m *FeatureRolloutPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the feature property value. Possible values are: passthroughAuthentication, seamlessSso, passwordHashSync, emailAsAlternateId, unknownFutureValue.
func (m *FeatureRolloutPolicy) GetFeature()(*StagedFeatureName) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
// Gets the isAppliedToOrganization property value. Indicates whether this feature rollout policy should be applied to the entire organization.
func (m *FeatureRolloutPolicy) GetIsAppliedToOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAppliedToOrganization
    }
}
// Gets the isEnabled property value. Indicates whether the feature rollout is enabled.
func (m *FeatureRolloutPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// The deserialization information for the current model
func (m *FeatureRolloutPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
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
            cast := val.(StagedFeatureName)
            m.SetFeature(&cast)
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
func (m *FeatureRolloutPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *FeatureRolloutPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := m.GetFeature().String()
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
// Sets the appliesTo property value. Nullable. Specifies a list of directoryObjects that feature is enabled for.
// Parameters:
//  - value : Value to set for the appliesTo property.
func (m *FeatureRolloutPolicy) SetAppliesTo(value []DirectoryObject)() {
    m.appliesTo = value
}
// Sets the description property value. A description for this feature rollout policy.
// Parameters:
//  - value : Value to set for the description property.
func (m *FeatureRolloutPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for this  feature rollout policy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *FeatureRolloutPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the feature property value. Possible values are: passthroughAuthentication, seamlessSso, passwordHashSync, emailAsAlternateId, unknownFutureValue.
// Parameters:
//  - value : Value to set for the feature property.
func (m *FeatureRolloutPolicy) SetFeature(value *StagedFeatureName)() {
    m.feature = value
}
// Sets the isAppliedToOrganization property value. Indicates whether this feature rollout policy should be applied to the entire organization.
// Parameters:
//  - value : Value to set for the isAppliedToOrganization property.
func (m *FeatureRolloutPolicy) SetIsAppliedToOrganization(value *bool)() {
    m.isAppliedToOrganization = value
}
// Sets the isEnabled property value. Indicates whether the feature rollout is enabled.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *FeatureRolloutPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
