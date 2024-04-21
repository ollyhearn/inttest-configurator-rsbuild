# IntTestConfigurator.UsersApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**auth**](UsersApi.md#auth) | **POST** /auth | create user
[**createRole**](UsersApi.md#createRole) | **POST** /roles | create role
[**createUser**](UsersApi.md#createUser) | **POST** /users | create user
[**deleteRole**](UsersApi.md#deleteRole) | **DELETE** /roles/{id} | delete role
[**deleteUser**](UsersApi.md#deleteUser) | **DELETE** /users/{id} | delete user
[**listRoles**](UsersApi.md#listRoles) | **GET** /roles | list all the roles in the system
[**listUsers**](UsersApi.md#listUsers) | **GET** /users | create user
[**updateRole**](UsersApi.md#updateRole) | **PUT** /roles/{id} | update role



## auth

> Object auth(form)

create user

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
let form = new IntTestConfigurator.ConfiguratorInternalApiAuthAuthRequest(); // ConfiguratorInternalApiAuthAuthRequest | login/pass form
apiInstance.auth(form, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **form** | [**ConfiguratorInternalApiAuthAuthRequest**](ConfiguratorInternalApiAuthAuthRequest.md)| login/pass form | 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createRole

> ConfiguratorInternalApiAuthRoleCreateResponse createRole(form)

create role

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
let form = new IntTestConfigurator.ConfiguratorInternalApiAuthRoleCreateRequest(); // ConfiguratorInternalApiAuthRoleCreateRequest | create role form
apiInstance.createRole(form, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **form** | [**ConfiguratorInternalApiAuthRoleCreateRequest**](ConfiguratorInternalApiAuthRoleCreateRequest.md)| create role form | 

### Return type

[**ConfiguratorInternalApiAuthRoleCreateResponse**](ConfiguratorInternalApiAuthRoleCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createUser

> ConfiguratorInternalApiAuthCreateUserResponse createUser(form)

create user

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
let form = new IntTestConfigurator.ConfiguratorInternalApiAuthCreateUserRequest(); // ConfiguratorInternalApiAuthCreateUserRequest | create user request model
apiInstance.createUser(form, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **form** | [**ConfiguratorInternalApiAuthCreateUserRequest**](ConfiguratorInternalApiAuthCreateUserRequest.md)| create user request model | 

### Return type

[**ConfiguratorInternalApiAuthCreateUserResponse**](ConfiguratorInternalApiAuthCreateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteRole

> Object deleteRole(id)

delete role

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
let id = 3.4; // Number | id of a role to delete
apiInstance.deleteRole(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| id of a role to delete | 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteUser

> Object deleteUser(id)

delete user

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
let id = 3.4; // Number | id of a user to delete
apiInstance.deleteUser(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| id of a user to delete | 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listRoles

> [ConfiguratorInternalApiAuthListRoleResponseItem] listRoles()

list all the roles in the system

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
apiInstance.listRoles((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**[ConfiguratorInternalApiAuthListRoleResponseItem]**](ConfiguratorInternalApiAuthListRoleResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listUsers

> [ConfiguratorInternalApiAuthListUsersResponseItem] listUsers()

create user

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
apiInstance.listUsers((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**[ConfiguratorInternalApiAuthListUsersResponseItem]**](ConfiguratorInternalApiAuthListUsersResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateRole

> Object updateRole(id, form)

update role

### Example

```javascript
import IntTestConfigurator from 'int_test_configurator';

let apiInstance = new IntTestConfigurator.UsersApi();
let id = 3.4; // Number | id of a role to update
let form = new IntTestConfigurator.ConfiguratorInternalApiAuthUpdateRoleRequest(); // ConfiguratorInternalApiAuthUpdateRoleRequest | new role data to store
apiInstance.updateRole(id, form, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| id of a role to update | 
 **form** | [**ConfiguratorInternalApiAuthUpdateRoleRequest**](ConfiguratorInternalApiAuthUpdateRoleRequest.md)| new role data to store | 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

