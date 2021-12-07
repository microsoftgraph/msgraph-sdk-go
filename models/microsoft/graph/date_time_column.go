package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DateTimeColumn 
type DateTimeColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default.
    displayAs *string;
    // Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime
    format *string;
}
// NewDateTimeColumn instantiates a new dateTimeColumn and sets the default values.
func NewDateTimeColumn()(*DateTimeColumn) {
    m := &DateTimeColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DateTimeColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayAs gets the displayAs property value. How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default.
func (m *DateTimeColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
// GetFormat gets the format property value. Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime
func (m *DateTimeColumn) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DateTimeColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAs(val)
        }
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    return res
}
func (m *DateTimeColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DateTimeColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("format", m.GetFormat())
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
func (m *DateTimeColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayAs sets the displayAs property value. How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default.
func (m *DateTimeColumn) SetDisplayAs(value *string)() {
    if m != nil {
        m.displayAs = value
    }
}
// SetFormat sets the format property value. Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime
func (m *DateTimeColumn) SetFormat(value *string)() {
    if m != nil {
        m.format = value
    }
}
