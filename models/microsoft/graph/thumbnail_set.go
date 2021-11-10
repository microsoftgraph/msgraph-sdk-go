package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new thumbnailSet and sets the default values.
func NewThumbnailSet()(*ThumbnailSet) {
    m := &ThumbnailSet{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) GetLarge()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.large
    }
}
// Gets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) GetMedium()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.medium
    }
}
// Gets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) GetSmall()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.small
    }
}
// Gets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) GetSource()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the large property value. A 1920x1920 scaled thumbnail.
// Parameters:
//  - value : Value to set for the large property.
func (m *ThumbnailSet) SetLarge(value *Thumbnail)() {
    m.large = value
}
// Sets the medium property value. A 176x176 scaled thumbnail.
// Parameters:
//  - value : Value to set for the medium property.
func (m *ThumbnailSet) SetMedium(value *Thumbnail)() {
    m.medium = value
}
// Sets the small property value. A 48x48 cropped thumbnail.
// Parameters:
//  - value : Value to set for the small property.
func (m *ThumbnailSet) SetSmall(value *Thumbnail)() {
    m.small = value
}
// Sets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
// Parameters:
//  - value : Value to set for the source property.
func (m *ThumbnailSet) SetSource(value *Thumbnail)() {
    m.source = value
}
