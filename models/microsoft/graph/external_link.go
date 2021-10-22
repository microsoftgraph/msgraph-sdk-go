package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExternalLink struct {
    additionalData map[string]interface{};
    href *string;
}
func NewExternalLink()(*ExternalLink) {
    m := &ExternalLink{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExternalLink) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ExternalLink) GetHref()(*string) {
    if m == nil {
        return nil
    } else {
        return m.href
    }
}
func (m *ExternalLink) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["href"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHref(val)
        return nil
    }
    return res
}
func (m *ExternalLink) IsNil()(bool) {
    return m == nil
}
func (m *ExternalLink) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("href", m.GetHref())
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
func (m *ExternalLink) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExternalLink) SetHref(value *string)() {
    m.href = value
}
