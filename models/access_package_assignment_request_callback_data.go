package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestCallbackData 
type AccessPackageAssignmentRequestCallbackData struct {
    CustomExtensionData
}
// NewAccessPackageAssignmentRequestCallbackData instantiates a new accessPackageAssignmentRequestCallbackData and sets the default values.
func NewAccessPackageAssignmentRequestCallbackData()(*AccessPackageAssignmentRequestCallbackData) {
    m := &AccessPackageAssignmentRequestCallbackData{
        CustomExtensionData: *NewCustomExtensionData(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageAssignmentRequestCallbackData"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessPackageAssignmentRequestCallbackDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestCallbackDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequestCallbackData(), nil
}
// GetCustomExtensionStageInstanceDetail gets the customExtensionStageInstanceDetail property value. The customExtensionStageInstanceDetail property
func (m *AccessPackageAssignmentRequestCallbackData) GetCustomExtensionStageInstanceDetail()(*string) {
    val, err := m.GetBackingStore().Get("customExtensionStageInstanceDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomExtensionStageInstanceId gets the customExtensionStageInstanceId property value. The customExtensionStageInstanceId property
func (m *AccessPackageAssignmentRequestCallbackData) GetCustomExtensionStageInstanceId()(*string) {
    val, err := m.GetBackingStore().Get("customExtensionStageInstanceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequestCallbackData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionData.GetFieldDeserializers()
    res["customExtensionStageInstanceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtensionStageInstanceDetail(val)
        }
        return nil
    }
    res["customExtensionStageInstanceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtensionStageInstanceId(val)
        }
        return nil
    }
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*AccessPackageCustomExtensionStage))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetStage gets the stage property value. The stage property
func (m *AccessPackageAssignmentRequestCallbackData) GetStage()(*AccessPackageCustomExtensionStage) {
    val, err := m.GetBackingStore().Get("stage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AccessPackageCustomExtensionStage)
    }
    return nil
}
// GetState gets the state property value. The state property
func (m *AccessPackageAssignmentRequestCallbackData) GetState()(*string) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestCallbackData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionData.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("customExtensionStageInstanceDetail", m.GetCustomExtensionStageInstanceDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customExtensionStageInstanceId", m.GetCustomExtensionStageInstanceId())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err = writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtensionStageInstanceDetail sets the customExtensionStageInstanceDetail property value. The customExtensionStageInstanceDetail property
func (m *AccessPackageAssignmentRequestCallbackData) SetCustomExtensionStageInstanceDetail(value *string)() {
    err := m.GetBackingStore().Set("customExtensionStageInstanceDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtensionStageInstanceId sets the customExtensionStageInstanceId property value. The customExtensionStageInstanceId property
func (m *AccessPackageAssignmentRequestCallbackData) SetCustomExtensionStageInstanceId(value *string)() {
    err := m.GetBackingStore().Set("customExtensionStageInstanceId", value)
    if err != nil {
        panic(err)
    }
}
// SetStage sets the stage property value. The stage property
func (m *AccessPackageAssignmentRequestCallbackData) SetStage(value *AccessPackageCustomExtensionStage)() {
    err := m.GetBackingStore().Set("stage", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *AccessPackageAssignmentRequestCallbackData) SetState(value *string)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageAssignmentRequestCallbackDataable 
type AccessPackageAssignmentRequestCallbackDataable interface {
    CustomExtensionDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomExtensionStageInstanceDetail()(*string)
    GetCustomExtensionStageInstanceId()(*string)
    GetStage()(*AccessPackageCustomExtensionStage)
    GetState()(*string)
    SetCustomExtensionStageInstanceDetail(value *string)()
    SetCustomExtensionStageInstanceId(value *string)()
    SetStage(value *AccessPackageCustomExtensionStage)()
    SetState(value *string)()
}
