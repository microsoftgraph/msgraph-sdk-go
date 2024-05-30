package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\appManagementPolicies\{appManagementPolicy-id}
type ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilderInternal instantiates a new ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder and sets the default values.
func NewItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder) {
    m := &ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/appManagementPolicies/{appManagementPolicy%2Did}", pathParameters),
    }
    return m
}
// NewItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder instantiates a new ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder and sets the default values.
func NewItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of application entities.
// returns a *ItemAppmanagementpoliciesItemRefRequestBuilder when successful
func (m *ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder) Ref()(*ItemAppmanagementpoliciesItemRefRequestBuilder) {
    return NewItemAppmanagementpoliciesItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
