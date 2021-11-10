package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SubscribedSku struct {
    Entity
    // For example, 'User' or 'Company'.
    appliesTo *string;
    // Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
    capabilityStatus *string;
    // The number of licenses that have been assigned.
    consumedUnits *int32;
    // Information about the number and status of prepaid licenses.
    prepaidUnits *LicenseUnitsDetail;
    // Information about the service plans that are available with the SKU. Not nullable
    servicePlans []ServicePlanInfo;
    // The unique identifier (GUID) for the service SKU.
    skuId *string;
    // The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
    skuPartNumber *string;
}
// Instantiates a new subscribedSku and sets the default values.
func NewSubscribedSku()(*SubscribedSku) {
    m := &SubscribedSku{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appliesTo property value. For example, 'User' or 'Company'.
func (m *SubscribedSku) GetAppliesTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// Gets the capabilityStatus property value. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
func (m *SubscribedSku) GetCapabilityStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capabilityStatus
    }
}
// Gets the consumedUnits property value. The number of licenses that have been assigned.
func (m *SubscribedSku) GetConsumedUnits()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.consumedUnits
    }
}
// Gets the prepaidUnits property value. Information about the number and status of prepaid licenses.
func (m *SubscribedSku) GetPrepaidUnits()(*LicenseUnitsDetail) {
    if m == nil {
        return nil
    } else {
        return m.prepaidUnits
    }
}
// Gets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable
func (m *SubscribedSku) GetServicePlans()([]ServicePlanInfo) {
    if m == nil {
        return nil
    } else {
        return m.servicePlans
    }
}
// Gets the skuId property value. The unique identifier (GUID) for the service SKU.
func (m *SubscribedSku) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// Gets the skuPartNumber property value. The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
func (m *SubscribedSku) GetSkuPartNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuPartNumber
    }
}
// The deserialization information for the current model
func (m *SubscribedSku) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val)
        }
        return nil
    }
    res["capabilityStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilityStatus(val)
        }
        return nil
    }
    res["consumedUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsumedUnits(val)
        }
        return nil
    }
    res["prepaidUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseUnitsDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrepaidUnits(val.(*LicenseUnitsDetail))
        }
        return nil
    }
    res["servicePlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePlanInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePlanInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServicePlanInfo))
            }
            m.SetServicePlans(res)
        }
        return nil
    }
    res["skuId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    res["skuPartNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuPartNumber(val)
        }
        return nil
    }
    return res
}
func (m *SubscribedSku) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SubscribedSku) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appliesTo", m.GetAppliesTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("capabilityStatus", m.GetCapabilityStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("consumedUnits", m.GetConsumedUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("prepaidUnits", m.GetPrepaidUnits())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServicePlans()))
        for i, v := range m.GetServicePlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("servicePlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuId", m.GetSkuId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuPartNumber", m.GetSkuPartNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appliesTo property value. For example, 'User' or 'Company'.
// Parameters:
//  - value : Value to set for the appliesTo property.
func (m *SubscribedSku) SetAppliesTo(value *string)() {
    m.appliesTo = value
}
// Sets the capabilityStatus property value. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
// Parameters:
//  - value : Value to set for the capabilityStatus property.
func (m *SubscribedSku) SetCapabilityStatus(value *string)() {
    m.capabilityStatus = value
}
// Sets the consumedUnits property value. The number of licenses that have been assigned.
// Parameters:
//  - value : Value to set for the consumedUnits property.
func (m *SubscribedSku) SetConsumedUnits(value *int32)() {
    m.consumedUnits = value
}
// Sets the prepaidUnits property value. Information about the number and status of prepaid licenses.
// Parameters:
//  - value : Value to set for the prepaidUnits property.
func (m *SubscribedSku) SetPrepaidUnits(value *LicenseUnitsDetail)() {
    m.prepaidUnits = value
}
// Sets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable
// Parameters:
//  - value : Value to set for the servicePlans property.
func (m *SubscribedSku) SetServicePlans(value []ServicePlanInfo)() {
    m.servicePlans = value
}
// Sets the skuId property value. The unique identifier (GUID) for the service SKU.
// Parameters:
//  - value : Value to set for the skuId property.
func (m *SubscribedSku) SetSkuId(value *string)() {
    m.skuId = value
}
// Sets the skuPartNumber property value. The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
// Parameters:
//  - value : Value to set for the skuPartNumber property.
func (m *SubscribedSku) SetSkuPartNumber(value *string)() {
    m.skuPartNumber = value
}
