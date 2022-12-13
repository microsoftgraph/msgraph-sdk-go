package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DrivesItemItemsItemPreviewPostRequestBody provides operations to call the preview method.
type DrivesItemItemsItemPreviewPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The page property
    page *string
    // The zoom property
    zoom *float64
}
// NewDrivesItemItemsItemPreviewPostRequestBody instantiates a new DrivesItemItemsItemPreviewPostRequestBody and sets the default values.
func NewDrivesItemItemsItemPreviewPostRequestBody()(*DrivesItemItemsItemPreviewPostRequestBody) {
    m := &DrivesItemItemsItemPreviewPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDrivesItemItemsItemPreviewPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDrivesItemItemsItemPreviewPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDrivesItemItemsItemPreviewPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DrivesItemItemsItemPreviewPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DrivesItemItemsItemPreviewPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val)
        }
        return nil
    }
    res["zoom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoom(val)
        }
        return nil
    }
    return res
}
// GetPage gets the page property value. The page property
func (m *DrivesItemItemsItemPreviewPostRequestBody) GetPage()(*string) {
    return m.page
}
// GetZoom gets the zoom property value. The zoom property
func (m *DrivesItemItemsItemPreviewPostRequestBody) GetZoom()(*float64) {
    return m.zoom
}
// Serialize serializes information the current object
func (m *DrivesItemItemsItemPreviewPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("page", m.GetPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("zoom", m.GetZoom())
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
func (m *DrivesItemItemsItemPreviewPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPage sets the page property value. The page property
func (m *DrivesItemItemsItemPreviewPostRequestBody) SetPage(value *string)() {
    m.page = value
}
// SetZoom sets the zoom property value. The zoom property
func (m *DrivesItemItemsItemPreviewPostRequestBody) SetZoom(value *float64)() {
    m.zoom = value
}
