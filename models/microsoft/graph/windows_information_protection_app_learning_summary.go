package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// windowsInformationProtectionAppLearningSummary 
type WindowsInformationProtectionAppLearningSummary struct {
    Entity
    // Application Name
    applicationName *string;
    // Application Type. Possible values are: universal, desktop.
    applicationType *ApplicationType;
    // Device Count
    deviceCount *int32;
}
// NewWindowsInformationProtectionAppLearningSummary instantiates a new windowsInformationProtectionAppLearningSummary and sets the default values.
func NewWindowsInformationProtectionAppLearningSummary()(*WindowsInformationProtectionAppLearningSummary) {
    m := &WindowsInformationProtectionAppLearningSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicationName gets the applicationName property value. Application Name
func (m *WindowsInformationProtectionAppLearningSummary) GetApplicationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationName
    }
}
// GetApplicationType gets the applicationType property value. Application Type. Possible values are: universal, desktop.
func (m *WindowsInformationProtectionAppLearningSummary) GetApplicationType()(*ApplicationType) {
    if m == nil {
        return nil
    } else {
        return m.applicationType
    }
}
// GetDeviceCount gets the deviceCount property value. Device Count
func (m *WindowsInformationProtectionAppLearningSummary) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionAppLearningSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
        }
        return nil
    }
    res["applicationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ApplicationType)
            m.SetApplicationType(&cast)
        }
        return nil
    }
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionAppLearningSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionAppLearningSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationName", m.GetApplicationName())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationType() != nil {
        cast := m.GetApplicationType().String()
        err = writer.WriteStringValue("applicationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationName sets the applicationName property value. Application Name
func (m *WindowsInformationProtectionAppLearningSummary) SetApplicationName(value *string)() {
    m.applicationName = value
}
// SetApplicationType sets the applicationType property value. Application Type. Possible values are: universal, desktop.
func (m *WindowsInformationProtectionAppLearningSummary) SetApplicationType(value *ApplicationType)() {
    m.applicationType = value
}
// SetDeviceCount sets the deviceCount property value. Device Count
func (m *WindowsInformationProtectionAppLearningSummary) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
