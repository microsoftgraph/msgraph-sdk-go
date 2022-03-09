package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookTableRow provides operations to manage the collection of drive entities.
type WorkbookTableRow struct {
    Entity
    // Returns the index number of the row within the rows collection of the table. Zero-indexed. Read-only.
    index *int32;
    // Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values Jsonable;
}
// NewWorkbookTableRow instantiates a new workbookTableRow and sets the default values.
func NewWorkbookTableRow()(*WorkbookTableRow) {
    m := &WorkbookTableRow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookTableRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookTableRowFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookTableRow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. Returns the index number of the row within the rows collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableRow) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// GetValues gets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableRow) GetValues()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *WorkbookTableRow) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetIndex sets the index property value. Returns the index number of the row within the rows collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableRow) SetIndex(value *int32)() {
    if m != nil {
        m.index = value
    }
}
// SetValues sets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableRow) SetValues(value Jsonable)() {
    if m != nil {
        m.values = value
    }
}
