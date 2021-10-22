package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ImplicitGrantSettings struct {
    additionalData map[string]interface{};
    enableAccessTokenIssuance *bool;
    enableIdTokenIssuance *bool;
}
func NewImplicitGrantSettings()(*ImplicitGrantSettings) {
    m := &ImplicitGrantSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ImplicitGrantSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ImplicitGrantSettings) GetEnableAccessTokenIssuance()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableAccessTokenIssuance
    }
}
func (m *ImplicitGrantSettings) GetEnableIdTokenIssuance()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableIdTokenIssuance
    }
}
func (m *ImplicitGrantSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enableAccessTokenIssuance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableAccessTokenIssuance(val)
        return nil
    }
    res["enableIdTokenIssuance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableIdTokenIssuance(val)
        return nil
    }
    return res
}
func (m *ImplicitGrantSettings) IsNil()(bool) {
    return m == nil
}
func (m *ImplicitGrantSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableAccessTokenIssuance", m.GetEnableAccessTokenIssuance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableIdTokenIssuance", m.GetEnableIdTokenIssuance())
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
func (m *ImplicitGrantSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ImplicitGrantSettings) SetEnableAccessTokenIssuance(value *bool)() {
    m.enableAccessTokenIssuance = value
}
func (m *ImplicitGrantSettings) SetEnableIdTokenIssuance(value *bool)() {
    m.enableIdTokenIssuance = value
}
