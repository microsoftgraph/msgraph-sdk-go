package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse provides operations to call the import method.
type DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ImportedWindowsAutopilotDeviceIdentityable
}
// NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse instantiates a new DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse and sets the default values.
func NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse()(*DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse) {
    m := &DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedWindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ImportedWindowsAutopilotDeviceIdentityable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse) GetValue()([]ImportedWindowsAutopilotDeviceIdentityable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponse) SetValue(value []ImportedWindowsAutopilotDeviceIdentityable)() {
    m.value = value
}
