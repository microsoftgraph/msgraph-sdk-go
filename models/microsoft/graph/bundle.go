package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Bundle provides operations to manage the collection of drive entities.
type Bundle struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If the bundle is an [album][], then the album property is included
    album Albumable;
    // Number of children contained immediately within this container.
    childCount *int32;
}
// NewBundle instantiates a new bundle and sets the default values.
func NewBundle()(*Bundle) {
    m := &Bundle{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBundleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBundleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBundle(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Bundle) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlbum gets the album property value. If the bundle is an [album][], then the album property is included
func (m *Bundle) GetAlbum()(Albumable) {
    if m == nil {
        return nil
    } else {
        return m.album
    }
}
// GetChildCount gets the childCount property value. Number of children contained immediately within this container.
func (m *Bundle) GetChildCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.childCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Bundle) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["album"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlbumFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlbum(val.(Albumable))
        }
        return nil
    }
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
    return res
}
func (m *Bundle) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Bundle) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("album", m.GetAlbum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("childCount", m.GetChildCount())
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
func (m *Bundle) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlbum sets the album property value. If the bundle is an [album][], then the album property is included
func (m *Bundle) SetAlbum(value Albumable)() {
    if m != nil {
        m.album = value
    }
}
// SetChildCount sets the childCount property value. Number of children contained immediately within this container.
func (m *Bundle) SetChildCount(value *int32)() {
    if m != nil {
        m.childCount = value
    }
}
