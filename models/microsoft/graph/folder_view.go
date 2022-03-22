package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FolderView 
type FolderView struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The method by which the folder should be sorted.
    sortBy *string;
    // If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
    sortOrder *string;
    // The type of view that should be used to represent the folder.
    viewType *string;
}
// NewFolderView instantiates a new folderView and sets the default values.
func NewFolderView()(*FolderView) {
    m := &FolderView{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFolderViewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFolderViewFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFolderView(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FolderView) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FolderView) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["sortBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortBy(val)
        }
        return nil
    }
    res["sortOrder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortOrder(val)
        }
        return nil
    }
    res["viewType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewType(val)
        }
        return nil
    }
    return res
}
// GetSortBy gets the sortBy property value. The method by which the folder should be sorted.
func (m *FolderView) GetSortBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortBy
    }
}
// GetSortOrder gets the sortOrder property value. If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
func (m *FolderView) GetSortOrder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortOrder
    }
}
// GetViewType gets the viewType property value. The type of view that should be used to represent the folder.
func (m *FolderView) GetViewType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.viewType
    }
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FolderView) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSortBy sets the sortBy property value. The method by which the folder should be sorted.
func (m *FolderView) SetSortBy(value *string)() {
    if m != nil {
        m.sortBy = value
    }
}
// SetSortOrder sets the sortOrder property value. If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
func (m *FolderView) SetSortOrder(value *string)() {
    if m != nil {
        m.sortOrder = value
    }
}
// SetViewType sets the viewType property value. The type of view that should be used to represent the folder.
func (m *FolderView) SetViewType(value *string)() {
    if m != nil {
        m.viewType = value
    }
}
