package main

import (
	"fmt"
	"sync"
)

// handlers for setting accepted commands (https://redis.io/docs/latest/commands/)

var Handlers = map[string]func([]Value) Value{
	"APPEND":  append_set,
	"COMMAND": command,
	// "DECR": decr,
	// "DECRBY": decrby,
	"DEL":    del,
	"GET":    get,
	"GETDEL": getdel,
	// "GETRANGE": getrange,
	// "ECHO":    echo,
	"EXISTS": exists,
	// "HDEL":    hdel,
	// "HEXISTS": hexists,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
	// "HINCRBY": hincrby,
	// "HKEYS":   hkeys,
	// "HSTRLEN": hstrlen,
	// "INCR":    incr,
	// "INCRBY":  incrby,
	// "KEYS": keys,
	// "RENAME": rename,
	"SET":  set,
	"PING": ping,
}

func command(args []Value) Value {
	return Value{typ: "string", str: "COMMAND RESPONSE"}
}

func ping(args []Value) Value {
	if len(args) > 0 {
		return Value{typ: "error", str: "PING takes no arguments"}
	}

	return Value{typ: "string", str: "PONG"}
}

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "SET requires 2 arguments (key and value)"}
	}

	key := args[0].bulk
	value := args[1].bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func append_set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "SET requires 2 arguments (key and value)"}
	}

	key := args[0].bulk
	value := args[1].bulk

	SETsMu.Lock()
	SETs[key] = SETs[key] + value
	SETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "GET requires 1 argument (key)"}
	}

	key := args[0].bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: value}
}

func getdel(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "GET requires 1 argument (key)"}
	}

	key := args[0].bulk
	var value string

	SETsMu.Lock()
	if val, ok := SETs[key]; ok {
		value = val
		delete(SETs, key)
	} else {
		return Value{typ: "null"}
	}
	SETsMu.Unlock()

	return Value{typ: "string", str: value}
}

func del(args []Value) Value {
	var count int

	SETsMu.Lock()
	for i := 0; i < len(args); i++ {
		key := args[i].bulk
		if _, ok := SETs[key]; ok {
			delete(SETs, key)
			count++
		}
	}
	SETsMu.Unlock()

	return Value{typ: "integer", int: count}
}

func exists(args []Value) Value {
	var count int

	SETsMu.RLock()
	for i := 0; i < len(args); i++ {
		if _, ok := SETs[args[i].bulk]; ok {
			count++
		}
	}
	SETsMu.RUnlock()

	return Value{typ: "integer", int: count}
}

var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "HSET requires 3 arguments (hash, key, value)"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "HGET requires 2 arguments (hash and key)"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return Value{typ: "error", str: fmt.Sprintf("'%s'/'%s' not found", hash, key)}
	}

	return Value{typ: "bulk", bulk: value}
}

func hgetall(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "HGETALL requires 1 argument (hash)"}
	}

	hash := args[0].bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash]
	HSETsMu.RUnlock()

	if !ok {
		return Value{typ: "error", str: fmt.Sprintf("%s not found", hash)}
	}

	values := []Value{}
	for k, v := range value {
		values = append(values, Value{typ: "bulk", bulk: k})
		values = append(values, Value{typ: "bulk", bulk: v})
	}

	return Value{typ: "array", array: values}
}
