package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TextColumn struct {
    additionalData map[string]interface{};
    allowMultipleLines *bool;
    appendChangesToExistingText *bool;
    linesForEditing *int32;
    maxLength *int32;
    textType *string;
}
func NewTextColumn()(*TextColumn) {
    m := &TextColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TextColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TextColumn) GetAllowMultipleLines()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleLines
    }
}
func (m *TextColumn) GetAppendChangesToExistingText()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appendChangesToExistingText
    }
}
func (m *TextColumn) GetLinesForEditing()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.linesForEditing
    }
}
func (m *TextColumn) GetMaxLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxLength
    }
}
func (m *TextColumn) GetTextType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.textType
    }
}
func (m *TextColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowMultipleLines(val)
        return nil
    }
    res["appendChangesToExistingText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAppendChangesToExistingText(val)
        return nil
    }
    res["linesForEditing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLinesForEditing(val)
        return nil
    }
    res["maxLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaxLength(val)
        return nil
    }
    res["textType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTextType(val)
        return nil
    }
    return res
}
func (m *TextColumn) IsNil()(bool) {
    return m == nil
}
func (m *TextColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleLines", m.GetAllowMultipleLines())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("appendChangesToExistingText", m.GetAppendChangesToExistingText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("linesForEditing", m.GetLinesForEditing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxLength", m.GetMaxLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("textType", m.GetTextType())
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
func (m *TextColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TextColumn) SetAllowMultipleLines(value *bool)() {
    m.allowMultipleLines = value
}
func (m *TextColumn) SetAppendChangesToExistingText(value *bool)() {
    m.appendChangesToExistingText = value
}
func (m *TextColumn) SetLinesForEditing(value *int32)() {
    m.linesForEditing = value
}
func (m *TextColumn) SetMaxLength(value *int32)() {
    m.maxLength = value
}
func (m *TextColumn) SetTextType(value *string)() {
    m.textType = value
}
