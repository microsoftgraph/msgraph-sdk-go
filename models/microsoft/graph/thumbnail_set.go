package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ThumbnailSet struct {
    Entity
    large *Thumbnail;
    medium *Thumbnail;
    small *Thumbnail;
    source *Thumbnail;
}
func NewThumbnailSet()(*ThumbnailSet) {
    m := &ThumbnailSet{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ThumbnailSet) GetLarge()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.large
    }
}
func (m *ThumbnailSet) GetMedium()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.medium
    }
}
func (m *ThumbnailSet) GetSmall()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.small
    }
}
func (m *ThumbnailSet) GetSource()(*Thumbnail) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
func (m *ThumbnailSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["large"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        m.SetLarge(val.(*Thumbnail))
        return nil
    }
    res["medium"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        m.SetMedium(val.(*Thumbnail))
        return nil
    }
    res["small"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        m.SetSmall(val.(*Thumbnail))
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnail() })
        if err != nil {
            return err
        }
        m.SetSource(val.(*Thumbnail))
        return nil
    }
    return res
}
func (m *ThumbnailSet) IsNil()(bool) {
    return m == nil
}
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
func (m *ThumbnailSet) SetLarge(value *Thumbnail)() {
    m.large = value
}
func (m *ThumbnailSet) SetMedium(value *Thumbnail)() {
    m.medium = value
}
func (m *ThumbnailSet) SetSmall(value *Thumbnail)() {
    m.small = value
}
func (m *ThumbnailSet) SetSource(value *Thumbnail)() {
    m.source = value
}
