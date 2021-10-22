package supportedtimezones

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SupportedTimeZones struct {
    additionalData map[string]interface{};
    alias *string;
    displayName *string;
}
func NewSupportedTimeZones()(*SupportedTimeZones) {
    m := &SupportedTimeZones{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SupportedTimeZones) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SupportedTimeZones) GetAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alias
    }
}
func (m *SupportedTimeZones) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *SupportedTimeZones) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alias"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlias(val)
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
    return res
}
func (m *SupportedTimeZones) IsNil()(bool) {
    return m == nil
}
func (m *SupportedTimeZones) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alias", m.GetAlias())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SupportedTimeZones) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SupportedTimeZones) SetAlias(value *string)() {
    m.alias = value
}
func (m *SupportedTimeZones) SetDisplayName(value *string)() {
    m.displayName = value
}
