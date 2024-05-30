package identity

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ConditionalaccessConditionalAccessRequestBuilder builds and executes requests for operations under \identity\conditionalAccess
type ConditionalaccessConditionalAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationContextClassReferences provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
// returns a *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferencesRequestBuilder when successful
func (m *ConditionalaccessConditionalAccessRequestBuilder) AuthenticationContextClassReferences()(*ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferencesRequestBuilder) {
    return NewConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuthenticationStrength provides operations to manage the authenticationStrength property of the microsoft.graph.conditionalAccessRoot entity.
// returns a *ConditionalaccessAuthenticationstrengthAuthenticationStrengthRequestBuilder when successful
func (m *ConditionalaccessConditionalAccessRequestBuilder) AuthenticationStrength()(*ConditionalaccessAuthenticationstrengthAuthenticationStrengthRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthAuthenticationStrengthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalaccessConditionalAccessRequestBuilderInternal instantiates a new ConditionalaccessConditionalAccessRequestBuilder and sets the default values.
func NewConditionalaccessConditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessConditionalAccessRequestBuilder) {
    m := &ConditionalaccessConditionalAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess", pathParameters),
    }
    return m
}
// NewConditionalaccessConditionalAccessRequestBuilder instantiates a new ConditionalaccessConditionalAccessRequestBuilder and sets the default values.
func NewConditionalaccessConditionalAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessConditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessConditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// NamedLocations provides operations to manage the namedLocations property of the microsoft.graph.conditionalAccessRoot entity.
// returns a *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder when successful
func (m *ConditionalaccessConditionalAccessRequestBuilder) NamedLocations()(*ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) {
    return NewConditionalaccessNamedlocationsNamedLocationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Policies provides operations to manage the policies property of the microsoft.graph.conditionalAccessRoot entity.
// returns a *ConditionalaccessPoliciesRequestBuilder when successful
func (m *ConditionalaccessConditionalAccessRequestBuilder) Policies()(*ConditionalaccessPoliciesRequestBuilder) {
    return NewConditionalaccessPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Templates provides operations to manage the templates property of the microsoft.graph.conditionalAccessRoot entity.
// returns a *ConditionalaccessTemplatesRequestBuilder when successful
func (m *ConditionalaccessConditionalAccessRequestBuilder) Templates()(*ConditionalaccessTemplatesRequestBuilder) {
    return NewConditionalaccessTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
