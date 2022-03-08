package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Subscriptionable 
type Subscriptionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationId()(*string)
    GetChangeType()(*string)
    GetClientState()(*string)
    GetCreatorId()(*string)
    GetEncryptionCertificate()(*string)
    GetEncryptionCertificateId()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIncludeResourceData()(*bool)
    GetLatestSupportedTlsVersion()(*string)
    GetLifecycleNotificationUrl()(*string)
    GetNotificationQueryOptions()(*string)
    GetNotificationUrl()(*string)
    GetNotificationUrlAppId()(*string)
    GetResource()(*string)
    SetApplicationId(value *string)()
    SetChangeType(value *string)()
    SetClientState(value *string)()
    SetCreatorId(value *string)()
    SetEncryptionCertificate(value *string)()
    SetEncryptionCertificateId(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIncludeResourceData(value *bool)()
    SetLatestSupportedTlsVersion(value *string)()
    SetLifecycleNotificationUrl(value *string)()
    SetNotificationQueryOptions(value *string)()
    SetNotificationUrl(value *string)()
    SetNotificationUrlAppId(value *string)()
    SetResource(value *string)()
}
