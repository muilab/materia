package mui

import (
	"github.com/aerogo/api"
	"github.com/aerogo/nano"
)

// Node represents the database node.
var Node = nano.New(5000)

// DB is the main database client.
var DB = Node.Namespace("mui").RegisterTypes(
	(*Book)(nil),
	(*EmailToUser)(nil),
	(*GoogleToUser)(nil),
	(*Material)(nil),
	(*MaterialSample)(nil),
	(*User)(nil),
	(*Session)(nil),
)

// API is the automatically created API for the database.
var API = api.New("/api/", DB)
