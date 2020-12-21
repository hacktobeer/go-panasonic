module github.com/hacktobeer/go-panasonic/cloudcontrol

go 1.15

replace github.com/hacktobeer/go-panasonic => ./

replace github.com/hacktobeer/go-panasonic/types => ./types/

require (
	github.com/google/go-cmp v0.3.0
	github.com/hacktobeer/go-panasonic v1.0.0
	github.com/hacktobeer/go-panasonic/types v0.0.0-00010101000000-000000000000
	github.com/m7shapan/njson v1.0.1
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/viper v1.7.1
)
