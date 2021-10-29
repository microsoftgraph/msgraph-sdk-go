package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsInformationProtectionAppLearningSummary struct {
    Entity
    // Application Name
    applicationName *string;
    // Application Type. Possible values are: universal, desktop.
    applicationType *ApplicationType;
    // Device Count
    deviceCount *int32;
}
// Instantiates a new windowsInformationProtectionAppLearningSummary and sets the default values.
func NewWindowsInformationProtectionAppLearningSummary()(*WindowsInformationProtectionAppLearningSummary) {
    m := &WindowsInformationProtectionAppLearningSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicationName property value. Application Name
func (m *WindowsInformationProtectionAppLearningSummary) GetApplicationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationName
    }
}
// Gets the applicationType property value. Application Type. Possible values are: universal, desktop.
func (m *WindowsInformationProtectionAppLearningSummary) GetApplicationType()(*ApplicationType) {
    if m == nil {
        return nil
    } else {
        return m.applicationType
    }
}
// Gets the deviceCount property value. Device Count
func (m *WindowsInformationProtectionAppLearningSummary) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// The deserialization information for the current model
func (m *WindowsInformationProtectionAppLearningSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationName(val)
        return nil
    }
    res["applicationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationType)
        if err != nil {
            return err
        }
        cast := val.(ApplicationType)
        m.SetApplicationType(&cast)
        return nil
    }
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceCount(val)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionAppLearningSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the applicationName property value. Application Name
// Parameters:
//  - value : Value to set for the applicationName property.
func (m *WindowsInformationProtectionAppLearningSummary) SetApplicationName(value *string)() {
    m.applicationName = value
}
// Sets the applicationType property value. Application Type. Possible values are: universal, desktop.
// Parameters:
//  - value : Value to set for the applicationType property.
func (m *WindowsInformationProtectionAppLearningSummary) SetApplicationType(value *ApplicationType)() {
    m.applicationType = value
}
// Sets the deviceCount property value. Device Count
// Parameters:
//  - value : Value to set for the deviceCount property.
func (m *WindowsInformationProtectionAppLearningSummary) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
