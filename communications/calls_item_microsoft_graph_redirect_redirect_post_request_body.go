package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CallsItemMicrosoftGraphRedirectRedirectPostRequestBody 
type CallsItemMicrosoftGraphRedirectRedirectPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCallsItemMicrosoftGraphRedirectRedirectPostRequestBody instantiates a new CallsItemMicrosoftGraphRedirectRedirectPostRequestBody and sets the default values.
func NewCallsItemMicrosoftGraphRedirectRedirectPostRequestBody()(*CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) {
    m := &CallsItemMicrosoftGraphRedirectRedirectPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsItemMicrosoftGraphRedirectRedirectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemMicrosoftGraphRedirectRedirectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemMicrosoftGraphRedirectRedirectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) GetCallbackUri()(*string) {
    val, err := m.GetBackingStore().Get("callbackUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable)
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["timeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeout(val)
        }
        return nil
    }
    return res
}
// GetTargets gets the targets property value. The targets property
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) GetTargets()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable) {
    val, err := m.GetBackingStore().Get("targets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable)
    }
    return nil
}
// GetTimeout gets the timeout property value. The timeout property
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) GetTimeout()(*int32) {
    val, err := m.GetBackingStore().Get("timeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeout", m.GetTimeout())
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
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) SetCallbackUri(value *string)() {
    err := m.GetBackingStore().Set("callbackUri", value)
    if err != nil {
        panic(err)
    }
}
// SetTargets sets the targets property value. The targets property
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) SetTargets(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable)() {
    err := m.GetBackingStore().Set("targets", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeout sets the timeout property value. The timeout property
func (m *CallsItemMicrosoftGraphRedirectRedirectPostRequestBody) SetTimeout(value *int32)() {
    err := m.GetBackingStore().Set("timeout", value)
    if err != nil {
        panic(err)
    }
}
// CallsItemMicrosoftGraphRedirectRedirectPostRequestBodyable 
type CallsItemMicrosoftGraphRedirectRedirectPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCallbackUri()(*string)
    GetTargets()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable)
    GetTimeout()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCallbackUri(value *string)()
    SetTargets(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable)()
    SetTimeout(value *int32)()
}
