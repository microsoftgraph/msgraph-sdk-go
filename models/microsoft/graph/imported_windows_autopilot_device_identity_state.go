package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ImportedWindowsAutopilotDeviceIdentityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Device error code reported by Device Directory Service(DDS).
    deviceErrorCode *int32;
    // Device error name reported by Device Directory Service(DDS).
    deviceErrorName *string;
    // Device status reported by Device Directory Service(DDS). Possible values are: unknown, pending, partial, complete, error.
    deviceImportStatus *ImportedWindowsAutopilotDeviceIdentityImportStatus;
    // Device Registration ID for successfully added device reported by Device Directory Service(DDS).
    deviceRegistrationId *string;
}
// Instantiates a new importedWindowsAutopilotDeviceIdentityState and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentityState()(*ImportedWindowsAutopilotDeviceIdentityState) {
    m := &ImportedWindowsAutopilotDeviceIdentityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceErrorCode property value. Device error code reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceErrorCode
    }
}
// Gets the deviceErrorName property value. Device error name reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceErrorName
    }
}
// Gets the deviceImportStatus property value. Device status reported by Device Directory Service(DDS). Possible values are: unknown, pending, partial, complete, error.
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceImportStatus()(*ImportedWindowsAutopilotDeviceIdentityImportStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceImportStatus
    }
}
// Gets the deviceRegistrationId property value. Device Registration ID for successfully added device reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceRegistrationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegistrationId
    }
}
// The deserialization information for the current model
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceErrorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceErrorCode(val)
        return nil
    }
    res["deviceErrorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceErrorName(val)
        return nil
    }
    res["deviceImportStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportedWindowsAutopilotDeviceIdentityImportStatus)
        if err != nil {
            return err
        }
        cast := val.(ImportedWindowsAutopilotDeviceIdentityImportStatus)
        m.SetDeviceImportStatus(&cast)
        return nil
    }
    res["deviceRegistrationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceRegistrationId(val)
        return nil
    }
    return res
}
func (m *ImportedWindowsAutopilotDeviceIdentityState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ImportedWindowsAutopilotDeviceIdentityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("deviceErrorCode", m.GetDeviceErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceErrorName", m.GetDeviceErrorName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceImportStatus() != nil {
        cast := m.GetDeviceImportStatus().String()
        err := writer.WriteStringValue("deviceImportStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceRegistrationId", m.GetDeviceRegistrationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceErrorCode property value. Device error code reported by Device Directory Service(DDS).
// Parameters:
//  - value : Value to set for the deviceErrorCode property.
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceErrorCode(value *int32)() {
    m.deviceErrorCode = value
}
// Sets the deviceErrorName property value. Device error name reported by Device Directory Service(DDS).
// Parameters:
//  - value : Value to set for the deviceErrorName property.
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceErrorName(value *string)() {
    m.deviceErrorName = value
}
// Sets the deviceImportStatus property value. Device status reported by Device Directory Service(DDS). Possible values are: unknown, pending, partial, complete, error.
// Parameters:
//  - value : Value to set for the deviceImportStatus property.
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceImportStatus(value *ImportedWindowsAutopilotDeviceIdentityImportStatus)() {
    m.deviceImportStatus = value
}
// Sets the deviceRegistrationId property value. Device Registration ID for successfully added device reported by Device Directory Service(DDS).
// Parameters:
//  - value : Value to set for the deviceRegistrationId property.
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceRegistrationId(value *string)() {
    m.deviceRegistrationId = value
}
