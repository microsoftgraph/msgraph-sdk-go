package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody provides operations to call the getByIds method.
type IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ids property
    ids []string
    // The types property
    types []string
}
// NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody instantiates a new IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody and sets the default values.
func NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) {
    m := &IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetIds)
    res["types"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTypes)
    return res
}
// GetIds gets the ids property value. The ids property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetTypes gets the types property value. The types property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) GetTypes()([]string) {
    return m.types
}
// Serialize serializes information the current object
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    if m.GetTypes() != nil {
        err := writer.WriteCollectionOfStringValues("types", m.GetTypes())
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
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetTypes sets the types property value. The types property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBody) SetTypes(value []string)() {
    m.types = value
}
