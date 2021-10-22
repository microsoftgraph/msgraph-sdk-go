package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserScopeTeamsAppInstallation struct {
    TeamsAppInstallation
    chat *Chat;
}
func NewUserScopeTeamsAppInstallation()(*UserScopeTeamsAppInstallation) {
    m := &UserScopeTeamsAppInstallation{
        TeamsAppInstallation: *NewTeamsAppInstallation(),
    }
    return m
}
func (m *UserScopeTeamsAppInstallation) GetChat()(*Chat) {
    if m == nil {
        return nil
    } else {
        return m.chat
    }
}
func (m *UserScopeTeamsAppInstallation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TeamsAppInstallation.GetFieldDeserializers()
    res["chat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChat() })
        if err != nil {
            return err
        }
        m.SetChat(val.(*Chat))
        return nil
    }
    return res
}
func (m *UserScopeTeamsAppInstallation) IsNil()(bool) {
    return m == nil
}
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
func (m *UserScopeTeamsAppInstallation) SetChat(value *Chat)() {
    m.chat = value
}
