package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InsightIdentity struct {
    additionalData map[string]interface{};
    address *string;
    displayName *string;
    id *string;
}
func NewInsightIdentity()(*InsightIdentity) {
    m := &InsightIdentity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *InsightIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *InsightIdentity) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *InsightIdentity) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *InsightIdentity) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *InsightIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddress(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    return res
}
func (m *InsightIdentity) IsNil()(bool) {
    return m == nil
}
func (m *InsightIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *InsightIdentity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *InsightIdentity) SetAddress(value *string)() {
    m.address = value
}
func (m *InsightIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *InsightIdentity) SetId(value *string)() {
    m.id = value
}
