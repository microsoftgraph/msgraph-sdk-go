package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LicenseProcessingState struct {
    additionalData map[string]interface{};
    state *string;
}
func NewLicenseProcessingState()(*LicenseProcessingState) {
    m := &LicenseProcessingState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LicenseProcessingState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LicenseProcessingState) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *LicenseProcessingState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    return res
}
func (m *LicenseProcessingState) IsNil()(bool) {
    return m == nil
}
func (m *LicenseProcessingState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *LicenseProcessingState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LicenseProcessingState) SetState(value *string)() {
    m.state = value
}
