package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppAssignment struct {
    Entity
    // The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
    intent *InstallIntent;
    // The settings for target assignment defined by the admin.
    settings *MobileAppAssignmentSettings;
    // The target group assignment defined by the admin.
    target *DeviceAndAppManagementAssignmentTarget;
}
// Instantiates a new mobileAppAssignment and sets the default values.
func NewMobileAppAssignment()(*MobileAppAssignment) {
    m := &MobileAppAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the intent property value. The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *MobileAppAssignment) GetIntent()(*InstallIntent) {
    if m == nil {
        return nil
    } else {
        return m.intent
    }
}
// Gets the settings property value. The settings for target assignment defined by the admin.
func (m *MobileAppAssignment) GetSettings()(*MobileAppAssignmentSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the target property value. The target group assignment defined by the admin.
func (m *MobileAppAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// The deserialization information for the current model
func (m *MobileAppAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(InstallIntent)
            m.SetIntent(&cast)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppAssignmentSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*MobileAppAssignmentSettings))
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*DeviceAndAppManagementAssignmentTarget))
        }
        return nil
    }
    return res
}
func (m *MobileAppAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := m.GetIntent().String()
        err = writer.WriteStringValue("intent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the intent property value. The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
// Parameters:
//  - value : Value to set for the intent property.
func (m *MobileAppAssignment) SetIntent(value *InstallIntent)() {
    m.intent = value
}
// Sets the settings property value. The settings for target assignment defined by the admin.
// Parameters:
//  - value : Value to set for the settings property.
func (m *MobileAppAssignment) SetSettings(value *MobileAppAssignmentSettings)() {
    m.settings = value
}
// Sets the target property value. The target group assignment defined by the admin.
// Parameters:
//  - value : Value to set for the target property.
func (m *MobileAppAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    m.target = value
}
