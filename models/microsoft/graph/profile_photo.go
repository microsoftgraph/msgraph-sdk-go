package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProfilePhoto provides operations to manage the drive singleton.
type ProfilePhoto struct {
    Entity
    // The height of the photo. Read-only.
    height *int32;
    // The width of the photo. Read-only.
    width *int32;
}
// NewProfilePhoto instantiates a new profilePhoto and sets the default values.
func NewProfilePhoto()(*ProfilePhoto) {
    m := &ProfilePhoto{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProfilePhotoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProfilePhotoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewProfilePhoto(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProfilePhoto) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetHeight gets the height property value. The height of the photo. Read-only.
func (m *ProfilePhoto) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// GetWidth gets the width property value. The width of the photo. Read-only.
func (m *ProfilePhoto) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
func (m *ProfilePhoto) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProfilePhoto) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHeight sets the height property value. The height of the photo. Read-only.
func (m *ProfilePhoto) SetHeight(value *int32)() {
    if m != nil {
        m.height = value
    }
}
// SetWidth sets the width property value. The width of the photo. Read-only.
func (m *ProfilePhoto) SetWidth(value *int32)() {
    if m != nil {
        m.width = value
    }
}
