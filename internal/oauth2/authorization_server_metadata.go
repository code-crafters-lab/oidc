package oauth2

const (
	WellKnownEndpoint                   = "/.well-known"
	AuthorizationServerMetadataEndpoint = WellKnownEndpoint + "/oauth-authorization-server"
)

type ClientType string

const (
	ConfidentialClient = "confidential"
	PublicClient       = "public"
)

type AuthMethod string

const (
	AuthMethodBasic         AuthMethod = "client_secret_basic"
	AuthMethodPost          AuthMethod = "client_secret_post"
	AuthMethodNone          AuthMethod = "none"
	AuthMethodPrivateKeyJWT AuthMethod = "private_key_jwt"
)

var AllAuthMethods = []AuthMethod{
	AuthMethodBasic, AuthMethodPost, AuthMethodNone, AuthMethodPrivateKeyJWT,
}

// AuthorizationServerMetadata 表示 OAuth 2.0 授权服务器元数据的结构体。
// https://datatracker.ietf.org/doc/html/rfc8414#section-2
type AuthorizationServerMetadata struct {
	// REQUIRED. 授权服务器的发布者标识符，是一个使用“https”方案且没有查询或片段组件的 URL。
	Issuer string `json:"issuer"`
	// URL of the authorization server's authorization endpoint. 除非不支持使用授权端点的授权类型，否则此为必填项。
	AuthorizationEndpoint *string `json:"authorization_endpoint,omitempty"`
	// URL of the authorization server's token endpoint. 除非仅支持隐式授权类型，否则此为必填项。
	TokenEndpoint *string `json:"token_endpoint,omitempty"`
	// OPTIONAL. 授权服务器的 JWK Set 文档的 URL。此 URL 必须使用“https”方案。
	JwksURI *string `json:"jwks_uri,omitempty"`
	// OPTIONAL. 授权服务器的 OAuth 2.0 动态客户端注册端点的 URL。
	RegistrationEndpoint *string `json:"registration_endpoint,omitempty"`
	// RECOMMENDED. JSON 数组，包含此授权服务器支持的 OAuth 2.0 “scope”值列表。
	ScopesSupported []string `json:"scopes_supported,omitempty"`
	// REQUIRED. JSON 数组，包含此授权服务器支持的 OAuth 2.0 “response_type”值列表。
	ResponseTypesSupported []string `json:"response_types_supported"`
	// OPTIONAL. JSON 数组，包含此授权服务器支持的 OAuth 2.0 “response_mode”值列表。如果省略，默认值为["query", "fragment"]。
	ResponseModesSupported []string `json:"response_modes_supported,omitempty"`
	// OPTIONAL. JSON 数组，包含此授权服务器支持的 OAuth 2.0 授权类型值列表。如果省略，默认值为["authorization_code", "implicit"]。
	GrantTypesSupported []string `json:"grant_types_supported,omitempty"`
	// OPTIONAL. JSON 数组，包含此令牌端点支持的客户端身份验证方法列表。
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	// OPTIONAL. JSON 数组，包含此令牌端点为用于在“private_key_jwt”和“client_secret_jwt”身份验证方法中对 JWT 进行签名而支持的 JWS 签名算法（“alg”值）列表。
	TokenEndpointAuthSigningAlgValues []string `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`
	// OPTIONAL. 包含开发人员在使用授权服务器时可能需要了解的人类可读信息的页面的 URL。
	ServiceDocumentation *string `json:"service_documentation,omitempty"`
	// OPTIONAL. 以 BCP 47 语言标签值表示的用户界面支持的语言和脚本的 JSON 数组。如果省略，支持的语言和脚本集未指定。
	UILocalesSupported []string `json:"ui_locales_supported,omitempty"`
	// OPTIONAL. 授权服务器提供给注册客户端的人员以阅读授权服务器关于客户端如何使用授权服务器提供的数据的要求的 URL。
	OpPolicyURI *string `json:"op_policy_uri,omitempty"`
	// OPTIONAL. 授权服务器提供给注册客户端的人员以阅读授权服务器的服务条款的 URL。
	OpTosURI *string `json:"op_tos_uri,omitempty"`
	// OPTIONAL. 授权服务器的 OAuth 2.0 撤销端点的 URL。
	RevocationEndpoint *string `json:"revocation_endpoint,omitempty"`
	// OPTIONAL. JSON 数组，包含此撤销端点支持的客户端身份验证方法列表。如果省略，默认值为“client_secret_basic”。
	RevocationEndpointAuthMethods []string `json:"revocation_endpoint_auth_methods_supported,omitempty"`
	// OPTIONAL. JSON 数组，包含此撤销端点为用于在“private_key_jwt”和“client_secret_jwt”身份验证方法中对 JWT 进行签名而支持的 JWS 签名算法（“alg”值）列表。
	RevocationEndpointAuthSigningAlg []string `json:"revocation_endpoint_auth_signing_alg_values_supported,omitempty"`
	// OPTIONAL. 授权服务器的 OAuth 2.0 内省端点的 URL。
	IntrospectionEndpoint *string `json:"introspection_endpoint,omitempty"`
	// OPTIONAL. JSON 数组，包含此内省端点支持的客户端身份验证方法列表。
	IntrospectionEndpointAuthMethods []string `json:"introspection_endpoint_auth_methods_supported,omitempty"`
	// OPTIONAL. JSON 数组，包含此内省端点为用于在“private_key_jwt”和“client_secret_jwt”身份验证方法中对 JWT 进行签名而支持的 JWS 签名算法（“alg”值）列表。
	IntrospectionEndpointAuthSigningAlg []string `json:"introspection_endpoint_auth_signing_alg_values_supported,omitempty"`
	// OPTIONAL. JSON 数组，包含此授权服务器支持的 PKCE 代码挑战方法列表。如果省略，授权服务器不支持 PKCE。
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported,omitempty"`
	// OPTIONAL. 包含授权服务器元数据值作为声明的 JWT。这是一个字符串值，由整个签名的 JWT 组成。“signed_metadata”元数据值不应作为 JWT 中的声明出现。
	SignedMetadata *string `json:"signed_metadata,omitempty"`
}

// SetIssuer 设置 Issuer。
func (m *AuthorizationServerMetadata) SetIssuer(issuer string) {
	m.Issuer = issuer
}

// SetAuthorizationEndpoint 设置 AuthorizationEndpoint。
func (m *AuthorizationServerMetadata) SetAuthorizationEndpoint(endpoint string) {
	m.AuthorizationEndpoint = &endpoint
}

// SetTokenEndpoint 设置 TokenEndpoint。
func (m *AuthorizationServerMetadata) SetTokenEndpoint(endpoint string) {
	m.TokenEndpoint = &endpoint
}

// SetJwksURI 设置 JwksURI。
func (m *AuthorizationServerMetadata) SetJwksURI(uri string) {
	m.JwksURI = &uri
}

// SetRegistrationEndpoint 设置 RegistrationEndpoint。
func (m *AuthorizationServerMetadata) SetRegistrationEndpoint(endpoint string) {
	m.RegistrationEndpoint = &endpoint
}

// SetServiceDocumentation 设置 ServiceDocumentation。
func (m *AuthorizationServerMetadata) SetServiceDocumentation(uri string) {
	m.ServiceDocumentation = &uri
}

// SetOpPolicyURI 设置 OpPolicyURI。
func (m *AuthorizationServerMetadata) SetOpPolicyURI(uri string) {
	m.OpPolicyURI = &uri
}

// SetOpTosURI 设置 OpTosURI。
func (m *AuthorizationServerMetadata) SetOpTosURI(uri string) {
	m.OpTosURI = &uri
}

// SetRevocationEndpoint 设置 RevocationEndpoint。
func (m *AuthorizationServerMetadata) SetRevocationEndpoint(endpoint string) {
	m.RevocationEndpoint = &endpoint
}

// SetIntrospectionEndpoint 设置 IntrospectionEndpoint。
func (m *AuthorizationServerMetadata) SetIntrospectionEndpoint(endpoint string) {
	m.IntrospectionEndpoint = &endpoint
}

// SetSignedMetadata 设置 SignedMetadata。
func (m *AuthorizationServerMetadata) SetSignedMetadata(metadata string) {
	m.SignedMetadata = &metadata
}
