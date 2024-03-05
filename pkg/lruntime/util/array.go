package util

import (
	lua "github.com/yuin/gopher-lua"
)

func BuildStringArray(L *lua.LState, slice []string) *lua.LTable {
	array := L.NewTable()
	for i, v := range slice {
		array.RawSetInt(i+1, lua.LString(v)) // Lua convention starts at 1
	}

	return array
}
