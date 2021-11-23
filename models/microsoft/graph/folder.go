package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// folder 
type Folder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of children contained immediately within this container.
    childCount *int32;
    // A collection of properties defining the recommended view for the folder.
    view *FolderView;
}
// NewFolder instantiates a new folder and sets the default values.
func NewFolder()(*Folder) {
    m := &Folder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Folder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChildCount gets the childCount property value. Number of children contained immediately within this container.
func (m *Folder) GetChildCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.childCount
    }
}
// GetView gets the view property value. A collection of properties defining the recommended view for the folder.
func (m *Folder) GetView()(*FolderView) {
    if m == nil {
        return nil
    } else {
        return m.view
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Folder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["childCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildCount(val)
        }
        return nil
    }
    res["view"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFolderView() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetView(val.(*FolderView))
        }
        return nil
    }
    return res
}
func (m *Folder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Folder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("childCount", m.GetChildCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("view", m.GetView())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Folder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetChildCount sets the childCount property value. Number of children contained immediately within this container.
func (m *Folder) SetChildCount(value *int32)() {
    m.childCount = value
}
// SetView sets the view property value. A collection of properties defining the recommended view for the folder.
func (m *Folder) SetView(value *FolderView)() {
    m.view = value
}
