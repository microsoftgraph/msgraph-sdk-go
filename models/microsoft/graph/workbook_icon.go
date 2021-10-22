package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookIcon struct {
    additionalData map[string]interface{};
    index *int32;
    set *string;
}
func NewWorkbookIcon()(*WorkbookIcon) {
    m := &WorkbookIcon{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WorkbookIcon) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WorkbookIcon) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
func (m *WorkbookIcon) GetSet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.set
    }
}
func (m *WorkbookIcon) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIndex(val)
        return nil
    }
    res["set"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSet(val)
        return nil
    }
    return res
}
func (m *WorkbookIcon) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookIcon) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("set", m.GetSet())
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
func (m *WorkbookIcon) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WorkbookIcon) SetIndex(value *int32)() {
    m.index = value
}
func (m *WorkbookIcon) SetSet(value *string)() {
    m.set = value
}
