module github.com/ganehag/protist/plugin/json

require (
	github.com/ganehag/protist/filter v0.0.0-20190423161215-902d358b97f0
	github.com/ganehag/protist/plugin v0.0.0-00010101000000-000000000000
)

replace github.com/ganehag/protist/plugin => ../

replace github.com/ganehag/protist/filter => ../../filter

go 1.12
