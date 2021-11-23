package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ThumbnailSet 
type ThumbnailSet struct {
    Entity
    // A 1920x1920 scaled thumbnail.
    large *Thumbnail;
    // A 176x176 scaled thumbnail.
    medium *Thumbnail;
    // A 48x48 cropped thumbnail.
    small *Thumbnail;
    // A custom thumbnail image or the original image used to generate other thumbnails.
    source *Thumbnail;
}
// NewThumbnailSet instantiates a new thumbnailSet and sets the default values.
func NewThumbnailSet()(*ThumbnailSet) {
    m := &ThumbnailSet{
        Entity: *NewEntity(),
    }
    return m
}
// GetLarge gets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) GetLarge()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.large
    }
}
// GetMedium gets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) GetMedium()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.medium
    }
}
// GetSmall gets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) GetSmall()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.small
    }
}
// GetSource gets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) GetSource()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThumbnailSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["large"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLarge(val.(*Thumbnail))
        }
        return nil
    }
    res["medium"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedium(val.(*Thumbnail))
        }
        return nil
    }
    res["small"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmall(val.(*Thumbnail))
        }
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*Thumbnail))
        }
        return nil
    }
    return res
}
func (m *ThumbnailSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ThumbnailSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("large", m.GetLarge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("medium", m.GetMedium())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("small", m.GetSmall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLarge sets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) SetLarge(value *Thumbnail)() {
    m.large = value
}
// SetMedium sets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) SetMedium(value *Thumbnail)() {
    m.medium = value
}
// SetSmall sets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) SetSmall(value *Thumbnail)() {
    m.small = value
}
// SetSource sets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) SetSource(value *Thumbnail)() {
    m.source = value
}
