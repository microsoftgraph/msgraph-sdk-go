package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody provides operations to call the scheduleActionsForRules method.
type DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceComplianceScheduledActionForRules property
    deviceComplianceScheduledActionForRules []DeviceComplianceScheduledActionForRuleable
}
// NewDeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody instantiates a new DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody()(*DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody) {
    m := &DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceComplianceScheduledActionForRules gets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody) GetDeviceComplianceScheduledActionForRules()([]DeviceComplianceScheduledActionForRuleable) {
    return m.deviceComplianceScheduledActionForRules
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceComplianceScheduledActionForRules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue , m.SetDeviceComplianceScheduledActionForRules)
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceComplianceScheduledActionForRules() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceComplianceScheduledActionForRules())
        err := writer.WriteCollectionOfObjectValues("deviceComplianceScheduledActionForRules", cast)
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
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceComplianceScheduledActionForRules sets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesPostRequestBody) SetDeviceComplianceScheduledActionForRules(value []DeviceComplianceScheduledActionForRuleable)() {
    m.deviceComplianceScheduledActionForRules = value
}
