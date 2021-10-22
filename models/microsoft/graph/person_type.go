package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PersonType struct {
    additionalData map[string]interface{};
    class *string;
    subclass *string;
}
func NewPersonType()(*PersonType) {
    m := &PersonType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PersonType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PersonType) GetClass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.class
    }
}
func (m *PersonType) GetSubclass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subclass
    }
}
func (m *PersonType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["class"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClass(val)
        return nil
    }
    res["subclass"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubclass(val)
        return nil
    }
    return res
}
func (m *PersonType) IsNil()(bool) {
    return m == nil
}
func (m *PersonType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("class", m.GetClass())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subclass", m.GetSubclass())
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
func (m *PersonType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PersonType) SetClass(value *string)() {
    m.class = value
}
func (m *PersonType) SetSubclass(value *string)() {
    m.subclass = value
}
