// Code generated by 'micro gen' command.
// DO NOT EDIT!

package api

import (
	tp "github.com/henrylee2cn/teleport/v6"
)

// Route registers handlers to router.
func Route(root string, router *tp.Router) {
	// root router group
	group := router.SubRoute(root)

	// custom router
	customRoute(group.ToRouter())

	// automatically generated router

	// CALL APIs...
	group.RouteCallFunc(IsGray)
	group.RouteCallFunc(Get)
	group.RouteCallFunc(Delete)
	group.RouteCallFunc(Set)
}
