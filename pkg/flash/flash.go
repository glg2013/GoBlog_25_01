// Package flash 用以支持在绘画中存储消息提示
package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

// Flashes Flash 消息数组类型，用已在会话中存储 map
type Flashes map[string]interface{}

// 存入会话数据里的 key
var flashKey = "_flashes"

func init() {
	// 在 gorilla/sessions 中存储
	// map 和 struct 数据需要提前注册 gob，方便后续 gob 序列化编码、解码
	gob.Register(Flashes{})
}

// Info 新增 Info 类型的消息提示
func Info(message string) {
	addFlash("info", message)
}

// Warning 新增 Warning 类型的消息提示
func Warning(message string) {
	addFlash("warning", message)
}

// Danger 新增 Danger 类型的消息提示
func Danger(message string) {
	addFlash("danger", message)
}

// Success 新增 Success 类型的消息提示
func Success(message string) {
	addFlash("success", message)
}

func All() Flashes {
	val := session.Get(flashKey)
	// 读取时必须做类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	// 读取即销毁，直接删除
	session.Forget(flashKey)
	return flashMessages
}

// 私有方法，新增一条提示
func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
}
