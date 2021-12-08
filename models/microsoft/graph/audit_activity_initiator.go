package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuditActivityInitiator 
type AuditActivityInitiator struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If the resource initiating the activity is an app, this property indicates all the app related information like appId, Name, servicePrincipalId, Name.
    app *AppIdentity;
    // If the resource initiating the activity is a user, this property Indicates all the user related information like userId, Name, UserPrinicpalName.
    user *UserIdentity;
}
// NewAuditActivityInitiator instantiates a new auditActivityInitiator and sets the default values.
func NewAuditActivityInitiator()(*AuditActivityInitiator) {
    m := &AuditActivityInitiator{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditActivityInitiator) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApp gets the app property value. If the resource initiating the activity is an app, this property indicates all the app related information like appId, Name, servicePrincipalId, Name.
func (m *AuditActivityInitiator) GetApp()(*AppIdentity) {
    if m == nil {
        return nil
    } else {
        return m.app
    }
}
// GetUser gets the user property value. If the resource initiating the activity is a user, this property Indicates all the user related information like userId, Name, UserPrinicpalName.
func (m *AuditActivityInitiator) GetUser()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditActivityInitiator) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["app"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(*AppIdentity))
        }
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(*UserIdentity))
        }
        return nil
    }
    return res
}
func (m *AuditActivityInitiator) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AuditActivityInitiator) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetApp sets the app property value. If the resource initiating the activity is an app, this property indicates all the app related information like appId, Name, servicePrincipalId, Name.
func (m *AuditActivityInitiator) SetApp(value *AppIdentity)() {
    if m != nil {
        m.app = value
    }
}
// SetUser sets the user property value. If the resource initiating the activity is a user, this property Indicates all the user related information like userId, Name, UserPrinicpalName.
func (m *AuditActivityInitiator) SetUser(value *UserIdentity)() {
    if m != nil {
        m.user = value
    }
}
