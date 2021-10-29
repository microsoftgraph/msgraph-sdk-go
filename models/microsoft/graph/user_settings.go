package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserSettings struct {
    Entity
    // Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
    contributionToContentDiscoveryAsOrganizationDisabled *bool;
    // When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
    contributionToContentDiscoveryDisabled *bool;
    // The shift preferences for the user.
    shiftPreferences *ShiftPreferences;
}
// Instantiates a new userSettings and sets the default values.
func NewUserSettings()(*UserSettings) {
    m := &UserSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
func (m *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contributionToContentDiscoveryAsOrganizationDisabled
    }
}
// Gets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
func (m *UserSettings) GetContributionToContentDiscoveryDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contributionToContentDiscoveryDisabled
    }
}
// Gets the shiftPreferences property value. The shift preferences for the user.
func (m *UserSettings) GetShiftPreferences()(*ShiftPreferences) {
    if m == nil {
        return nil
    } else {
        return m.shiftPreferences
    }
}
// The deserialization information for the current model
func (m *UserSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contributionToContentDiscoveryAsOrganizationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetContributionToContentDiscoveryAsOrganizationDisabled(val)
        return nil
    }
    res["contributionToContentDiscoveryDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetContributionToContentDiscoveryDisabled(val)
        return nil
    }
    res["shiftPreferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftPreferences() })
        if err != nil {
            return err
        }
        m.SetShiftPreferences(val.(*ShiftPreferences))
        return nil
    }
    return res
}
func (m *UserSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryAsOrganizationDisabled", m.GetContributionToContentDiscoveryAsOrganizationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryDisabled", m.GetContributionToContentDiscoveryDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shiftPreferences", m.GetShiftPreferences())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
// Parameters:
//  - value : Value to set for the contributionToContentDiscoveryAsOrganizationDisabled property.
func (m *UserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)() {
    m.contributionToContentDiscoveryAsOrganizationDisabled = value
}
// Sets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
// Parameters:
//  - value : Value to set for the contributionToContentDiscoveryDisabled property.
func (m *UserSettings) SetContributionToContentDiscoveryDisabled(value *bool)() {
    m.contributionToContentDiscoveryDisabled = value
}
// Sets the shiftPreferences property value. The shift preferences for the user.
// Parameters:
//  - value : Value to set for the shiftPreferences property.
func (m *UserSettings) SetShiftPreferences(value *ShiftPreferences)() {
    m.shiftPreferences = value
}
