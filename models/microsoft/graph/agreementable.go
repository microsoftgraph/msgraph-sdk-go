package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Agreementable 
type Agreementable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAcceptances()([]AgreementAcceptanceable)
    GetDisplayName()(*string)
    GetFile()(AgreementFileable)
    GetFiles()([]AgreementFileLocalizationable)
    GetIsPerDeviceAcceptanceRequired()(*bool)
    GetIsViewingBeforeAcceptanceRequired()(*bool)
    GetTermsExpiration()(TermsExpirationable)
    GetUserReacceptRequiredFrequency()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    SetAcceptances(value []AgreementAcceptanceable)()
    SetDisplayName(value *string)()
    SetFile(value AgreementFileable)()
    SetFiles(value []AgreementFileLocalizationable)()
    SetIsPerDeviceAcceptanceRequired(value *bool)()
    SetIsViewingBeforeAcceptanceRequired(value *bool)()
    SetTermsExpiration(value TermsExpirationable)()
    SetUserReacceptRequiredFrequency(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
}
