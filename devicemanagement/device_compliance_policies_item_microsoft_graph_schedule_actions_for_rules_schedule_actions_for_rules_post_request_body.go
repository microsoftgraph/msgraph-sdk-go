package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody 
type DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody instantiates a new DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody and sets the default values.
func NewDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody()(*DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) {
    m := &DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceComplianceScheduledActionForRules gets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) GetDeviceComplianceScheduledActionForRules()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable) {
    val, err := m.GetBackingStore().Get("deviceComplianceScheduledActionForRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceComplianceScheduledActionForRules sets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) SetDeviceComplianceScheduledActionForRules(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)() {
    err := m.GetBackingStore().Set("deviceComplianceScheduledActionForRules", value)
    if err != nil {
        panic(err)
    }
}
// DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBodyable 
type DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceComplianceScheduledActionForRules()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceComplianceScheduledActionForRules(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)()
}
