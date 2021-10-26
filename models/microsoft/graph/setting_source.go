package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SettingSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Not yet documented
    displayName *string;
    // Not yet documented
    id *string;
    // Not yet documented. Possible values are: deviceConfiguration, deviceIntent.
    sourceType *SettingSourceType;
}
// Instantiates a new settingSource and sets the default values.
func NewSettingSource()(*SettingSource) {
    m := &SettingSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingSource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. Not yet documented
func (m *SettingSource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the id property value. Not yet documented
func (m *SettingSource) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the sourceType property value. Not yet documented. Possible values are: deviceConfiguration, deviceIntent.
func (m *SettingSource) GetSourceType()(*SettingSourceType) {
    if m == nil {
        return nil
    } else {
        return m.sourceType
    }
}
// The deserialization information for the current model
func (m *SettingSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["sourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSettingSourceType)
        if err != nil {
            return err
        }
        cast := val.(SettingSourceType)
        m.SetSourceType(&cast)
        return nil
    }
    return res
}
func (m *SettingSource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SettingSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetSourceType() != nil {
        cast := m.GetSourceType().String()
        err := writer.WriteStringValue("sourceType", &cast)
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
func (m *SettingSource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. Not yet documented
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SettingSource) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the id property value. Not yet documented
// Parameters:
//  - value : Value to set for the id property.
func (m *SettingSource) SetId(value *string)() {
    m.id = value
}
// Sets the sourceType property value. Not yet documented. Possible values are: deviceConfiguration, deviceIntent.
// Parameters:
//  - value : Value to set for the sourceType property.
func (m *SettingSource) SetSourceType(value *SettingSourceType)() {
    m.sourceType = value
}
