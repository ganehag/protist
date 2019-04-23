module github.com/ganehag/protist

require (
	github.com/ganehag/protist/filter v0.0.0-00010101000000-000000000000
	github.com/ganehag/protist/plugin v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.3.2
)

replace github.com/ganehag/protist/plugin => ./plugin

replace github.com/ganehag/protist/filter => ./filter

go 1.12
