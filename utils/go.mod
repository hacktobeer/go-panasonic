module cli

go 1.15

replace github.com/hacktobeer/gopanasonic/cloudcontrol => ../

replace github.com/hacktobeer/gopanasonic/types => ../types/

require (
	github.com/hacktobeer/gopanasonic/cloudcontrol v0.0.0-00010101000000-000000000000
	github.com/hacktobeer/gopanasonic/types v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.7.1
)
