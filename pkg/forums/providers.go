package forums

import "github.com/google/wire"

// Set of providers for forums components.
var Providers = wire.NewSet(NewData, ListForums, CreateForum)
