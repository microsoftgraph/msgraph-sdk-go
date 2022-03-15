package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookSortField provides operations to manage the collection of drive entities.
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
    icon WorkbookIconable;
    // Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset from the first column (or row).
    key *int32;
    // Represents the type of sorting of this condition. The possible values are: Value, CellColor, FontColor, Icon.
    sortOn *string;
}
// NewWorkbookSortField instantiates a new workbookSortField and sets the default values.
func NewWorkbookSortField()(*WorkbookSortField) {
    m := &WorkbookSortField{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbookSortFieldFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookSortFieldFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookSortField(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookSortField) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAscending gets the ascending property value. Represents whether the sorting is done in an ascending fashion.
func (m *WorkbookSortField) GetAscending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ascending
    }
}
// GetColor gets the color property value. Represents the color that is the target of the condition if the sorting is on font or cell color.
func (m *WorkbookSortField) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetDataOption gets the dataOption property value. Represents additional sorting options for this field. The possible values are: Normal, TextAsNumber.
func (m *WorkbookSortField) GetDataOption()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataOption
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookSortField) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ascending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAscending(val)
        }
        return nil
    }
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["dataOption"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataOption(val)
        }
        return nil
    }
    res["icon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookIconFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIcon(val.(WorkbookIconable))
        }
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["sortOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortOn(val)
        }
        return nil
    }
    return res
}
// GetIcon gets the icon property value. Represents the icon that is the target of the condition if the sorting is on the cell's icon.
func (m *WorkbookSortField) GetIcon()(WorkbookIconable) {
    if m == nil {
        return nil
    } else {
        return m.icon
    }
}
// GetKey gets the key property value. Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset from the first column (or row).
func (m *WorkbookSortField) GetKey()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetSortOn gets the sortOn property value. Represents the type of sorting of this condition. The possible values are: Value, CellColor, FontColor, Icon.
func (m *WorkbookSortField) GetSortOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortOn
    }
}
func (m *WorkbookSortField) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookSortField) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAscending sets the ascending property value. Represents whether the sorting is done in an ascending fashion.
func (m *WorkbookSortField) SetAscending(value *bool)() {
    if m != nil {
        m.ascending = value
    }
}
// SetColor sets the color property value. Represents the color that is the target of the condition if the sorting is on font or cell color.
func (m *WorkbookSortField) SetColor(value *string)() {
    if m != nil {
        m.color = value
    }
}
// SetDataOption sets the dataOption property value. Represents additional sorting options for this field. The possible values are: Normal, TextAsNumber.
func (m *WorkbookSortField) SetDataOption(value *string)() {
    if m != nil {
        m.dataOption = value
    }
}
// SetIcon sets the icon property value. Represents the icon that is the target of the condition if the sorting is on the cell's icon.
func (m *WorkbookSortField) SetIcon(value WorkbookIconable)() {
    if m != nil {
        m.icon = value
    }
}
// SetKey sets the key property value. Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset from the first column (or row).
func (m *WorkbookSortField) SetKey(value *int32)() {
    if m != nil {
        m.key = value
    }
}
// SetSortOn sets the sortOn property value. Represents the type of sorting of this condition. The possible values are: Value, CellColor, FontColor, Icon.
func (m *WorkbookSortField) SetSortOn(value *string)() {
    if m != nil {
        m.sortOn = value
    }
}
