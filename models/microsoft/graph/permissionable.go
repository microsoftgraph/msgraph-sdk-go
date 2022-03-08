package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Permissionable 
type Permissionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGrantedTo()(IdentitySetable)
    GetGrantedToIdentities()([]IdentitySetable)
    GetGrantedToIdentitiesV2()([]SharePointIdentitySetable)
    GetGrantedToV2()(SharePointIdentitySetable)
    GetHasPassword()(*bool)
    GetInheritedFrom()(ItemReferenceable)
    GetInvitation()(SharingInvitationable)
    GetLink()(SharingLinkable)
    GetRoles()([]string)
    GetShareId()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGrantedTo(value IdentitySetable)()
    SetGrantedToIdentities(value []IdentitySetable)()
    SetGrantedToIdentitiesV2(value []SharePointIdentitySetable)()
    SetGrantedToV2(value SharePointIdentitySetable)()
    SetHasPassword(value *bool)()
    SetInheritedFrom(value ItemReferenceable)()
    SetInvitation(value SharingInvitationable)()
    SetLink(value SharingLinkable)()
    SetRoles(value []string)()
    SetShareId(value *string)()
}
