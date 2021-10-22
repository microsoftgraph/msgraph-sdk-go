package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AddIn struct {
    additionalData map[string]interface{};
    id *string;
    properties []KeyValue;
    type_escpaped *string;
}
func NewAddIn()(*AddIn) {
    m := &AddIn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AddIn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AddIn) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *AddIn) GetProperties()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
func (m *AddIn) GetType_escpaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *AddIn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        res := make([]KeyValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValue))
        }
        m.SetProperties(res)
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
func (m *AddIn) IsNil()(bool) {
    return m == nil
}
func (m *AddIn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("properties", cast)
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
func (m *AddIn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AddIn) SetId(value *string)() {
    m.id = value
}
func (m *AddIn) SetProperties(value []KeyValue)() {
    m.properties = value
}
func (m *AddIn) SetType_escpaped(value *string)() {
    m.type_escpaped = value
}
