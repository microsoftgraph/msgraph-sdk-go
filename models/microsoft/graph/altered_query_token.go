package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AlteredQueryToken provides operations to call the query method.
type AlteredQueryToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Defines the length of a changed segment.
    length *int32;
    // Defines the offset of a changed segment.
    offset *int32;
    // Represents the corrected segment string.
    suggestion *string;
}
// NewAlteredQueryToken instantiates a new alteredQueryToken and sets the default values.
func NewAlteredQueryToken()(*AlteredQueryToken) {
    m := &AlteredQueryToken{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlteredQueryTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlteredQueryTokenFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAlteredQueryToken(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlteredQueryToken) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlteredQueryToken) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["length"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
        }
        return nil
    }
    res["offset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
        }
        return nil
    }
    res["suggestion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestion(val)
        }
        return nil
    }
    return res
}
// GetLength gets the length property value. Defines the length of a changed segment.
func (m *AlteredQueryToken) GetLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.length
    }
}
// GetOffset gets the offset property value. Defines the offset of a changed segment.
func (m *AlteredQueryToken) GetOffset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
// GetSuggestion gets the suggestion property value. Represents the corrected segment string.
func (m *AlteredQueryToken) GetSuggestion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.suggestion
    }
}
func (m *AlteredQueryToken) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AlteredQueryToken) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("length", m.GetLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("offset", m.GetOffset())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("suggestion", m.GetSuggestion())
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
func (m *AlteredQueryToken) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLength sets the length property value. Defines the length of a changed segment.
func (m *AlteredQueryToken) SetLength(value *int32)() {
    if m != nil {
        m.length = value
    }
}
// SetOffset sets the offset property value. Defines the offset of a changed segment.
func (m *AlteredQueryToken) SetOffset(value *int32)() {
    if m != nil {
        m.offset = value
    }
}
// SetSuggestion sets the suggestion property value. Represents the corrected segment string.
func (m *AlteredQueryToken) SetSuggestion(value *string)() {
    if m != nil {
        m.suggestion = value
    }
}
