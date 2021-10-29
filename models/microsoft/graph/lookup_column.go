package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LookupColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether multiple values can be selected from the source.
    allowMultipleValues *bool;
    // Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
    allowUnlimitedLength *bool;
    // The name of the lookup source column.
    columnName *string;
    // The unique identifier of the lookup source list.
    listId *string;
    // If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here.
    primaryLookupColumnId *string;
}
// Instantiates a new lookupColumn and sets the default values.
func NewLookupColumn()(*LookupColumn) {
    m := &LookupColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LookupColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowMultipleValues property value. Indicates whether multiple values can be selected from the source.
func (m *LookupColumn) GetAllowMultipleValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleValues
    }
}
// Gets the allowUnlimitedLength property value. Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
func (m *LookupColumn) GetAllowUnlimitedLength()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUnlimitedLength
    }
}
// Gets the columnName property value. The name of the lookup source column.
func (m *LookupColumn) GetColumnName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.columnName
    }
}
// Gets the listId property value. The unique identifier of the lookup source list.
func (m *LookupColumn) GetListId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listId
    }
}
// Gets the primaryLookupColumnId property value. If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here.
func (m *LookupColumn) GetPrimaryLookupColumnId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.primaryLookupColumnId
    }
}
// The deserialization information for the current model
func (m *LookupColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowMultipleValues(val)
        return nil
    }
    res["allowUnlimitedLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowUnlimitedLength(val)
        return nil
    }
    res["columnName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColumnName(val)
        return nil
    }
    res["listId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetListId(val)
        return nil
    }
    res["primaryLookupColumnId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrimaryLookupColumnId(val)
        return nil
    }
    return res
}
func (m *LookupColumn) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LookupColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleValues", m.GetAllowMultipleValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowUnlimitedLength", m.GetAllowUnlimitedLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("columnName", m.GetColumnName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("listId", m.GetListId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("primaryLookupColumnId", m.GetPrimaryLookupColumnId())
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
func (m *LookupColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowMultipleValues property value. Indicates whether multiple values can be selected from the source.
// Parameters:
//  - value : Value to set for the allowMultipleValues property.
func (m *LookupColumn) SetAllowMultipleValues(value *bool)() {
    m.allowMultipleValues = value
}
// Sets the allowUnlimitedLength property value. Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
// Parameters:
//  - value : Value to set for the allowUnlimitedLength property.
func (m *LookupColumn) SetAllowUnlimitedLength(value *bool)() {
    m.allowUnlimitedLength = value
}
// Sets the columnName property value. The name of the lookup source column.
// Parameters:
//  - value : Value to set for the columnName property.
func (m *LookupColumn) SetColumnName(value *string)() {
    m.columnName = value
}
// Sets the listId property value. The unique identifier of the lookup source list.
// Parameters:
//  - value : Value to set for the listId property.
func (m *LookupColumn) SetListId(value *string)() {
    m.listId = value
}
// Sets the primaryLookupColumnId property value. If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here.
// Parameters:
//  - value : Value to set for the primaryLookupColumnId property.
func (m *LookupColumn) SetPrimaryLookupColumnId(value *string)() {
    m.primaryLookupColumnId = value
}
