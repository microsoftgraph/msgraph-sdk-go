package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServicePlanInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The object the service plan can be assigned to. Possible values:'User' - service plan can be assigned to individual users.'Company' - service plan can be assigned to the entire tenant.
    appliesTo *string;
    // The provisioning status of the service plan. Possible values:'Success' - Service is fully provisioned.'Disabled' - Service has been disabled.'PendingInput' - Service is not yet provisioned; awaiting service confirmation.'PendingActivation' - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)'PendingProvisioning' - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
    provisioningStatus *string;
    // The unique identifier of the service plan.
    servicePlanId *string;
    // The name of the service plan.
    servicePlanName *string;
}
// Instantiates a new servicePlanInfo and sets the default values.
func NewServicePlanInfo()(*ServicePlanInfo) {
    m := &ServicePlanInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePlanInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appliesTo property value. The object the service plan can be assigned to. Possible values:'User' - service plan can be assigned to individual users.'Company' - service plan can be assigned to the entire tenant.
func (m *ServicePlanInfo) GetAppliesTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// Gets the provisioningStatus property value. The provisioning status of the service plan. Possible values:'Success' - Service is fully provisioned.'Disabled' - Service has been disabled.'PendingInput' - Service is not yet provisioned; awaiting service confirmation.'PendingActivation' - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)'PendingProvisioning' - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
func (m *ServicePlanInfo) GetProvisioningStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatus
    }
}
// Gets the servicePlanId property value. The unique identifier of the service plan.
func (m *ServicePlanInfo) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// Gets the servicePlanName property value. The name of the service plan.
func (m *ServicePlanInfo) GetServicePlanName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanName
    }
}
// The deserialization information for the current model
func (m *ServicePlanInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppliesTo(val)
        return nil
    }
    res["provisioningStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProvisioningStatus(val)
        return nil
    }
    res["servicePlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePlanId(val)
        return nil
    }
    res["servicePlanName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePlanName(val)
        return nil
    }
    return res
}
func (m *ServicePlanInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServicePlanInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appliesTo", m.GetAppliesTo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("provisioningStatus", m.GetProvisioningStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePlanName", m.GetServicePlanName())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ServicePlanInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appliesTo property value. The object the service plan can be assigned to. Possible values:'User' - service plan can be assigned to individual users.'Company' - service plan can be assigned to the entire tenant.
// Parameters:
//  - value : Value to set for the appliesTo property.
func (m *ServicePlanInfo) SetAppliesTo(value *string)() {
    m.appliesTo = value
}
// Sets the provisioningStatus property value. The provisioning status of the service plan. Possible values:'Success' - Service is fully provisioned.'Disabled' - Service has been disabled.'PendingInput' - Service is not yet provisioned; awaiting service confirmation.'PendingActivation' - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)'PendingProvisioning' - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
// Parameters:
//  - value : Value to set for the provisioningStatus property.
func (m *ServicePlanInfo) SetProvisioningStatus(value *string)() {
    m.provisioningStatus = value
}
// Sets the servicePlanId property value. The unique identifier of the service plan.
// Parameters:
//  - value : Value to set for the servicePlanId property.
func (m *ServicePlanInfo) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// Sets the servicePlanName property value. The name of the service plan.
// Parameters:
//  - value : Value to set for the servicePlanName property.
func (m *ServicePlanInfo) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
