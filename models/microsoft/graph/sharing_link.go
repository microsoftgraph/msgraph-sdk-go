package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SharingLink struct {
    additionalData map[string]interface{};
    application *Identity;
    preventsDownload *bool;
    scope *string;
    type_escpaped *string;
    webHtml *string;
    webUrl *string;
}
func NewSharingLink()(*SharingLink) {
    m := &SharingLink{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SharingLink) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SharingLink) GetApplication()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
func (m *SharingLink) GetPreventsDownload()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preventsDownload
    }
}
func (m *SharingLink) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
func (m *SharingLink) GetType_escpaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *SharingLink) GetWebHtml()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webHtml
    }
}
func (m *SharingLink) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *SharingLink) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetApplication(val.(*Identity))
        return nil
    }
    res["preventsDownload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPreventsDownload(val)
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScope(val)
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
    res["webHtml"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebHtml(val)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *SharingLink) IsNil()(bool) {
    return m == nil
}
func (m *SharingLink) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preventsDownload", m.GetPreventsDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
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
        err := writer.WriteStringValue("webHtml", m.GetWebHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *SharingLink) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SharingLink) SetApplication(value *Identity)() {
    m.application = value
}
func (m *SharingLink) SetPreventsDownload(value *bool)() {
    m.preventsDownload = value
}
func (m *SharingLink) SetScope(value *string)() {
    m.scope = value
}
func (m *SharingLink) SetType_escpaped(value *string)() {
    m.type_escpaped = value
}
func (m *SharingLink) SetWebHtml(value *string)() {
    m.webHtml = value
}
func (m *SharingLink) SetWebUrl(value *string)() {
    m.webUrl = value
}
