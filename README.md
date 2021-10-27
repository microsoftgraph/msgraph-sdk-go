# Microsoft Graph SDK for Java

[![PkgGoDev](https://pkg.go.dev/badge/github.com/microsoftgraph/msgraph-sdk-go/)](https://pkg.go.dev/github.com/microsoftgraph/msgraph-sdk-go/)

Get started with the Microsoft Graph SDK for Go by integrating the [Microsoft Graph API](https://developer.microsoft.com/en-us/graph/get-started/go) into your Go application!

> **Note:** this SDK allows you to build applications using the [v1.0](https://docs.microsoft.com/en-us/graph/use-the-api#version) of Microsoft Graph. If you want to try the latest Microsoft Graph APIs under beta, use our [beta SDK](https://github.com/microsoftgraph/msgraph-beta-sdk-go) instead.
> **Note:** the Microsoft Graph Go SDK is currently in Community Preview. During this period we're expecting breaking changes to happen to the SDK based on community's feedback. Checkout the [known limitations](https://github.com/microsoftgraph/msgraph-sdk-go-core/issues/1).

## 1. Installation

```Shell
go get github.com/microsoftgraph/msgraph-sdk-go
go get github.com/Azure/azure-sdk-for-go/sdk/azidentity
```

## 2. Getting started

### 2.1 Register your application

Register your application by following the steps at [Register your app with the Microsoft Identity Platform](https://docs.microsoft.com/graph/auth-register-app-v2).

### 2.2 Create an AuthenticationProvider object

An instance of the **GraphRequestAdapter** class handles building client. To create a new instance of this class, you need to provide an instance of `AuthenticationProvider`, which can authenticate requests to Microsoft Graph.

For an example of how to get an authentication provider, see [choose a Microsoft Graph authentication provider](https://docs.microsoft.com/graph/sdks/choose-authentication-providers?tabs=Go).

> Note: we are working to add the getting started information for Go to our public documentation, in the meantime the follwing sample should help you getting started.

```Golang


// azidentity is an import of https://github.com/Azure/azure-sdk-for-go/tree/main/sdk/azidentity

cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
        TenantID: "<the tenant id from your app registration>",
        ClientID: "<the client id from your app registration>",
        UserPrompt: func(message azidentity.DeviceCodeMessage) {
                fmt.Println(message.Message)
        },
})

if err != nil {
        fmt.Printf("Error creating credentials: %v\n", err)
}

// a is an import of https://github.com/microsoft/kiota/authentication/go/azure
auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"Mail.Read", "Mail.Send"})
if err != nil {
        fmt.Printf("Error authentication provider: %v\n", err)
        return
}

```

### 2.3 Get a Graph Service Client Adapter object

You must get a **GraphRequestAdapter** object to make requests against the service.

```Golang
// msgraphsdk is an import of this library
adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
if err != nil {
        fmt.Printf("Error creating adapter: %v\n", err)
        return
}
client := msgraphsdk.NewGraphServiceClient(adapter)
```

## 3. Make requests against the service

After you have a GraphServiceClient that is authenticated, you can begin making calls against the service. The requests against the service look like our [REST API](https://docs.microsoft.com/graph/overview).

### 3.1 Get the user's drive

To retrieve the user's drive:

```Golang
result, err := client
  .Me()
  .Drive()
  .Get(nil, nil, nil, nil)
if err != nil {
	fmt.Printf("Error getting the drive: %v\n", err)
}
fmt.Printf("Found Drive : %v\n", result.GetId())
```

## 4. Documentation

For more detailed documentation, see:

* [Overview](https://docs.microsoft.com/graph/overview)
* [Collections](https://docs.microsoft.com/graph/sdks/paging)
* [Making requests](https://docs.microsoft.com/graph/sdks/create-requests)
* [Known issues](https://github.com/MicrosoftGraph/msgraph-sdk-go/issues)
* [Contributions](https://github.com/microsoftgraph/msgraph-sdk-go/blob/main/CONTRIBUTING.md)

## 5. Issues

For known issues, see [issues](https://github.com/MicrosoftGraph/msgraph-sdk-go/issues).

## 6. Contributions

The Microsoft Graph SDK is open for contribution. To contribute to this project, see [Contributing](https://github.com/microsoftgraph/msgraph-sdk-go/blob/main/CONTRIBUTING.md).

## 7. License

Copyright (c) Microsoft Corporation. All Rights Reserved. Licensed under the [MIT license](LICENSE).

## 8. Third-party notices

[Third-party notices](THIRD%20PARTY%20NOTICES)
