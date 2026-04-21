// Package steve is a stand-in module for the automation playground.
// Imports wrangler so go mod tidy keeps the require entry.
package steve

import "github.com/tomleb/frameworks-automation-wrangler/pkg/wrangler"

const Version = "0.8"

func Hello() string { return "hello from steve " + Version + " using " + wrangler.Hello() }
