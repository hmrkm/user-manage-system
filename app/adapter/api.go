// Package adapter provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package adapter

// ページャー
type Page struct {
	// 最終位置
	Last int `json:"last"`

	// 現在位置
	Now int `json:"now"`
}

// 認証リクエストボディ
type RequestAuth struct {
	// メールアドレス
	Email string `json:"email"`

	// パスワード
	Password string `json:"password"`
}

// ユーザー新規作成リクエスト
type RequestUsersCreate struct {
	// メールアドレス
	Email string `json:"email"`

	// パスワード
	Password string `json:"password"`

	// 確認用パスワード
	PasswordConf string `json:"password_conf"`
}

// ユーザー削除リクエスト
type RequestUsersDelete struct {
	// ユーザーID
	Id string `json:"id"`
}

// ユーザー詳細情報取得リクエスト
type RequestUsersDetail struct {
	// ユーザーID
	Id string `json:"id"`
}

// ユーザーリスト取得リクエスト
type RequestUsersList struct {
	// 件数
	Limit int `json:"limit"`

	// ページャー
	Page int `json:"page"`
}

// ユーザー詳細更新リクエスト
type RequestUsersUpdate struct {
	// メールアドレス
	Email string `json:"email"`

	// ユーザーID
	Id string `json:"id"`

	// 更新後パスワード
	Password *string `json:"password,omitempty"`

	// 確認用パスワード
	PasswordConf *string `json:"password_conf,omitempty"`
}

// ユーザー
type User struct {
	// メールアドレス
	Email string `json:"email"`

	// ユーザーID
	Id string `json:"id"`
}

// AuthToken defines model for auth_token.
type AuthToken string

// ResponseAuth defines model for ResponseAuth.
type ResponseAuth struct {
	// 認証トークンの有効期限Unixtimeミリ秒
	ExpiredAt int `json:"expired_at"`

	// 認証トークン
	Token string `json:"token"`
}

// ユーザー
type ResponseUsersDetail User

// ResponseUsersList defines model for ResponseUsersList.
type ResponseUsersList struct {
	// ユーザーリスト
	List []User `json:"list"`

	// ページャー
	Page Page `json:"page"`
}

// PostV1AuthJSONBody defines parameters for PostV1Auth.
type PostV1AuthJSONBody RequestAuth

// PostV1UsersCreateJSONBody defines parameters for PostV1UsersCreate.
type PostV1UsersCreateJSONBody RequestUsersCreate

// PostV1UsersCreateParams defines parameters for PostV1UsersCreate.
type PostV1UsersCreateParams struct {
	// 認証トークン
	AuthToken *AuthToken `json:"auth_token,omitempty"`
}

// PostV1UsersDeleteJSONBody defines parameters for PostV1UsersDelete.
type PostV1UsersDeleteJSONBody RequestUsersDelete

// PostV1UsersDeleteParams defines parameters for PostV1UsersDelete.
type PostV1UsersDeleteParams struct {
	// 認証トークン
	AuthToken *AuthToken `json:"auth_token,omitempty"`
}

// PostV1UsersDetailJSONBody defines parameters for PostV1UsersDetail.
type PostV1UsersDetailJSONBody RequestUsersDetail

// PostV1UsersDetailParams defines parameters for PostV1UsersDetail.
type PostV1UsersDetailParams struct {
	// 認証トークン
	AuthToken *AuthToken `json:"auth_token,omitempty"`
}

// PostV1UsersListJSONBody defines parameters for PostV1UsersList.
type PostV1UsersListJSONBody RequestUsersList

// PostV1UsersListParams defines parameters for PostV1UsersList.
type PostV1UsersListParams struct {
	// 認証トークン
	AuthToken *AuthToken `json:"auth_token,omitempty"`
}

// PostV1UsersUpdateJSONBody defines parameters for PostV1UsersUpdate.
type PostV1UsersUpdateJSONBody RequestUsersUpdate

// PostV1UsersUpdateParams defines parameters for PostV1UsersUpdate.
type PostV1UsersUpdateParams struct {
	// 認証トークン
	AuthToken *AuthToken `json:"auth_token,omitempty"`
}

// PostV1AuthJSONRequestBody defines body for PostV1Auth for application/json ContentType.
type PostV1AuthJSONRequestBody PostV1AuthJSONBody

// PostV1UsersCreateJSONRequestBody defines body for PostV1UsersCreate for application/json ContentType.
type PostV1UsersCreateJSONRequestBody PostV1UsersCreateJSONBody

// PostV1UsersDeleteJSONRequestBody defines body for PostV1UsersDelete for application/json ContentType.
type PostV1UsersDeleteJSONRequestBody PostV1UsersDeleteJSONBody

// PostV1UsersDetailJSONRequestBody defines body for PostV1UsersDetail for application/json ContentType.
type PostV1UsersDetailJSONRequestBody PostV1UsersDetailJSONBody

// PostV1UsersListJSONRequestBody defines body for PostV1UsersList for application/json ContentType.
type PostV1UsersListJSONRequestBody PostV1UsersListJSONBody

// PostV1UsersUpdateJSONRequestBody defines body for PostV1UsersUpdate for application/json ContentType.
type PostV1UsersUpdateJSONRequestBody PostV1UsersUpdateJSONBody

