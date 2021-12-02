package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SettingSource 
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
// NewSettingSource instantiates a new settingSource and sets the default values.
func NewSettingSource()(*SettingSource) {
    m := &SettingSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingSource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Not yet documented
func (m *SettingSource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetId gets the id property value. Not yet documented
func (m *SettingSource) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetSourceType gets the sourceType property value. Not yet documented. Possible values are: deviceConfiguration, deviceIntent.
func (m *SettingSource) GetSourceType()(*SettingSourceType) {
    if m == nil {
        return nil
    } else {
        return m.sourceType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["sourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSettingSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SettingSourceType)
            m.SetSourceType(&cast)
        }
        return nil
    }
    return res
}
func (m *SettingSource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingSource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Not yet documented
func (m *SettingSource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. Not yet documented
func (m *SettingSource) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetSourceType sets the sourceType property value. Not yet documented. Possible values are: deviceConfiguration, deviceIntent.
func (m *SettingSource) SetSourceType(value *SettingSourceType)() {
    if m != nil {
        m.sourceType = value
    }
}
