package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// displayNameLocalization 
type DisplayNameLocalization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field.
    displayName *string;
    // Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
    languageTag *string;
}
// NewDisplayNameLocalization instantiates a new displayNameLocalization and sets the default values.
func NewDisplayNameLocalization()(*DisplayNameLocalization) {
    m := &DisplayNameLocalization{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DisplayNameLocalization) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field.
func (m *DisplayNameLocalization) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLanguageTag gets the languageTag property value. Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
func (m *DisplayNameLocalization) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DisplayNameLocalization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["languageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    return res
}
func (m *DisplayNameLocalization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DisplayNameLocalization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DisplayNameLocalization) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field.
func (m *DisplayNameLocalization) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLanguageTag sets the languageTag property value. Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
func (m *DisplayNameLocalization) SetLanguageTag(value *string)() {
    m.languageTag = value
}
