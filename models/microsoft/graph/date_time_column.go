package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DateTimeColumn struct {
    additionalData map[string]interface{};
    displayAs *string;
    format *string;
}
func NewDateTimeColumn()(*DateTimeColumn) {
    m := &DateTimeColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DateTimeColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DateTimeColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
func (m *DateTimeColumn) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *DateTimeColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayAs(val)
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFormat(val)
        return nil
    }
    return res
}
func (m *DateTimeColumn) IsNil()(bool) {
    return m == nil
}
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
func (m *DateTimeColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DateTimeColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
func (m *DateTimeColumn) SetFormat(value *string)() {
    m.format = value
}
