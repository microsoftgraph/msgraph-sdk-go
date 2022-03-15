package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DriveItemVersion provides operations to manage the collection of drive entities.
type DriveItemVersion struct {
    BaseItemVersion
    // The content stream for this version of the item.
    content []byte;
    // Indicates the size of the content stream for this version of the item.
    size *int32;
}
// NewDriveItemVersion instantiates a new driveItemVersion and sets the default values.
func NewDriveItemVersion()(*DriveItemVersion) {
    m := &DriveItemVersion{
        BaseItemVersion: *NewBaseItemVersion(),
    }
    return m
}
// CreateDriveItemVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveItemVersionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDriveItemVersion(), nil
}
// GetContent gets the content property value. The content stream for this version of the item.
func (m *DriveItemVersion) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItemVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItemVersion.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.Get[]byteValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetSize gets the size property value. Indicates the size of the content stream for this version of the item.
func (m *DriveItemVersion) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *DriveItemVersion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DriveItemVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BaseItemVersion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content stream for this version of the item.
func (m *DriveItemVersion) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetSize sets the size property value. Indicates the size of the content stream for this version of the item.
func (m *DriveItemVersion) SetSize(value *int32)() {
    if m != nil {
        m.size = value
    }
}
