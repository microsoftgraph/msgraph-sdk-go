package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnPremisesConditionalAccessSettings struct {
    Entity
    // Indicates if on premises conditional access is enabled for this organization
    enabled *bool;
    // User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
    excludedGroups []string;
    // User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
    includedGroups []string;
    // Override the default access rule when allowing a device to ensure access is granted.
    overrideDefaultRule *bool;
}
// Instantiates a new onPremisesConditionalAccessSettings and sets the default values.
func NewOnPremisesConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    m := &OnPremisesConditionalAccessSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the enabled property value. Indicates if on premises conditional access is enabled for this organization
func (m *OnPremisesConditionalAccessSettings) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// Gets the excludedGroups property value. User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
func (m *OnPremisesConditionalAccessSettings) GetExcludedGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludedGroups
    }
}
// Gets the includedGroups property value. User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
func (m *OnPremisesConditionalAccessSettings) GetIncludedGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includedGroups
    }
}
// Gets the overrideDefaultRule property value. Override the default access rule when allowing a device to ensure access is granted.
func (m *OnPremisesConditionalAccessSettings) GetOverrideDefaultRule()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overrideDefaultRule
    }
}
// The deserialization information for the current model
func (m *OnPremisesConditionalAccessSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnabled(val)
        return nil
    }
    res["excludedGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludedGroups(res)
        return nil
    }
    res["includedGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludedGroups(res)
        return nil
    }
    res["overrideDefaultRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOverrideDefaultRule(val)
        return nil
    }
    return res
}
func (m *OnPremisesConditionalAccessSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnPremisesConditionalAccessSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("excludedGroups", m.GetExcludedGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("includedGroups", m.GetIncludedGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("overrideDefaultRule", m.GetOverrideDefaultRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the enabled property value. Indicates if on premises conditional access is enabled for this organization
// Parameters:
//  - value : Value to set for the enabled property.
func (m *OnPremisesConditionalAccessSettings) SetEnabled(value *bool)() {
    m.enabled = value
}
// Sets the excludedGroups property value. User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
// Parameters:
//  - value : Value to set for the excludedGroups property.
func (m *OnPremisesConditionalAccessSettings) SetExcludedGroups(value []string)() {
    m.excludedGroups = value
}
// Sets the includedGroups property value. User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
// Parameters:
//  - value : Value to set for the includedGroups property.
func (m *OnPremisesConditionalAccessSettings) SetIncludedGroups(value []string)() {
    m.includedGroups = value
}
// Sets the overrideDefaultRule property value. Override the default access rule when allowing a device to ensure access is granted.
// Parameters:
//  - value : Value to set for the overrideDefaultRule property.
func (m *OnPremisesConditionalAccessSettings) SetOverrideDefaultRule(value *bool)() {
    m.overrideDefaultRule = value
}
