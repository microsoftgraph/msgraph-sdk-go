package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Album struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Unique identifier of the [driveItem][] that is the cover of the album.
    coverImageItemId *string;
}
// Instantiates a new album and sets the default values.
func NewAlbum()(*Album) {
    m := &Album{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Album) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the coverImageItemId property value. Unique identifier of the [driveItem][] that is the cover of the album.
func (m *Album) GetCoverImageItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.coverImageItemId
    }
}
// The deserialization information for the current model
func (m *Album) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["coverImageItemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCoverImageItemId(val)
        return nil
    }
    return res
}
func (m *Album) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Album) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("coverImageItemId", m.GetCoverImageItemId())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Album) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the coverImageItemId property value. Unique identifier of the [driveItem][] that is the cover of the album.
// Parameters:
//  - value : Value to set for the coverImageItemId property.
func (m *Album) SetCoverImageItemId(value *string)() {
    m.coverImageItemId = value
}
