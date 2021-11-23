package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// userScopeTeamsAppInstallation 
type UserScopeTeamsAppInstallation struct {
    TeamsAppInstallation
    // The chat between the user and Teams app.
    chat *Chat;
}
// NewUserScopeTeamsAppInstallation instantiates a new userScopeTeamsAppInstallation and sets the default values.
func NewUserScopeTeamsAppInstallation()(*UserScopeTeamsAppInstallation) {
    m := &UserScopeTeamsAppInstallation{
        TeamsAppInstallation: *NewTeamsAppInstallation(),
    }
    return m
}
// GetChat gets the chat property value. The chat between the user and Teams app.
func (m *UserScopeTeamsAppInstallation) GetChat()(*Chat) {
    if m == nil {
        return nil
    } else {
        return m.chat
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserScopeTeamsAppInstallation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TeamsAppInstallation.GetFieldDeserializers()
    res["chat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChat(val.(*Chat))
        }
        return nil
    }
    return res
}
func (m *UserScopeTeamsAppInstallation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserScopeTeamsAppInstallation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.TeamsAppInstallation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("chat", m.GetChat())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChat sets the chat property value. The chat between the user and Teams app.
func (m *UserScopeTeamsAppInstallation) SetChat(value *Chat)() {
    m.chat = value
}
