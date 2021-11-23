package cleanwindowsdevice

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CleanWindowsDeviceRequestBody 
type CleanWindowsDeviceRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keepUserData *bool;
}
// NewCleanWindowsDeviceRequestBody instantiates a new cleanWindowsDeviceRequestBody and sets the default values.
func NewCleanWindowsDeviceRequestBody()(*CleanWindowsDeviceRequestBody) {
    m := &CleanWindowsDeviceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CleanWindowsDeviceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKeepUserData gets the keepUserData property value. 
func (m *CleanWindowsDeviceRequestBody) GetKeepUserData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepUserData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CleanWindowsDeviceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keepUserData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    return res
}
func (m *CleanWindowsDeviceRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CleanWindowsDeviceRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CleanWindowsDeviceRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKeepUserData sets the keepUserData property value. 
func (m *CleanWindowsDeviceRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
