package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookSortField struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents whether the sorting is done in an ascending fashion.
    ascending *bool;
    // Represents the color that is the target of the condition if the sorting is on font or cell color.
    color *string;
    // Represents additional sorting options for this field. The possible values are: Normal, TextAsNumber.
    dataOption *string;
    // Represents the icon that is the target of the condition if the sorting is on the cell's icon.
    icon *WorkbookIcon;
    // Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset from the first column (or row).
    key *int32;
    // Represents the type of sorting of this condition. The possible values are: Value, CellColor, FontColor, Icon.
    sortOn *string;
}
// Instantiates a new workbookSortField and sets the default values.
func NewWorkbookSortField()(*WorkbookSortField) {
    m := &WorkbookSortField{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookSortField) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the ascending property value. Represents whether the sorting is done in an ascending fashion.
func (m *WorkbookSortField) GetAscending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ascending
    }
}
// Gets the color property value. Represents the color that is the target of the condition if the sorting is on font or cell color.
func (m *WorkbookSortField) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// Gets the dataOption property value. Represents additional sorting options for this field. The possible values are: Normal, TextAsNumber.
func (m *WorkbookSortField) GetDataOption()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataOption
    }
}
// Gets the icon property value. Represents the icon that is the target of the condition if the sorting is on the cell's icon.
func (m *WorkbookSortField) GetIcon()(*WorkbookIcon) {
    if m == nil {
        return nil
    } else {
        return m.icon
    }
}
// Gets the key property value. Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset from the first column (or row).
func (m *WorkbookSortField) GetKey()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// Gets the sortOn property value. Represents the type of sorting of this condition. The possible values are: Value, CellColor, FontColor, Icon.
func (m *WorkbookSortField) GetSortOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortOn
    }
}
// The deserialization information for the current model
func (m *WorkbookSortField) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ascending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAscending(val)
        return nil
    }
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColor(val)
        return nil
    }
    res["dataOption"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDataOption(val)
        return nil
    }
    res["icon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookIcon() })
        if err != nil {
            return err
        }
        m.SetIcon(val.(*WorkbookIcon))
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["sortOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSortOn(val)
        return nil
    }
    return res
}
func (m *WorkbookSortField) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookSortField) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("ascending", m.GetAscending())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataOption", m.GetDataOption())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("icon", m.GetIcon())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sortOn", m.GetSortOn())
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
func (m *WorkbookSortField) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the ascending property value. Represents whether the sorting is done in an ascending fashion.
// Parameters:
//  - value : Value to set for the ascending property.
func (m *WorkbookSortField) SetAscending(value *bool)() {
    m.ascending = value
}
// Sets the color property value. Represents the color that is the target of the condition if the sorting is on font or cell color.
// Parameters:
//  - value : Value to set for the color property.
func (m *WorkbookSortField) SetColor(value *string)() {
    m.color = value
}
// Sets the dataOption property value. Represents additional sorting options for this field. The possible values are: Normal, TextAsNumber.
// Parameters:
//  - value : Value to set for the dataOption property.
func (m *WorkbookSortField) SetDataOption(value *string)() {
    m.dataOption = value
}
// Sets the icon property value. Represents the icon that is the target of the condition if the sorting is on the cell's icon.
// Parameters:
//  - value : Value to set for the icon property.
func (m *WorkbookSortField) SetIcon(value *WorkbookIcon)() {
    m.icon = value
}
// Sets the key property value. Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset from the first column (or row).
// Parameters:
//  - value : Value to set for the key property.
func (m *WorkbookSortField) SetKey(value *int32)() {
    m.key = value
}
// Sets the sortOn property value. Represents the type of sorting of this condition. The possible values are: Value, CellColor, FontColor, Icon.
// Parameters:
//  - value : Value to set for the sortOn property.
func (m *WorkbookSortField) SetSortOn(value *string)() {
    m.sortOn = value
}
