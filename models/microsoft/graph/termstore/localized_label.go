package termstore

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LocalizedLabel provides operations to manage the collection of drive entities.
type LocalizedLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the label is the default label.
    isDefault *bool;
    // The language tag for the label.
    languageTag *string;
    // The name of the label.
    name *string;
}
// NewLocalizedLabel instantiates a new localizedLabel and sets the default values.
func NewLocalizedLabel()(*LocalizedLabel) {
    m := &LocalizedLabel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLocalizedLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocalizedLabelFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewLocalizedLabel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocalizedLabel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocalizedLabel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates whether the label is the default label.
func (m *LocalizedLabel) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetLanguageTag gets the languageTag property value. The language tag for the label.
func (m *LocalizedLabel) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetName gets the name property value. The name of the label.
func (m *LocalizedLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *LocalizedLabel) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LocalizedLabel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
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
        err := writer.WriteStringValue("name", m.GetName())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocalizedLabel) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsDefault sets the isDefault property value. Indicates whether the label is the default label.
func (m *LocalizedLabel) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetLanguageTag sets the languageTag property value. The language tag for the label.
func (m *LocalizedLabel) SetLanguageTag(value *string)() {
    if m != nil {
        m.languageTag = value
    }
}
// SetName sets the name property value. The name of the label.
func (m *LocalizedLabel) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
