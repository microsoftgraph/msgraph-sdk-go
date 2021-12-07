package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityGovernance 
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
// NewIdentityGovernance instantiates a new IdentityGovernance and sets the default values.
func NewIdentityGovernance()(*IdentityGovernance) {
    m := &IdentityGovernance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessReviews gets the accessReviews property value. 
func (m *IdentityGovernance) GetAccessReviews()(*AccessReviewSet) {
    if m == nil {
        return nil
    } else {
        return m.accessReviews
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppConsent gets the appConsent property value. 
func (m *IdentityGovernance) GetAppConsent()(*AppConsentApprovalRoute) {
    if m == nil {
        return nil
    } else {
        return m.appConsent
    }
}
// GetEntitlementManagement gets the entitlementManagement property value. 
func (m *IdentityGovernance) GetEntitlementManagement()(*EntitlementManagement) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
// GetTermsOfUse gets the termsOfUse property value. 
func (m *IdentityGovernance) GetTermsOfUse()(*TermsOfUseContainer) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUse
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAccessReviews sets the accessReviews property value. 
func (m *IdentityGovernance) SetAccessReviews(value *AccessReviewSet)() {
    if m != nil {
        m.accessReviews = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppConsent sets the appConsent property value. 
func (m *IdentityGovernance) SetAppConsent(value *AppConsentApprovalRoute)() {
    if m != nil {
        m.appConsent = value
    }
}
// SetEntitlementManagement sets the entitlementManagement property value. 
func (m *IdentityGovernance) SetEntitlementManagement(value *EntitlementManagement)() {
    if m != nil {
        m.entitlementManagement = value
    }
}
// SetTermsOfUse sets the termsOfUse property value. 
func (m *IdentityGovernance) SetTermsOfUse(value *TermsOfUseContainer)() {
    if m != nil {
        m.termsOfUse = value
    }
}
