package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationDetails 
type UserRegistrationDetails struct {
    Entity
}
// NewUserRegistrationDetails instantiates a new userRegistrationDetails and sets the default values.
func NewUserRegistrationDetails()(*UserRegistrationDetails) {
    m := &UserRegistrationDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserRegistrationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserRegistrationDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isAdmin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAdmin(val)
        }
        return nil
    }
    res["isMfaCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaCapable(val)
        }
        return nil
    }
    res["isMfaRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaRegistered(val)
        }
        return nil
    }
    res["isPasswordlessCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordlessCapable(val)
        }
        return nil
    }
    res["isSsprCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprCapable(val)
        }
        return nil
    }
    res["isSsprEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprEnabled(val)
        }
        return nil
    }
    res["isSsprRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprRegistered(val)
        }
        return nil
    }
    res["isSystemPreferredAuthenticationMethodEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSystemPreferredAuthenticationMethodEnabled(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["methodsRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetMethodsRegistered(res)
        }
        return nil
    }
    res["systemPreferredAuthenticationMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSystemPreferredAuthenticationMethods(res)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userPreferredMethodForSecondaryAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserDefaultAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPreferredMethodForSecondaryAuthentication(val.(*UserDefaultAuthenticationMethod))
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*SignInUserType))
        }
        return nil
    }
    return res
}
// GetIsAdmin gets the isAdmin property value. The isAdmin property
func (m *UserRegistrationDetails) GetIsAdmin()(*bool) {
    val, err := m.GetBackingStore().Get("isAdmin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsMfaCapable gets the isMfaCapable property value. The isMfaCapable property
func (m *UserRegistrationDetails) GetIsMfaCapable()(*bool) {
    val, err := m.GetBackingStore().Get("isMfaCapable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsMfaRegistered gets the isMfaRegistered property value. The isMfaRegistered property
func (m *UserRegistrationDetails) GetIsMfaRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("isMfaRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPasswordlessCapable gets the isPasswordlessCapable property value. The isPasswordlessCapable property
func (m *UserRegistrationDetails) GetIsPasswordlessCapable()(*bool) {
    val, err := m.GetBackingStore().Get("isPasswordlessCapable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSsprCapable gets the isSsprCapable property value. The isSsprCapable property
func (m *UserRegistrationDetails) GetIsSsprCapable()(*bool) {
    val, err := m.GetBackingStore().Get("isSsprCapable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSsprEnabled gets the isSsprEnabled property value. The isSsprEnabled property
func (m *UserRegistrationDetails) GetIsSsprEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSsprEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSsprRegistered gets the isSsprRegistered property value. The isSsprRegistered property
func (m *UserRegistrationDetails) GetIsSsprRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("isSsprRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSystemPreferredAuthenticationMethodEnabled gets the isSystemPreferredAuthenticationMethodEnabled property value. The isSystemPreferredAuthenticationMethodEnabled property
func (m *UserRegistrationDetails) GetIsSystemPreferredAuthenticationMethodEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSystemPreferredAuthenticationMethodEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *UserRegistrationDetails) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMethodsRegistered gets the methodsRegistered property value. The methodsRegistered property
func (m *UserRegistrationDetails) GetMethodsRegistered()([]string) {
    val, err := m.GetBackingStore().Get("methodsRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSystemPreferredAuthenticationMethods gets the systemPreferredAuthenticationMethods property value. The systemPreferredAuthenticationMethods property
func (m *UserRegistrationDetails) GetSystemPreferredAuthenticationMethods()([]string) {
    val, err := m.GetBackingStore().Get("systemPreferredAuthenticationMethods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. The userDisplayName property
func (m *UserRegistrationDetails) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPreferredMethodForSecondaryAuthentication gets the userPreferredMethodForSecondaryAuthentication property value. The userPreferredMethodForSecondaryAuthentication property
func (m *UserRegistrationDetails) GetUserPreferredMethodForSecondaryAuthentication()(*UserDefaultAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("userPreferredMethodForSecondaryAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserDefaultAuthenticationMethod)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *UserRegistrationDetails) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserType gets the userType property value. The userType property
func (m *UserRegistrationDetails) GetUserType()(*SignInUserType) {
    val, err := m.GetBackingStore().Get("userType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SignInUserType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserRegistrationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isAdmin", m.GetIsAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaCapable", m.GetIsMfaCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaRegistered", m.GetIsMfaRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPasswordlessCapable", m.GetIsPasswordlessCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprCapable", m.GetIsSsprCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprEnabled", m.GetIsSsprEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprRegistered", m.GetIsSsprRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSystemPreferredAuthenticationMethodEnabled", m.GetIsSystemPreferredAuthenticationMethodEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMethodsRegistered() != nil {
        err = writer.WriteCollectionOfStringValues("methodsRegistered", m.GetMethodsRegistered())
        if err != nil {
            return err
        }
    }
    if m.GetSystemPreferredAuthenticationMethods() != nil {
        err = writer.WriteCollectionOfStringValues("systemPreferredAuthenticationMethods", m.GetSystemPreferredAuthenticationMethods())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetUserPreferredMethodForSecondaryAuthentication() != nil {
        cast := (*m.GetUserPreferredMethodForSecondaryAuthentication()).String()
        err = writer.WriteStringValue("userPreferredMethodForSecondaryAuthentication", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err = writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsAdmin sets the isAdmin property value. The isAdmin property
func (m *UserRegistrationDetails) SetIsAdmin(value *bool)() {
    err := m.GetBackingStore().Set("isAdmin", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMfaCapable sets the isMfaCapable property value. The isMfaCapable property
func (m *UserRegistrationDetails) SetIsMfaCapable(value *bool)() {
    err := m.GetBackingStore().Set("isMfaCapable", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMfaRegistered sets the isMfaRegistered property value. The isMfaRegistered property
func (m *UserRegistrationDetails) SetIsMfaRegistered(value *bool)() {
    err := m.GetBackingStore().Set("isMfaRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPasswordlessCapable sets the isPasswordlessCapable property value. The isPasswordlessCapable property
func (m *UserRegistrationDetails) SetIsPasswordlessCapable(value *bool)() {
    err := m.GetBackingStore().Set("isPasswordlessCapable", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSsprCapable sets the isSsprCapable property value. The isSsprCapable property
func (m *UserRegistrationDetails) SetIsSsprCapable(value *bool)() {
    err := m.GetBackingStore().Set("isSsprCapable", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSsprEnabled sets the isSsprEnabled property value. The isSsprEnabled property
func (m *UserRegistrationDetails) SetIsSsprEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSsprEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSsprRegistered sets the isSsprRegistered property value. The isSsprRegistered property
func (m *UserRegistrationDetails) SetIsSsprRegistered(value *bool)() {
    err := m.GetBackingStore().Set("isSsprRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSystemPreferredAuthenticationMethodEnabled sets the isSystemPreferredAuthenticationMethodEnabled property value. The isSystemPreferredAuthenticationMethodEnabled property
func (m *UserRegistrationDetails) SetIsSystemPreferredAuthenticationMethodEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSystemPreferredAuthenticationMethodEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *UserRegistrationDetails) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMethodsRegistered sets the methodsRegistered property value. The methodsRegistered property
func (m *UserRegistrationDetails) SetMethodsRegistered(value []string)() {
    err := m.GetBackingStore().Set("methodsRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPreferredAuthenticationMethods sets the systemPreferredAuthenticationMethods property value. The systemPreferredAuthenticationMethods property
func (m *UserRegistrationDetails) SetSystemPreferredAuthenticationMethods(value []string)() {
    err := m.GetBackingStore().Set("systemPreferredAuthenticationMethods", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. The userDisplayName property
func (m *UserRegistrationDetails) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPreferredMethodForSecondaryAuthentication sets the userPreferredMethodForSecondaryAuthentication property value. The userPreferredMethodForSecondaryAuthentication property
func (m *UserRegistrationDetails) SetUserPreferredMethodForSecondaryAuthentication(value *UserDefaultAuthenticationMethod)() {
    err := m.GetBackingStore().Set("userPreferredMethodForSecondaryAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *UserRegistrationDetails) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserType sets the userType property value. The userType property
func (m *UserRegistrationDetails) SetUserType(value *SignInUserType)() {
    err := m.GetBackingStore().Set("userType", value)
    if err != nil {
        panic(err)
    }
}
// UserRegistrationDetailsable 
type UserRegistrationDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsAdmin()(*bool)
    GetIsMfaCapable()(*bool)
    GetIsMfaRegistered()(*bool)
    GetIsPasswordlessCapable()(*bool)
    GetIsSsprCapable()(*bool)
    GetIsSsprEnabled()(*bool)
    GetIsSsprRegistered()(*bool)
    GetIsSystemPreferredAuthenticationMethodEnabled()(*bool)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMethodsRegistered()([]string)
    GetSystemPreferredAuthenticationMethods()([]string)
    GetUserDisplayName()(*string)
    GetUserPreferredMethodForSecondaryAuthentication()(*UserDefaultAuthenticationMethod)
    GetUserPrincipalName()(*string)
    GetUserType()(*SignInUserType)
    SetIsAdmin(value *bool)()
    SetIsMfaCapable(value *bool)()
    SetIsMfaRegistered(value *bool)()
    SetIsPasswordlessCapable(value *bool)()
    SetIsSsprCapable(value *bool)()
    SetIsSsprEnabled(value *bool)()
    SetIsSsprRegistered(value *bool)()
    SetIsSystemPreferredAuthenticationMethodEnabled(value *bool)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMethodsRegistered(value []string)()
    SetSystemPreferredAuthenticationMethods(value []string)()
    SetUserDisplayName(value *string)()
    SetUserPreferredMethodForSecondaryAuthentication(value *UserDefaultAuthenticationMethod)()
    SetUserPrincipalName(value *string)()
    SetUserType(value *SignInUserType)()
}
