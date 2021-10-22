package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemPreviewInfo struct {
    additionalData map[string]interface{};
    getUrl *string;
    postParameters *string;
    postUrl *string;
}
func NewItemPreviewInfo()(*ItemPreviewInfo) {
    m := &ItemPreviewInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ItemPreviewInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ItemPreviewInfo) GetGetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.getUrl
    }
}
func (m *ItemPreviewInfo) GetPostParameters()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postParameters
    }
}
func (m *ItemPreviewInfo) GetPostUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postUrl
    }
}
func (m *ItemPreviewInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["getUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGetUrl(val)
        return nil
    }
    res["postParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostParameters(val)
        return nil
    }
    res["postUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostUrl(val)
        return nil
    }
    return res
}
func (m *ItemPreviewInfo) IsNil()(bool) {
    return m == nil
}
func (m *ItemPreviewInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("getUrl", m.GetGetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postParameters", m.GetPostParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postUrl", m.GetPostUrl())
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
func (m *ItemPreviewInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ItemPreviewInfo) SetGetUrl(value *string)() {
    m.getUrl = value
}
func (m *ItemPreviewInfo) SetPostParameters(value *string)() {
    m.postParameters = value
}
func (m *ItemPreviewInfo) SetPostUrl(value *string)() {
    m.postUrl = value
}
