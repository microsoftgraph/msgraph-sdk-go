package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookTableSort 
type WorkbookTableSort struct {
    Entity
    // Represents the current conditions used to last sort the table. Read-only.
    fields []WorkbookSortField;
    // Represents whether the casing impacted the last sort of the table. Read-only.
    matchCase *bool;
    // Represents Chinese character ordering method last used to sort the table. Possible values are: PinYin, StrokeCount. Read-only.
    method *string;
}
// NewWorkbookTableSort instantiates a new workbookTableSort and sets the default values.
func NewWorkbookTableSort()(*WorkbookTableSort) {
    m := &WorkbookTableSort{
        Entity: *NewEntity(),
    }
    return m
}
// GetFields gets the fields property value. Represents the current conditions used to last sort the table. Read-only.
func (m *WorkbookTableSort) GetFields()([]WorkbookSortField) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// GetMatchCase gets the matchCase property value. Represents whether the casing impacted the last sort of the table. Read-only.
func (m *WorkbookTableSort) GetMatchCase()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.matchCase
    }
}
// GetMethod gets the method property value. Represents Chinese character ordering method last used to sort the table. Possible values are: PinYin, StrokeCount. Read-only.
func (m *WorkbookTableSort) GetMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.method
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookTableSort) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookSortField() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookSortField, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookSortField))
            }
            m.SetFields(res)
        }
        return nil
    }
    res["matchCase"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchCase(val)
        }
        return nil
    }
    res["method"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethod(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookTableSort) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookTableSort) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFields() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("fields", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("matchCase", m.GetMatchCase())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("method", m.GetMethod())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFields sets the fields property value. Represents the current conditions used to last sort the table. Read-only.
func (m *WorkbookTableSort) SetFields(value []WorkbookSortField)() {
    if m != nil {
        m.fields = value
    }
}
// SetMatchCase sets the matchCase property value. Represents whether the casing impacted the last sort of the table. Read-only.
func (m *WorkbookTableSort) SetMatchCase(value *bool)() {
    if m != nil {
        m.matchCase = value
    }
}
// SetMethod sets the method property value. Represents Chinese character ordering method last used to sort the table. Possible values are: PinYin, StrokeCount. Read-only.
func (m *WorkbookTableSort) SetMethod(value *string)() {
    if m != nil {
        m.method = value
    }
}
