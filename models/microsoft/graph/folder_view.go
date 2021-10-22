package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FolderView struct {
    additionalData map[string]interface{};
    sortBy *string;
    sortOrder *string;
    viewType *string;
}
func NewFolderView()(*FolderView) {
    m := &FolderView{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FolderView) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FolderView) GetSortBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortBy
    }
}
func (m *FolderView) GetSortOrder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortOrder
    }
}
func (m *FolderView) GetViewType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.viewType
    }
}
func (m *FolderView) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["sortBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSortBy(val)
        return nil
    }
    res["sortOrder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSortOrder(val)
        return nil
    }
    res["viewType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetViewType(val)
        return nil
    }
    return res
}
func (m *FolderView) IsNil()(bool) {
    return m == nil
}
func (m *FolderView) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sortBy", m.GetSortBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sortOrder", m.GetSortOrder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("viewType", m.GetViewType())
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
func (m *FolderView) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FolderView) SetSortBy(value *string)() {
    m.sortBy = value
}
func (m *FolderView) SetSortOrder(value *string)() {
    m.sortOrder = value
}
func (m *FolderView) SetViewType(value *string)() {
    m.viewType = value
}
