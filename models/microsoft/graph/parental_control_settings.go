package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ParentalControlSettings 
type ParentalControlSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list.
    countriesBlockedForMinors []string;
    // Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
    legalAgeGroupRule *string;
}
// NewParentalControlSettings instantiates a new parentalControlSettings and sets the default values.
func NewParentalControlSettings()(*ParentalControlSettings) {
    m := &ParentalControlSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParentalControlSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCountriesBlockedForMinors gets the countriesBlockedForMinors property value. Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list.
func (m *ParentalControlSettings) GetCountriesBlockedForMinors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.countriesBlockedForMinors
    }
}
// GetLegalAgeGroupRule gets the legalAgeGroupRule property value. Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
func (m *ParentalControlSettings) GetLegalAgeGroupRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.legalAgeGroupRule
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParentalControlSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["countriesBlockedForMinors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCountriesBlockedForMinors(res)
        }
        return nil
    }
    res["legalAgeGroupRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegalAgeGroupRule(val)
        }
        return nil
    }
    return res
}
func (m *ParentalControlSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ParentalControlSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("countriesBlockedForMinors", m.GetCountriesBlockedForMinors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("legalAgeGroupRule", m.GetLegalAgeGroupRule())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParentalControlSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCountriesBlockedForMinors sets the countriesBlockedForMinors property value. Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list.
func (m *ParentalControlSettings) SetCountriesBlockedForMinors(value []string)() {
    if m != nil {
        m.countriesBlockedForMinors = value
    }
}
// SetLegalAgeGroupRule sets the legalAgeGroupRule property value. Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
func (m *ParentalControlSettings) SetLegalAgeGroupRule(value *string)() {
    if m != nil {
        m.legalAgeGroupRule = value
    }
}
