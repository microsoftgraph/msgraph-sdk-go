package supportedtimezoneswithtimezonestandard

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SupportedTimeZonesWithTimeZoneStandard struct {
    additionalData map[string]interface{};
    alias *string;
    displayName *string;
}
func NewSupportedTimeZonesWithTimeZoneStandard()(*SupportedTimeZonesWithTimeZoneStandard) {
    m := &SupportedTimeZonesWithTimeZoneStandard{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SupportedTimeZonesWithTimeZoneStandard) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SupportedTimeZonesWithTimeZoneStandard) GetAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alias
    }
}
func (m *SupportedTimeZonesWithTimeZoneStandard) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *SupportedTimeZonesWithTimeZoneStandard) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *SupportedTimeZonesWithTimeZoneStandard) IsNil()(bool) {
    return m == nil
}
func (m *SupportedTimeZonesWithTimeZoneStandard) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *SupportedTimeZonesWithTimeZoneStandard) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SupportedTimeZonesWithTimeZoneStandard) SetAlias(value *string)() {
    m.alias = value
}
func (m *SupportedTimeZonesWithTimeZoneStandard) SetDisplayName(value *string)() {
    m.displayName = value
}
