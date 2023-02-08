package setup

import (
	"benchmarks/env"
	"fmt"

	"github.com/tarantool/go-tarantool"
)

func Tarantool() error {
	conn, err := tarantool.Connect(fmt.Sprintf("%s:%s", env.TntHost, env.TntPort), tarantool.Opts{
		User: env.TntUser,
		Pass: env.TntPassword,
	})
	if err != nil {
		return err
	}
	// Space Test
	// | id | name  |
	// |----|-------|
	// | 1  | Alice |
	// | 2  | Bob   |
	_, err = conn.Call("box.schema.space.create", []interface{}{
		"Test",
		map[string]bool{"if_not_exists": true}})
	// if err != nil {
	// 	return fmt.Errorf("unexpected error for conn.Call(space.create): %v", err)
	// }
	_, err = conn.Call("box.space.Test:format", [][]map[string]string{
		{
			{"name": "id", "type": "number"},
			{"name": "name1", "type": "string"},
			{"name": "name2", "type": "string"},
			{"name": "name3", "type": "string"},
			{"name": "name4", "type": "string"},
		}})
	if err != nil {
		return fmt.Errorf("unexpected error for conn.Call(space.Test:format): %v", err)
	}
	_, err = conn.Call("box.space.Test:create_index", []interface{}{
		"primary",
		map[string]interface{}{
			"parts":         []string{"id"},
			"if_not_exists": true}})
	if err != nil {
		return fmt.Errorf("unexpected error for conn.Call(space.Test:create_index): %v", err)
	}
	return conn.Close()
}
