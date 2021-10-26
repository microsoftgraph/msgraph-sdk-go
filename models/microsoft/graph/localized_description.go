package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LocalizedDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The description in the localized language.
    description *string;
    // The language tag for the label.
    languageTag *string;
}
// Instantiates a new localizedDescription and sets the default values.
func NewLocalizedDescription()(*LocalizedDescription) {
    m := &LocalizedDescription{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocalizedDescription) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. The description in the localized language.
func (m *LocalizedDescription) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the languageTag property value. The language tag for the label.
func (m *LocalizedDescription) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// The deserialization information for the current model
func (m *LocalizedDescription) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["languageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguageTag(val)
        return nil
    }
    return res
}
func (m *LocalizedDescription) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LocalizedDescription) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
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
func (m *LocalizedDescription) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. The description in the localized language.
// Parameters:
//  - value : Value to set for the description property.
func (m *LocalizedDescription) SetDescription(value *string)() {
    m.description = value
}
// Sets the languageTag property value. The language tag for the label.
// Parameters:
//  - value : Value to set for the languageTag property.
func (m *LocalizedDescription) SetLanguageTag(value *string)() {
    m.languageTag = value
}
