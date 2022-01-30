package main

type Config struct {
	AuthEndpoint      string `split_words:"true"`
	VerifyEndpoint      string `split_words:"true"`
	RightsEndpoint    string `split_words:"true"`
	UserManageBaseurl string `split_words:"true"`
}
