package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookTableRow struct {
    Entity
    // Returns the index number of the row within the rows collection of the table. Zero-indexed. Read-only.
    index *int32;
    // Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values *Json;
}
// Instantiates a new workbookTableRow and sets the default values.
func NewWorkbookTableRow()(*WorkbookTableRow) {
    m := &WorkbookTableRow{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the index property value. Returns the index number of the row within the rows collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableRow) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// Gets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableRow) GetValues()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// The deserialization information for the current model
func (m *WorkbookTableRow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(*Json))
        }
        return nil
    }
    return res
}
func (m *WorkbookTableRow) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookTableRow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the index property value. Returns the index number of the row within the rows collection of the table. Zero-indexed. Read-only.
// Parameters:
//  - value : Value to set for the index property.
func (m *WorkbookTableRow) SetIndex(value *int32)() {
    m.index = value
}
// Sets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
// Parameters:
//  - value : Value to set for the values property.
func (m *WorkbookTableRow) SetValues(value *Json)() {
    m.values = value
}
