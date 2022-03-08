package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessagePolicyViolationable 
type ChatMessagePolicyViolationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDlpAction()(*ChatMessagePolicyViolationDlpActionTypes)
    GetJustificationText()(*string)
    GetPolicyTip()(ChatMessagePolicyViolationPolicyTipable)
    GetUserAction()(*ChatMessagePolicyViolationUserActionTypes)
    GetVerdictDetails()(*ChatMessagePolicyViolationVerdictDetailsTypes)
    SetDlpAction(value *ChatMessagePolicyViolationDlpActionTypes)()
    SetJustificationText(value *string)()
    SetPolicyTip(value ChatMessagePolicyViolationPolicyTipable)()
    SetUserAction(value *ChatMessagePolicyViolationUserActionTypes)()
    SetVerdictDetails(value *ChatMessagePolicyViolationVerdictDetailsTypes)()
}
