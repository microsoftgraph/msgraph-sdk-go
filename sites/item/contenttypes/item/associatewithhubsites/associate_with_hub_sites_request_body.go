package associatewithhubsites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssociateWithHubSitesRequestBody provides operations to call the associateWithHubSites method.
type AssociateWithHubSitesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The hubSiteUrls property
    hubSiteUrls []string;
    // The propagateToExistingLists property
    propagateToExistingLists *bool;
}
// NewAssociateWithHubSitesRequestBody instantiates a new associateWithHubSitesRequestBody and sets the default values.
func NewAssociateWithHubSitesRequestBody()(*AssociateWithHubSitesRequestBody) {
    m := &AssociateWithHubSitesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssociateWithHubSitesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssociateWithHubSitesRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssociateWithHubSitesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssociateWithHubSitesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssociateWithHubSitesRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hubSiteUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetHubSiteUrls(res)
        }
        return nil
    }
    res["propagateToExistingLists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropagateToExistingLists(val)
        }
        return nil
    }
    return res
}
// GetHubSiteUrls gets the hubSiteUrls property value. The hubSiteUrls property
func (m *AssociateWithHubSitesRequestBody) GetHubSiteUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.hubSiteUrls
    }
}
// GetPropagateToExistingLists gets the propagateToExistingLists property value. The propagateToExistingLists property
func (m *AssociateWithHubSitesRequestBody) GetPropagateToExistingLists()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateToExistingLists
    }
}
// Serialize serializes information the current object
func (m *AssociateWithHubSitesRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AssociateWithHubSitesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHubSiteUrls sets the hubSiteUrls property value. The hubSiteUrls property
func (m *AssociateWithHubSitesRequestBody) SetHubSiteUrls(value []string)() {
    if m != nil {
        m.hubSiteUrls = value
    }
}
// SetPropagateToExistingLists sets the propagateToExistingLists property value. The propagateToExistingLists property
func (m *AssociateWithHubSitesRequestBody) SetPropagateToExistingLists(value *bool)() {
    if m != nil {
        m.propagateToExistingLists = value
    }
}
