package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsDeviceAccount struct {
    additionalData map[string]interface{};
    password *string;
}
func NewWindowsDeviceAccount()(*WindowsDeviceAccount) {
    m := &WindowsDeviceAccount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WindowsDeviceAccount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WindowsDeviceAccount) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
func (m *WindowsDeviceAccount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPassword(val)
        return nil
    }
    return res
}
func (m *WindowsDeviceAccount) IsNil()(bool) {
    return m == nil
}
func (m *WindowsDeviceAccount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
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
func (m *WindowsDeviceAccount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WindowsDeviceAccount) SetPassword(value *string)() {
    m.password = value
}
