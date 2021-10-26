package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TextColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Whether to allow multiple lines of text.
    allowMultipleLines *bool;
    // Whether updates to this column should replace existing text, or append to it.
    appendChangesToExistingText *bool;
    // The size of the text box.
    linesForEditing *int32;
    // The maximum number of characters for the value.
    maxLength *int32;
    // The type of text being stored. Must be one of plain or richText
    textType *string;
}
// Instantiates a new textColumn and sets the default values.
func NewTextColumn()(*TextColumn) {
    m := &TextColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TextColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowMultipleLines property value. Whether to allow multiple lines of text.
func (m *TextColumn) GetAllowMultipleLines()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleLines
    }
}
// Gets the appendChangesToExistingText property value. Whether updates to this column should replace existing text, or append to it.
func (m *TextColumn) GetAppendChangesToExistingText()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appendChangesToExistingText
    }
}
// Gets the linesForEditing property value. The size of the text box.
func (m *TextColumn) GetLinesForEditing()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.linesForEditing
    }
}
// Gets the maxLength property value. The maximum number of characters for the value.
func (m *TextColumn) GetMaxLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxLength
    }
}
// Gets the textType property value. The type of text being stored. Must be one of plain or richText
func (m *TextColumn) GetTextType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.textType
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TextColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowMultipleLines property value. Whether to allow multiple lines of text.
// Parameters:
//  - value : Value to set for the allowMultipleLines property.
func (m *TextColumn) SetAllowMultipleLines(value *bool)() {
    m.allowMultipleLines = value
}
// Sets the appendChangesToExistingText property value. Whether updates to this column should replace existing text, or append to it.
// Parameters:
//  - value : Value to set for the appendChangesToExistingText property.
func (m *TextColumn) SetAppendChangesToExistingText(value *bool)() {
    m.appendChangesToExistingText = value
}
// Sets the linesForEditing property value. The size of the text box.
// Parameters:
//  - value : Value to set for the linesForEditing property.
func (m *TextColumn) SetLinesForEditing(value *int32)() {
    m.linesForEditing = value
}
// Sets the maxLength property value. The maximum number of characters for the value.
// Parameters:
//  - value : Value to set for the maxLength property.
func (m *TextColumn) SetMaxLength(value *int32)() {
    m.maxLength = value
}
// Sets the textType property value. The type of text being stored. Must be one of plain or richText
// Parameters:
//  - value : Value to set for the textType property.
func (m *TextColumn) SetTextType(value *string)() {
    m.textType = value
}
