package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RequiredResourceAccess struct {
    additionalData map[string]interface{};
    resourceAccess []ResourceAccess;
    resourceAppId *string;
}
func NewRequiredResourceAccess()(*RequiredResourceAccess) {
    m := &RequiredResourceAccess{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RequiredResourceAccess) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RequiredResourceAccess) GetResourceAccess()([]ResourceAccess) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccess
    }
}
func (m *RequiredResourceAccess) GetResourceAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceAppId
    }
}
func (m *RequiredResourceAccess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resourceAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceAccess() })
        if err != nil {
            return err
        }
        res := make([]ResourceAccess, len(val))
        for i, v := range val {
            res[i] = *(v.(*ResourceAccess))
        }
        m.SetResourceAccess(res)
        return nil
    }
    res["resourceAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceAppId(val)
        return nil
    }
    return res
}
func (m *RequiredResourceAccess) IsNil()(bool) {
    return m == nil
}
func (m *RequiredResourceAccess) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceAccess()))
        for i, v := range m.GetResourceAccess() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("resourceAccess", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceAppId", m.GetResourceAppId())
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
func (m *RequiredResourceAccess) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RequiredResourceAccess) SetResourceAccess(value []ResourceAccess)() {
    m.resourceAccess = value
}
func (m *RequiredResourceAccess) SetResourceAppId(value *string)() {
    m.resourceAppId = value
}
