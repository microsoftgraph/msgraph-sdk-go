package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody 
type OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody instantiates a new OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody and sets the default values.
func NewOnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody()(*OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) {
    m := &OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
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
    res["renameAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenameAs(val)
        }
        return nil
    }
    res["siteCollectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollectionId(val)
        }
        return nil
    }
    res["siteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetId gets the id property value. The id property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRenameAs gets the renameAs property value. The renameAs property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetRenameAs()(*string) {
    val, err := m.GetBackingStore().Get("renameAs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetSiteCollectionId()(*string) {
    val, err := m.GetBackingStore().Get("siteCollectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSiteId gets the siteId property value. The siteId property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) GetSiteId()(*string) {
    val, err := m.GetBackingStore().Get("siteId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("renameAs", m.GetRenameAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteCollectionId", m.GetSiteCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteId", m.GetSiteId())
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
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. The id property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetRenameAs sets the renameAs property value. The renameAs property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) SetRenameAs(value *string)() {
    err := m.GetBackingStore().Set("renameAs", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) SetSiteCollectionId(value *string)() {
    err := m.GetBackingStore().Set("siteCollectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteId sets the siteId property value. The siteId property
func (m *OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBody) SetSiteId(value *string)() {
    err := m.GetBackingStore().Set("siteId", value)
    if err != nil {
        panic(err)
    }
}
// OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBodyable 
type OnenoteSectionsItemMicrosoftGraphCopyToNotebookCopyToNotebookPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetGroupId()(*string)
    GetId()(*string)
    GetRenameAs()(*string)
    GetSiteCollectionId()(*string)
    GetSiteId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetGroupId(value *string)()
    SetId(value *string)()
    SetRenameAs(value *string)()
    SetSiteCollectionId(value *string)()
    SetSiteId(value *string)()
}
