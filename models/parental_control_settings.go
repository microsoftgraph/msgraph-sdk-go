package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParentalControlSettings 
type ParentalControlSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list.
    countriesBlockedForMinors []string
    // Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
    legalAgeGroupRule *string
}
// NewParentalControlSettings instantiates a new parentalControlSettings and sets the default values.
func NewParentalControlSettings()(*ParentalControlSettings) {
    m := &ParentalControlSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateParentalControlSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParentalControlSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParentalControlSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
// GetFieldDeserializers the deserialization information for the current model
func (m *ParentalControlSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["countriesBlockedForMinors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["legalAgeGroupRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLegalAgeGroupRule gets the legalAgeGroupRule property value. Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
func (m *ParentalControlSettings) GetLegalAgeGroupRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.legalAgeGroupRule
    }
}
// Serialize serializes information the current object
func (m *ParentalControlSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCountriesBlockedForMinors() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
