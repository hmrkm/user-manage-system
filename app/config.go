package main

type Config struct {
	MysqlUser         string `split_words:"true"`
	MysqlPassword     string `split_words:"true"`
	MysqlDatabase     string `split_words:"true"`
	TokenExpireHour   int    `split_words:"true"`
	AuthEndpoint      string `split_words:"true"`
	RightsEndpoint    string `split_words:"true"`
	UserManageBaseurl string `split_words:"true"`
}
