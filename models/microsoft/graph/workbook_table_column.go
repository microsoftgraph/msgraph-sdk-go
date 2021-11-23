package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workbookTableColumn 
type WorkbookTableColumn struct {
    Entity
    // Retrieve the filter applied to the column. Read-only.
    filter *WorkbookFilter;
    // Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
    index *int32;
    // Returns the name of the table column.
    name *string;
    // Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values *Json;
}
// NewWorkbookTableColumn instantiates a new workbookTableColumn and sets the default values.
func NewWorkbookTableColumn()(*WorkbookTableColumn) {
    m := &WorkbookTableColumn{
        Entity: *NewEntity(),
    }
    return m
}
// GetFilter gets the filter property value. Retrieve the filter applied to the column. Read-only.
func (m *WorkbookTableColumn) GetFilter()(*WorkbookFilter) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// GetIndex gets the index property value. Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableColumn) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// GetName gets the name property value. Returns the name of the table column.
func (m *WorkbookTableColumn) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetValues gets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableColumn) GetValues()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookTableColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFilter() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val.(*WorkbookFilter))
        }
        return nil
    }
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
func (m *WorkbookTableColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookTableColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
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
// SetFilter sets the filter property value. Retrieve the filter applied to the column. Read-only.
func (m *WorkbookTableColumn) SetFilter(value *WorkbookFilter)() {
    m.filter = value
}
// SetIndex sets the index property value. Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableColumn) SetIndex(value *int32)() {
    m.index = value
}
// SetName sets the name property value. Returns the name of the table column.
func (m *WorkbookTableColumn) SetName(value *string)() {
    m.name = value
}
// SetValues sets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableColumn) SetValues(value *Json)() {
    m.values = value
}
