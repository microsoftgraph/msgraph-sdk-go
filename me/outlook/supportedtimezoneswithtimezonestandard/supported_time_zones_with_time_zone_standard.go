package supportedtimezoneswithtimezonestandard

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SupportedTimeZonesWithTimeZoneStandard struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // An identifier for the time zone.
    alias *string;
    // A display string that represents the time zone.
    displayName *string;
}
// Instantiates a new supportedTimeZonesWithTimeZoneStandard and sets the default values.
func NewSupportedTimeZonesWithTimeZoneStandard()(*SupportedTimeZonesWithTimeZoneStandard) {
    m := &SupportedTimeZonesWithTimeZoneStandard{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SupportedTimeZonesWithTimeZoneStandard) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the alias property value. An identifier for the time zone.
func (m *SupportedTimeZonesWithTimeZoneStandard) GetAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alias
    }
}
// Gets the displayName property value. A display string that represents the time zone.
func (m *SupportedTimeZonesWithTimeZoneStandard) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// The deserialization information for the current model
func (m *SupportedTimeZonesWithTimeZoneStandard) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alias"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlias(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *SupportedTimeZonesWithTimeZoneStandard) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SupportedTimeZonesWithTimeZoneStandard) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the alias property value. An identifier for the time zone.
// Parameters:
//  - value : Value to set for the alias property.
func (m *SupportedTimeZonesWithTimeZoneStandard) SetAlias(value *string)() {
    m.alias = value
}
// Sets the displayName property value. A display string that represents the time zone.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SupportedTimeZonesWithTimeZoneStandard) SetDisplayName(value *string)() {
    m.displayName = value
}
