package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportedWindowsAutopilotDeviceIdentityState provides operations to manage the deviceManagement singleton.
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
// NewImportedWindowsAutopilotDeviceIdentityState instantiates a new importedWindowsAutopilotDeviceIdentityState and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentityState()(*ImportedWindowsAutopilotDeviceIdentityState) {
    m := &ImportedWindowsAutopilotDeviceIdentityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateImportedWindowsAutopilotDeviceIdentityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedWindowsAutopilotDeviceIdentityStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewImportedWindowsAutopilotDeviceIdentityState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceErrorCode gets the deviceErrorCode property value. Device error code reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceErrorCode
    }
}
// GetDeviceErrorName gets the deviceErrorName property value. Device error name reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceErrorName
    }
}
// GetDeviceImportStatus gets the deviceImportStatus property value. Device status reported by Device Directory Service(DDS). Possible values are: unknown, pending, partial, complete, error.
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceImportStatus()(*ImportedWindowsAutopilotDeviceIdentityImportStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceImportStatus
    }
}
// GetDeviceRegistrationId gets the deviceRegistrationId property value. Device Registration ID for successfully added device reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetDeviceRegistrationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegistrationId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedWindowsAutopilotDeviceIdentityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceErrorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceErrorCode(val)
        }
        return nil
    }
    res["deviceErrorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceErrorName(val)
        }
        return nil
    }
    res["deviceImportStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportedWindowsAutopilotDeviceIdentityImportStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceImportStatus(val.(*ImportedWindowsAutopilotDeviceIdentityImportStatus))
        }
        return nil
    }
    res["deviceRegistrationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationId(val)
        }
        return nil
    }
    return res
}
func (m *ImportedWindowsAutopilotDeviceIdentityState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetDeviceImportStatus()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceErrorCode sets the deviceErrorCode property value. Device error code reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceErrorCode(value *int32)() {
    if m != nil {
        m.deviceErrorCode = value
    }
}
// SetDeviceErrorName sets the deviceErrorName property value. Device error name reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceErrorName(value *string)() {
    if m != nil {
        m.deviceErrorName = value
    }
}
// SetDeviceImportStatus sets the deviceImportStatus property value. Device status reported by Device Directory Service(DDS). Possible values are: unknown, pending, partial, complete, error.
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceImportStatus(value *ImportedWindowsAutopilotDeviceIdentityImportStatus)() {
    if m != nil {
        m.deviceImportStatus = value
    }
}
// SetDeviceRegistrationId sets the deviceRegistrationId property value. Device Registration ID for successfully added device reported by Device Directory Service(DDS).
func (m *ImportedWindowsAutopilotDeviceIdentityState) SetDeviceRegistrationId(value *string)() {
    if m != nil {
        m.deviceRegistrationId = value
    }
}
