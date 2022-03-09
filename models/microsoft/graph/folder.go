package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Folder provides operations to manage the collection of drive entities.
type Folder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of children contained immediately within this container.
    childCount *int32;
    // A collection of properties defining the recommended view for the folder.
    view FolderViewable;
}
// NewFolder instantiates a new folder and sets the default values.
func NewFolder()(*Folder) {
    m := &Folder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFolderFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFolder(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
        val, err := n.GetObjectValue(CreateFolderViewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetView(val.(FolderViewable))
        }
        return nil
    }
    return res
}
// GetView gets the view property value. A collection of properties defining the recommended view for the folder.
func (m *Folder) GetView()(FolderViewable) {
    if m == nil {
        return nil
    } else {
        return m.view
    }
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Folder) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChildCount sets the childCount property value. Number of children contained immediately within this container.
func (m *Folder) SetChildCount(value *int32)() {
    if m != nil {
        m.childCount = value
    }
}
// SetView sets the view property value. A collection of properties defining the recommended view for the folder.
func (m *Folder) SetView(value FolderViewable)() {
    if m != nil {
        m.view = value
    }
}
