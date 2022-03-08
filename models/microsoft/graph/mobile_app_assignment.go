package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppAssignment provides operations to manage the deviceAppManagement singleton.
type MobileAppAssignment struct {
    Entity
    // The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
    intent *InstallIntent;
    // The settings for target assignment defined by the admin.
    settings MobileAppAssignmentSettingsable;
    // The target group assignment defined by the admin.
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewMobileAppAssignment instantiates a new mobileAppAssignment and sets the default values.
func NewMobileAppAssignment()(*MobileAppAssignment) {
    m := &MobileAppAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMobileAppAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*InstallIntent))
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppAssignmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(MobileAppAssignmentSettingsable))
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetIntent gets the intent property value. The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *MobileAppAssignment) GetIntent()(*InstallIntent) {
    if m == nil {
        return nil
    } else {
        return m.intent
    }
}
// GetSettings gets the settings property value. The settings for target assignment defined by the admin.
func (m *MobileAppAssignment) GetSettings()(MobileAppAssignmentSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetTarget gets the target property value. The target group assignment defined by the admin.
func (m *MobileAppAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *MobileAppAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MobileAppAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
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
// SetIntent sets the intent property value. The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *MobileAppAssignment) SetIntent(value *InstallIntent)() {
    if m != nil {
        m.intent = value
    }
}
// SetSettings sets the settings property value. The settings for target assignment defined by the admin.
func (m *MobileAppAssignment) SetSettings(value MobileAppAssignmentSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
// SetTarget sets the target property value. The target group assignment defined by the admin.
func (m *MobileAppAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
