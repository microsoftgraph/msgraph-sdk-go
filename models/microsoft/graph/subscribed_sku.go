package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubscribedSku 
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
// NewSubscribedSku instantiates a new subscribedSku and sets the default values.
func NewSubscribedSku()(*SubscribedSku) {
    m := &SubscribedSku{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppliesTo gets the appliesTo property value. For example, 'User' or 'Company'.
func (m *SubscribedSku) GetAppliesTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetCapabilityStatus gets the capabilityStatus property value. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
func (m *SubscribedSku) GetCapabilityStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capabilityStatus
    }
}
// GetConsumedUnits gets the consumedUnits property value. The number of licenses that have been assigned.
func (m *SubscribedSku) GetConsumedUnits()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.consumedUnits
    }
}
// GetPrepaidUnits gets the prepaidUnits property value. Information about the number and status of prepaid licenses.
func (m *SubscribedSku) GetPrepaidUnits()(*LicenseUnitsDetail) {
    if m == nil {
        return nil
    } else {
        return m.prepaidUnits
    }
}
// GetServicePlans gets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable
func (m *SubscribedSku) GetServicePlans()([]ServicePlanInfo) {
    if m == nil {
        return nil
    } else {
        return m.servicePlans
    }
}
// GetSkuId gets the skuId property value. The unique identifier (GUID) for the service SKU.
func (m *SubscribedSku) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// GetSkuPartNumber gets the skuPartNumber property value. The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
func (m *SubscribedSku) GetSkuPartNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuPartNumber
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
    if m.GetServicePlans() != nil {
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
// SetAppliesTo sets the appliesTo property value. For example, 'User' or 'Company'.
func (m *SubscribedSku) SetAppliesTo(value *string)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetCapabilityStatus sets the capabilityStatus property value. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
func (m *SubscribedSku) SetCapabilityStatus(value *string)() {
    if m != nil {
        m.capabilityStatus = value
    }
}
// SetConsumedUnits sets the consumedUnits property value. The number of licenses that have been assigned.
func (m *SubscribedSku) SetConsumedUnits(value *int32)() {
    if m != nil {
        m.consumedUnits = value
    }
}
// SetPrepaidUnits sets the prepaidUnits property value. Information about the number and status of prepaid licenses.
func (m *SubscribedSku) SetPrepaidUnits(value *LicenseUnitsDetail)() {
    if m != nil {
        m.prepaidUnits = value
    }
}
// SetServicePlans sets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable
func (m *SubscribedSku) SetServicePlans(value []ServicePlanInfo)() {
    if m != nil {
        m.servicePlans = value
    }
}
// SetSkuId sets the skuId property value. The unique identifier (GUID) for the service SKU.
func (m *SubscribedSku) SetSkuId(value *string)() {
    if m != nil {
        m.skuId = value
    }
}
// SetSkuPartNumber sets the skuPartNumber property value. The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
func (m *SubscribedSku) SetSkuPartNumber(value *string)() {
    if m != nil {
        m.skuPartNumber = value
    }
}
