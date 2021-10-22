package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LookupColumn struct {
    additionalData map[string]interface{};
    allowMultipleValues *bool;
    allowUnlimitedLength *bool;
    columnName *string;
    listId *string;
    primaryLookupColumnId *string;
}
func NewLookupColumn()(*LookupColumn) {
    m := &LookupColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LookupColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LookupColumn) GetAllowMultipleValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleValues
    }
}
func (m *LookupColumn) GetAllowUnlimitedLength()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUnlimitedLength
    }
}
func (m *LookupColumn) GetColumnName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.columnName
    }
}
func (m *LookupColumn) GetListId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listId
    }
}
func (m *LookupColumn) GetPrimaryLookupColumnId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.primaryLookupColumnId
    }
}
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
func (m *LookupColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LookupColumn) SetAllowMultipleValues(value *bool)() {
    m.allowMultipleValues = value
}
func (m *LookupColumn) SetAllowUnlimitedLength(value *bool)() {
    m.allowUnlimitedLength = value
}
func (m *LookupColumn) SetColumnName(value *string)() {
    m.columnName = value
}
func (m *LookupColumn) SetListId(value *string)() {
    m.listId = value
}
func (m *LookupColumn) SetPrimaryLookupColumnId(value *string)() {
    m.primaryLookupColumnId = value
}
