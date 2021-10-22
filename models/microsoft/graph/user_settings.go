package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserSettings struct {
    Entity
    contributionToContentDiscoveryAsOrganizationDisabled *bool;
    contributionToContentDiscoveryDisabled *bool;
    shiftPreferences *ShiftPreferences;
}
func NewUserSettings()(*UserSettings) {
    m := &UserSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contributionToContentDiscoveryAsOrganizationDisabled
    }
}
func (m *UserSettings) GetContributionToContentDiscoveryDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contributionToContentDiscoveryDisabled
    }
}
func (m *UserSettings) GetShiftPreferences()(*ShiftPreferences) {
    if m == nil {
        return nil
    } else {
        return m.shiftPreferences
    }
}
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
func (m *UserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)() {
    m.contributionToContentDiscoveryAsOrganizationDisabled = value
}
func (m *UserSettings) SetContributionToContentDiscoveryDisabled(value *bool)() {
    m.contributionToContentDiscoveryDisabled = value
}
func (m *UserSettings) SetShiftPreferences(value *ShiftPreferences)() {
    m.shiftPreferences = value
}
