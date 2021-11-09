package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IdentityGovernance struct {
    // 
    accessReviews *AccessReviewSet;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    appConsent *AppConsentApprovalRoute;
    // 
    entitlementManagement *EntitlementManagement;
    // 
    termsOfUse *TermsOfUseContainer;
}
// Instantiates a new IdentityGovernance and sets the default values.
func NewIdentityGovernance()(*IdentityGovernance) {
    m := &IdentityGovernance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the accessReviews property value. 
func (m *IdentityGovernance) GetAccessReviews()(*AccessReviewSet) {
    if m == nil {
        return nil
    } else {
        return m.accessReviews
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appConsent property value. 
func (m *IdentityGovernance) GetAppConsent()(*AppConsentApprovalRoute) {
    if m == nil {
        return nil
    } else {
        return m.appConsent
    }
}
// Gets the entitlementManagement property value. 
func (m *IdentityGovernance) GetEntitlementManagement()(*EntitlementManagement) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
// Gets the termsOfUse property value. 
func (m *IdentityGovernance) GetTermsOfUse()(*TermsOfUseContainer) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUse
    }
}
// The deserialization information for the current model
func (m *IdentityGovernance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessReviews"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviews(val.(*AccessReviewSet))
        }
        return nil
    }
    res["appConsent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppConsentApprovalRoute() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppConsent(val.(*AppConsentApprovalRoute))
        }
        return nil
    }
    res["entitlementManagement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntitlementManagement() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitlementManagement(val.(*EntitlementManagement))
        }
        return nil
    }
    res["termsOfUse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsOfUseContainer() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsOfUse(val.(*TermsOfUseContainer))
        }
        return nil
    }
    return res
}
func (m *IdentityGovernance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IdentityGovernance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviews", m.GetAccessReviews())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("appConsent", m.GetAppConsent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("entitlementManagement", m.GetEntitlementManagement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("termsOfUse", m.GetTermsOfUse())
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
// Sets the accessReviews property value. 
// Parameters:
//  - value : Value to set for the accessReviews property.
func (m *IdentityGovernance) SetAccessReviews(value *AccessReviewSet)() {
    m.accessReviews = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *IdentityGovernance) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appConsent property value. 
// Parameters:
//  - value : Value to set for the appConsent property.
func (m *IdentityGovernance) SetAppConsent(value *AppConsentApprovalRoute)() {
    m.appConsent = value
}
// Sets the entitlementManagement property value. 
// Parameters:
//  - value : Value to set for the entitlementManagement property.
func (m *IdentityGovernance) SetEntitlementManagement(value *EntitlementManagement)() {
    m.entitlementManagement = value
}
// Sets the termsOfUse property value. 
// Parameters:
//  - value : Value to set for the termsOfUse property.
func (m *IdentityGovernance) SetTermsOfUse(value *TermsOfUseContainer)() {
    m.termsOfUse = value
}
