package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InvestigationSecurityState struct {
    additionalData map[string]interface{};
    name *string;
    status *string;
}
func NewInvestigationSecurityState()(*InvestigationSecurityState) {
    m := &InvestigationSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *InvestigationSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *InvestigationSecurityState) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *InvestigationSecurityState) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *InvestigationSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *InvestigationSecurityState) IsNil()(bool) {
    return m == nil
}
func (m *InvestigationSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *InvestigationSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *InvestigationSecurityState) SetName(value *string)() {
    m.name = value
}
func (m *InvestigationSecurityState) SetStatus(value *string)() {
    m.status = value
}
