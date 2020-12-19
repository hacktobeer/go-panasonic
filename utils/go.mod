module cli

replace github.com/hacktobeer/go-panasonic/cloudcontrol => ../

replace github.com/hacktobeer/go-panasonic/types => ../types/

go 1.15

require (
	github.com/hacktobeer/go-panasonic/cloudcontrol v0.0.0-00010101000000-000000000000
	github.com/hacktobeer/go-panasonic/types v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.7.1
)
