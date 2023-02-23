package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody 
type ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody instantiates a new ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody and sets the default values.
func NewItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody()(*ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) {
    m := &ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceTag gets the deviceTag property value. The deviceTag property
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) GetDeviceTag()(*string) {
    val, err := m.GetBackingStore().Get("deviceTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceTag", m.GetDeviceTag())
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
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceTag sets the deviceTag property value. The deviceTag property
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) SetDeviceTag(value *string)() {
    err := m.GetBackingStore().Set("deviceTag", value)
    if err != nil {
        panic(err)
    }
}
// ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable 
type ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceTag()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceTag(value *string)()
}
