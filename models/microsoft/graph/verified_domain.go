package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type VerifiedDomain struct {
    additionalData map[string]interface{};
    capabilities *string;
    isDefault *bool;
    isInitial *bool;
    name *string;
    type_escpaped *string;
}
func NewVerifiedDomain()(*VerifiedDomain) {
    m := &VerifiedDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *VerifiedDomain) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *VerifiedDomain) GetCapabilities()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
func (m *VerifiedDomain) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
func (m *VerifiedDomain) GetIsInitial()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInitial
    }
}
func (m *VerifiedDomain) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *VerifiedDomain) GetType_escpaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *VerifiedDomain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["capabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCapabilities(val)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["isInitial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInitial(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escpaped(val)
        return nil
    }
    return res
}
func (m *VerifiedDomain) IsNil()(bool) {
    return m == nil
}
func (m *VerifiedDomain) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("capabilities", m.GetCapabilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInitial", m.GetIsInitial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escpaped", m.GetType_escpaped())
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
func (m *VerifiedDomain) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *VerifiedDomain) SetCapabilities(value *string)() {
    m.capabilities = value
}
func (m *VerifiedDomain) SetIsDefault(value *bool)() {
    m.isDefault = value
}
func (m *VerifiedDomain) SetIsInitial(value *bool)() {
    m.isInitial = value
}
func (m *VerifiedDomain) SetName(value *string)() {
    m.name = value
}
func (m *VerifiedDomain) SetType_escpaped(value *string)() {
    m.type_escpaped = value
}
