package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentTypeInfo 
type ContentTypeInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The id of the content type.
    id *string
    // The name of the content type.
    name *string
    // The OdataType property
    odataType *string
}
// NewContentTypeInfo instantiates a new contentTypeInfo and sets the default values.
func NewContentTypeInfo()(*ContentTypeInfo) {
    m := &ContentTypeInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.contentTypeInfo";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateContentTypeInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentTypeInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentTypeInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentTypeInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentTypeInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id of the content type.
func (m *ContentTypeInfo) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetName gets the name property value. The name of the content type.
func (m *ContentTypeInfo) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentTypeInfo) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// Serialize serializes information the current object
func (m *ContentTypeInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *ContentTypeInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. The id of the content type.
func (m *ContentTypeInfo) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetName sets the name property value. The name of the content type.
func (m *ContentTypeInfo) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentTypeInfo) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
