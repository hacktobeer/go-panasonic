module github.com/hacktobeer/gopanasonic/cloudcontrol

go 1.15

replace github.com/hacktobeer/gopanasonic/types => ./types/

require (
	github.com/google/go-cmp v0.3.0
	github.com/hacktobeer/gopanasonic/types v0.0.0-00010101000000-000000000000
	github.com/m7shapan/njson v1.0.1
	github.com/spf13/viper v1.7.1 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/tools v0.0.0-20201218024724-ae774e9781d2 // indirect
)
