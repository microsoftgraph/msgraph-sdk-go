package supportedlanguages

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SupportedLanguages struct {
    additionalData map[string]interface{};
    displayName *string;
    locale *string;
}
func NewSupportedLanguages()(*SupportedLanguages) {
    m := &SupportedLanguages{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SupportedLanguages) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SupportedLanguages) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *SupportedLanguages) GetLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locale
    }
}
func (m *SupportedLanguages) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["locale"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocale(val)
        return nil
    }
    return res
}
func (m *SupportedLanguages) IsNil()(bool) {
    return m == nil
}
func (m *SupportedLanguages) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
func (m *SupportedLanguages) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SupportedLanguages) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *SupportedLanguages) SetLocale(value *string)() {
    m.locale = value
}
