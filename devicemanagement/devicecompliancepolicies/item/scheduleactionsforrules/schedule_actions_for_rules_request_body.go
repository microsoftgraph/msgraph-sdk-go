package scheduleactionsforrules

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ScheduleActionsForRulesRequestBody provides operations to call the scheduleActionsForRules method.
type ScheduleActionsForRulesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The deviceComplianceScheduledActionForRules property
    deviceComplianceScheduledActionForRules []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable;
}
// NewScheduleActionsForRulesRequestBody instantiates a new scheduleActionsForRulesRequestBody and sets the default values.
func NewScheduleActionsForRulesRequestBody()(*ScheduleActionsForRulesRequestBody) {
    m := &ScheduleActionsForRulesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateScheduleActionsForRulesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScheduleActionsForRulesRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScheduleActionsForRulesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScheduleActionsForRulesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceComplianceScheduledActionForRules gets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *ScheduleActionsForRulesRequestBody) GetDeviceComplianceScheduledActionForRules()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScheduledActionForRules
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScheduleActionsForRulesRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceComplianceScheduledActionForRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)
            }
            m.SetDeviceComplianceScheduledActionForRules(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ScheduleActionsForRulesRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceComplianceScheduledActionForRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceComplianceScheduledActionForRules()))
        for i, v := range m.GetDeviceComplianceScheduledActionForRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
func (m *ScheduleActionsForRulesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceComplianceScheduledActionForRules sets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *ScheduleActionsForRulesRequestBody) SetDeviceComplianceScheduledActionForRules(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)() {
    if m != nil {
        m.deviceComplianceScheduledActionForRules = value
    }
}
