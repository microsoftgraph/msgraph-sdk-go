package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceCompliancePoliciesItemAssignResponse provides operations to call the assign method.
type DeviceManagementDeviceCompliancePoliciesItemAssignResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceCompliancePolicyAssignmentable
}
// NewDeviceManagementDeviceCompliancePoliciesItemAssignResponse instantiates a new DeviceManagementDeviceCompliancePoliciesItemAssignResponse and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesItemAssignResponse()(*DeviceManagementDeviceCompliancePoliciesItemAssignResponse) {
    m := &DeviceManagementDeviceCompliancePoliciesItemAssignResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceManagementDeviceCompliancePoliciesItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceCompliancePoliciesItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceCompliancePoliciesItemAssignResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceCompliancePoliciesItemAssignResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicyAssignmentFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceManagementDeviceCompliancePoliciesItemAssignResponse) GetValue()([]DeviceCompliancePolicyAssignmentable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceCompliancePoliciesItemAssignResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DeviceManagementDeviceCompliancePoliciesItemAssignResponse) SetValue(value []DeviceCompliancePolicyAssignmentable)() {
    m.value = value
}
