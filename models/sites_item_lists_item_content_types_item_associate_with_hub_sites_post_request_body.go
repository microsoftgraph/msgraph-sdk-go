package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody provides operations to call the associateWithHubSites method.
type SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The hubSiteUrls property
    hubSiteUrls []string
    // The propagateToExistingLists property
    propagateToExistingLists *bool
}
// NewSitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody instantiates a new SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody and sets the default values.
func NewSitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody()(*SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) {
    m := &SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hubSiteUrls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetHubSiteUrls)
    res["propagateToExistingLists"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPropagateToExistingLists)
    return res
}
// GetHubSiteUrls gets the hubSiteUrls property value. The hubSiteUrls property
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetHubSiteUrls()([]string) {
    return m.hubSiteUrls
}
// GetPropagateToExistingLists gets the propagateToExistingLists property value. The propagateToExistingLists property
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetPropagateToExistingLists()(*bool) {
    return m.propagateToExistingLists
}
// Serialize serializes information the current object
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHubSiteUrls() != nil {
        err := writer.WriteCollectionOfStringValues("hubSiteUrls", m.GetHubSiteUrls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateToExistingLists", m.GetPropagateToExistingLists())
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
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHubSiteUrls sets the hubSiteUrls property value. The hubSiteUrls property
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) SetHubSiteUrls(value []string)() {
    m.hubSiteUrls = value
}
// SetPropagateToExistingLists sets the propagateToExistingLists property value. The propagateToExistingLists property
func (m *SitesItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) SetPropagateToExistingLists(value *bool)() {
    m.propagateToExistingLists = value
}
