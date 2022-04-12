package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestorSettings 
type AccessPackageAssignmentRequestorSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If false, the requestor is not permitted to include a schedule in their request.
    allowCustomAssignmentSchedule *bool;
    // If true, allows on-behalf-of requestors to create a request to add access for another principal.
    enableOnBehalfRequestorsToAddAccess *bool;
    // If true, allows on-behalf-of requestors to create a request to remove access for another principal.
    enableOnBehalfRequestorsToRemoveAccess *bool;
    // If true, allows on-behalf-of requestors to create a request to update access for another principal.
    enableOnBehalfRequestorsToUpdateAccess *bool;
    // If true, allows requestors to create a request to add access for themselves.
    enableTargetsToSelfAddAccess *bool;
    // If true, allows requestors to create a request to remove their access.
    enableTargetsToSelfRemoveAccess *bool;
    // If true, allows requestors to create a request to update their access.
    enableTargetsToSelfUpdateAccess *bool;
    // The principals who can request on-behalf-of others.
    onBehalfRequestors []SubjectSetable;
}
// NewAccessPackageAssignmentRequestorSettings instantiates a new accessPackageAssignmentRequestorSettings and sets the default values.
func NewAccessPackageAssignmentRequestorSettings()(*AccessPackageAssignmentRequestorSettings) {
    m := &AccessPackageAssignmentRequestorSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageAssignmentRequestorSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestorSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequestorSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentRequestorSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowCustomAssignmentSchedule gets the allowCustomAssignmentSchedule property value. If false, the requestor is not permitted to include a schedule in their request.
func (m *AccessPackageAssignmentRequestorSettings) GetAllowCustomAssignmentSchedule()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCustomAssignmentSchedule
    }
}
// GetEnableOnBehalfRequestorsToAddAccess gets the enableOnBehalfRequestorsToAddAccess property value. If true, allows on-behalf-of requestors to create a request to add access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) GetEnableOnBehalfRequestorsToAddAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableOnBehalfRequestorsToAddAccess
    }
}
// GetEnableOnBehalfRequestorsToRemoveAccess gets the enableOnBehalfRequestorsToRemoveAccess property value. If true, allows on-behalf-of requestors to create a request to remove access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) GetEnableOnBehalfRequestorsToRemoveAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableOnBehalfRequestorsToRemoveAccess
    }
}
// GetEnableOnBehalfRequestorsToUpdateAccess gets the enableOnBehalfRequestorsToUpdateAccess property value. If true, allows on-behalf-of requestors to create a request to update access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) GetEnableOnBehalfRequestorsToUpdateAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableOnBehalfRequestorsToUpdateAccess
    }
}
// GetEnableTargetsToSelfAddAccess gets the enableTargetsToSelfAddAccess property value. If true, allows requestors to create a request to add access for themselves.
func (m *AccessPackageAssignmentRequestorSettings) GetEnableTargetsToSelfAddAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableTargetsToSelfAddAccess
    }
}
// GetEnableTargetsToSelfRemoveAccess gets the enableTargetsToSelfRemoveAccess property value. If true, allows requestors to create a request to remove their access.
func (m *AccessPackageAssignmentRequestorSettings) GetEnableTargetsToSelfRemoveAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableTargetsToSelfRemoveAccess
    }
}
// GetEnableTargetsToSelfUpdateAccess gets the enableTargetsToSelfUpdateAccess property value. If true, allows requestors to create a request to update their access.
func (m *AccessPackageAssignmentRequestorSettings) GetEnableTargetsToSelfUpdateAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableTargetsToSelfUpdateAccess
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequestorSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowCustomAssignmentSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCustomAssignmentSchedule(val)
        }
        return nil
    }
    res["enableOnBehalfRequestorsToAddAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOnBehalfRequestorsToAddAccess(val)
        }
        return nil
    }
    res["enableOnBehalfRequestorsToRemoveAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOnBehalfRequestorsToRemoveAccess(val)
        }
        return nil
    }
    res["enableOnBehalfRequestorsToUpdateAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOnBehalfRequestorsToUpdateAccess(val)
        }
        return nil
    }
    res["enableTargetsToSelfAddAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTargetsToSelfAddAccess(val)
        }
        return nil
    }
    res["enableTargetsToSelfRemoveAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTargetsToSelfRemoveAccess(val)
        }
        return nil
    }
    res["enableTargetsToSelfUpdateAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTargetsToSelfUpdateAccess(val)
        }
        return nil
    }
    res["onBehalfRequestors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetOnBehalfRequestors(res)
        }
        return nil
    }
    return res
}
// GetOnBehalfRequestors gets the onBehalfRequestors property value. The principals who can request on-behalf-of others.
func (m *AccessPackageAssignmentRequestorSettings) GetOnBehalfRequestors()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.onBehalfRequestors
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestorSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCustomAssignmentSchedule", m.GetAllowCustomAssignmentSchedule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableOnBehalfRequestorsToAddAccess", m.GetEnableOnBehalfRequestorsToAddAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableOnBehalfRequestorsToRemoveAccess", m.GetEnableOnBehalfRequestorsToRemoveAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableOnBehalfRequestorsToUpdateAccess", m.GetEnableOnBehalfRequestorsToUpdateAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTargetsToSelfAddAccess", m.GetEnableTargetsToSelfAddAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTargetsToSelfRemoveAccess", m.GetEnableTargetsToSelfRemoveAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTargetsToSelfUpdateAccess", m.GetEnableTargetsToSelfUpdateAccess())
        if err != nil {
            return err
        }
    }
    if m.GetOnBehalfRequestors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnBehalfRequestors()))
        for i, v := range m.GetOnBehalfRequestors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("onBehalfRequestors", cast)
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
func (m *AccessPackageAssignmentRequestorSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowCustomAssignmentSchedule sets the allowCustomAssignmentSchedule property value. If false, the requestor is not permitted to include a schedule in their request.
func (m *AccessPackageAssignmentRequestorSettings) SetAllowCustomAssignmentSchedule(value *bool)() {
    if m != nil {
        m.allowCustomAssignmentSchedule = value
    }
}
// SetEnableOnBehalfRequestorsToAddAccess sets the enableOnBehalfRequestorsToAddAccess property value. If true, allows on-behalf-of requestors to create a request to add access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableOnBehalfRequestorsToAddAccess(value *bool)() {
    if m != nil {
        m.enableOnBehalfRequestorsToAddAccess = value
    }
}
// SetEnableOnBehalfRequestorsToRemoveAccess sets the enableOnBehalfRequestorsToRemoveAccess property value. If true, allows on-behalf-of requestors to create a request to remove access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableOnBehalfRequestorsToRemoveAccess(value *bool)() {
    if m != nil {
        m.enableOnBehalfRequestorsToRemoveAccess = value
    }
}
// SetEnableOnBehalfRequestorsToUpdateAccess sets the enableOnBehalfRequestorsToUpdateAccess property value. If true, allows on-behalf-of requestors to create a request to update access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableOnBehalfRequestorsToUpdateAccess(value *bool)() {
    if m != nil {
        m.enableOnBehalfRequestorsToUpdateAccess = value
    }
}
// SetEnableTargetsToSelfAddAccess sets the enableTargetsToSelfAddAccess property value. If true, allows requestors to create a request to add access for themselves.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableTargetsToSelfAddAccess(value *bool)() {
    if m != nil {
        m.enableTargetsToSelfAddAccess = value
    }
}
// SetEnableTargetsToSelfRemoveAccess sets the enableTargetsToSelfRemoveAccess property value. If true, allows requestors to create a request to remove their access.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableTargetsToSelfRemoveAccess(value *bool)() {
    if m != nil {
        m.enableTargetsToSelfRemoveAccess = value
    }
}
// SetEnableTargetsToSelfUpdateAccess sets the enableTargetsToSelfUpdateAccess property value. If true, allows requestors to create a request to update their access.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableTargetsToSelfUpdateAccess(value *bool)() {
    if m != nil {
        m.enableTargetsToSelfUpdateAccess = value
    }
}
// SetOnBehalfRequestors sets the onBehalfRequestors property value. The principals who can request on-behalf-of others.
func (m *AccessPackageAssignmentRequestorSettings) SetOnBehalfRequestors(value []SubjectSetable)() {
    if m != nil {
        m.onBehalfRequestors = value
    }
}
