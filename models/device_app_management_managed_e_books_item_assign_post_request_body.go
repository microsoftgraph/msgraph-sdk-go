package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementManagedEBooksItemAssignPostRequestBody provides operations to call the assign method.
type DeviceAppManagementManagedEBooksItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managedEBookAssignments property
    managedEBookAssignments []ManagedEBookAssignmentable
}
// NewDeviceAppManagementManagedEBooksItemAssignPostRequestBody instantiates a new DeviceAppManagementManagedEBooksItemAssignPostRequestBody and sets the default values.
func NewDeviceAppManagementManagedEBooksItemAssignPostRequestBody()(*DeviceAppManagementManagedEBooksItemAssignPostRequestBody) {
    m := &DeviceAppManagementManagedEBooksItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementManagedEBooksItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementManagedEBooksItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementManagedEBooksItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementManagedEBooksItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementManagedEBooksItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedEBookAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedEBookAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedEBookAssignmentable)
            }
            m.SetManagedEBookAssignments(res)
        }
        return nil
    }
    return res
}
// GetManagedEBookAssignments gets the managedEBookAssignments property value. The managedEBookAssignments property
func (m *DeviceAppManagementManagedEBooksItemAssignPostRequestBody) GetManagedEBookAssignments()([]ManagedEBookAssignmentable) {
    return m.managedEBookAssignments
}
// Serialize serializes information the current object
func (m *DeviceAppManagementManagedEBooksItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedEBookAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedEBookAssignments()))
        for i, v := range m.GetManagedEBookAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("managedEBookAssignments", cast)
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
func (m *DeviceAppManagementManagedEBooksItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagedEBookAssignments sets the managedEBookAssignments property value. The managedEBookAssignments property
func (m *DeviceAppManagementManagedEBooksItemAssignPostRequestBody) SetManagedEBookAssignments(value []ManagedEBookAssignmentable)() {
    m.managedEBookAssignments = value
}
