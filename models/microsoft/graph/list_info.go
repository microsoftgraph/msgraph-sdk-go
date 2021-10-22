package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ListInfo struct {
    additionalData map[string]interface{};
    contentTypesEnabled *bool;
    hidden *bool;
    template *string;
}
func NewListInfo()(*ListInfo) {
    m := &ListInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ListInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ListInfo) GetContentTypesEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contentTypesEnabled
    }
}
func (m *ListInfo) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
func (m *ListInfo) GetTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.template
    }
}
func (m *ListInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentTypesEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetContentTypesEnabled(val)
        return nil
    }
    res["hidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHidden(val)
        return nil
    }
    res["template"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTemplate(val)
        return nil
    }
    return res
}
func (m *ListInfo) IsNil()(bool) {
    return m == nil
}
func (m *ListInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("contentTypesEnabled", m.GetContentTypesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hidden", m.GetHidden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("template", m.GetTemplate())
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
func (m *ListInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ListInfo) SetContentTypesEnabled(value *bool)() {
    m.contentTypesEnabled = value
}
func (m *ListInfo) SetHidden(value *bool)() {
    m.hidden = value
}
func (m *ListInfo) SetTemplate(value *string)() {
    m.template = value
}
