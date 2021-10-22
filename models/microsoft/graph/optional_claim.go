package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OptionalClaim struct {
    additionalData map[string]interface{};
    additionalProperties []string;
    essential *bool;
    name *string;
    source *string;
}
func NewOptionalClaim()(*OptionalClaim) {
    m := &OptionalClaim{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OptionalClaim) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OptionalClaim) GetAdditionalProperties()([]string) {
    if m == nil {
        return nil
    } else {
        return m.additionalProperties
    }
}
func (m *OptionalClaim) GetEssential()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.essential
    }
}
func (m *OptionalClaim) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *OptionalClaim) GetSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
func (m *OptionalClaim) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAdditionalProperties(res)
        return nil
    }
    res["essential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEssential(val)
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
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSource(val)
        return nil
    }
    return res
}
func (m *OptionalClaim) IsNil()(bool) {
    return m == nil
}
func (m *OptionalClaim) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("additionalProperties", m.GetAdditionalProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("essential", m.GetEssential())
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
        err := writer.WriteStringValue("source", m.GetSource())
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
func (m *OptionalClaim) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OptionalClaim) SetAdditionalProperties(value []string)() {
    m.additionalProperties = value
}
func (m *OptionalClaim) SetEssential(value *bool)() {
    m.essential = value
}
func (m *OptionalClaim) SetName(value *string)() {
    m.name = value
}
func (m *OptionalClaim) SetSource(value *string)() {
    m.source = value
}
