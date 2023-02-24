package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody 
type CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody instantiates a new CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody and sets the default values.
func NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody()(*CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) {
    m := &CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseScreenSharingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. The role property
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) GetRole()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
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
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetRole sets the role property value. The role property
func (m *CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBody) SetRole(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBodyable 
type CallsItemMicrosoftGraphChangeScreenSharingRoleChangeScreenSharingRolePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetRole()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetRole(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole)()
}
