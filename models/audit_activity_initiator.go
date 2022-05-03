package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditActivityInitiator 
type AuditActivityInitiator struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If the actor initiating the activity is an app, this property indicates all its identification information including appId, displayName, servicePrincipalId, and servicePrincipalName.
    app AppIdentityable
    // If the actor initiating the activity is a user, this property indicates their identification information including their id, displayName, and userPrincipalName.
    user UserIdentityable
}
// NewAuditActivityInitiator instantiates a new auditActivityInitiator and sets the default values.
func NewAuditActivityInitiator()(*AuditActivityInitiator) {
    m := &AuditActivityInitiator{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuditActivityInitiatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditActivityInitiatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditActivityInitiator(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditActivityInitiator) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApp gets the app property value. If the actor initiating the activity is an app, this property indicates all its identification information including appId, displayName, servicePrincipalId, and servicePrincipalName.
func (m *AuditActivityInitiator) GetApp()(AppIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.app
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditActivityInitiator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(AppIdentityable))
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(UserIdentityable))
        }
        return nil
    }
    return res
}
// GetUser gets the user property value. If the actor initiating the activity is a user, this property indicates their identification information including their id, displayName, and userPrincipalName.
func (m *AuditActivityInitiator) GetUser()(UserIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// Serialize serializes information the current object
func (m *AuditActivityInitiator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *AuditActivityInitiator) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApp sets the app property value. If the actor initiating the activity is an app, this property indicates all its identification information including appId, displayName, servicePrincipalId, and servicePrincipalName.
func (m *AuditActivityInitiator) SetApp(value AppIdentityable)() {
    if m != nil {
        m.app = value
    }
}
// SetUser sets the user property value. If the actor initiating the activity is a user, this property indicates their identification information including their id, displayName, and userPrincipalName.
func (m *AuditActivityInitiator) SetUser(value UserIdentityable)() {
    if m != nil {
        m.user = value
    }
}
