package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LicenseDetails struct {
    Entity
    // Information about the service plans assigned with the license. Read-only, Not nullable
    servicePlans []ServicePlanInfo;
    // Unique identifier (GUID) for the service SKU. Equal to the skuId property on the related SubscribedSku object. Read-only
    skuId *string;
    // Unique SKU display name. Equal to the skuPartNumber on the related SubscribedSku object; for example: 'AAD_Premium'. Read-only
    skuPartNumber *string;
}
// Instantiates a new licenseDetails and sets the default values.
func NewLicenseDetails()(*LicenseDetails) {
    m := &LicenseDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the servicePlans property value. Information about the service plans assigned with the license. Read-only, Not nullable
func (m *LicenseDetails) GetServicePlans()([]ServicePlanInfo) {
    if m == nil {
        return nil
    } else {
        return m.servicePlans
    }
}
// Gets the skuId property value. Unique identifier (GUID) for the service SKU. Equal to the skuId property on the related SubscribedSku object. Read-only
func (m *LicenseDetails) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// Gets the skuPartNumber property value. Unique SKU display name. Equal to the skuPartNumber on the related SubscribedSku object; for example: 'AAD_Premium'. Read-only
func (m *LicenseDetails) GetSkuPartNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuPartNumber
    }
}
// The deserialization information for the current model
func (m *LicenseDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["servicePlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePlanInfo() })
        if err != nil {
            return err
        }
        res := make([]ServicePlanInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServicePlanInfo))
        }
        m.SetServicePlans(res)
        return nil
    }
    res["skuId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSkuId(val)
        return nil
    }
    res["skuPartNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSkuPartNumber(val)
        return nil
    }
    return res
}
func (m *LicenseDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LicenseDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
// Sets the servicePlans property value. Information about the service plans assigned with the license. Read-only, Not nullable
// Parameters:
//  - value : Value to set for the servicePlans property.
func (m *LicenseDetails) SetServicePlans(value []ServicePlanInfo)() {
    m.servicePlans = value
}
// Sets the skuId property value. Unique identifier (GUID) for the service SKU. Equal to the skuId property on the related SubscribedSku object. Read-only
// Parameters:
//  - value : Value to set for the skuId property.
func (m *LicenseDetails) SetSkuId(value *string)() {
    m.skuId = value
}
// Sets the skuPartNumber property value. Unique SKU display name. Equal to the skuPartNumber on the related SubscribedSku object; for example: 'AAD_Premium'. Read-only
// Parameters:
//  - value : Value to set for the skuPartNumber property.
func (m *LicenseDetails) SetSkuPartNumber(value *string)() {
    m.skuPartNumber = value
}
