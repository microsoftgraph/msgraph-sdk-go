package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemItemsItemPreviewPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom composed type wrapper for classes float64, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom instantiates a new ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom and sets the default values.
func NewItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom()(*ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) {
    m := &ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
// NewItemItemsItemPreviewPostRequestBody instantiates a new ItemItemsItemPreviewPostRequestBody and sets the default values.
func NewItemItemsItemPreviewPostRequestBody()(*ItemItemsItemPreviewPostRequestBody) {
    m := &ItemItemsItemPreviewPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemPreviewPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemPreviewPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemPreviewPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemsItemPreviewPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ItemItemsItemPreviewPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemsItemPreviewPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoom(val.(*ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoom))
        }
        return nil
    }
    return res
}
// GetPage gets the page property value. The page property
// returns a *string when successful
func (m *ItemItemsItemPreviewPostRequestBody) GetPage()(*string) {
    val, err := m.GetBackingStore().Get("page")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetZoom gets the zoom property value. The zoom property
// returns a ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomable when successful
func (m *ItemItemsItemPreviewPostRequestBody) GetZoom()(ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomable) {
    val, err := m.GetBackingStore().Get("zoom")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemItemsItemPreviewPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("page", m.GetPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("zoom", m.GetZoom())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemPreviewPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemItemsItemPreviewPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetPage sets the page property value. The page property
func (m *ItemItemsItemPreviewPostRequestBody) SetPage(value *string)() {
    err := m.GetBackingStore().Set("page", value)
    if err != nil {
        panic(err)
    }
}
// SetZoom sets the zoom property value. The zoom property
func (m *ItemItemsItemPreviewPostRequestBody) SetZoom(value ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomable)() {
    err := m.GetBackingStore().Set("zoom", value)
    if err != nil {
        panic(err)
    }
}
type ItemItemsItemPreviewPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetPage()(*string)
    GetZoom()(ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetPage(value *string)()
    SetZoom(value ItemItemsItemPreviewPostRequestBody_PreviewPostRequestBody_zoomable)()
}
