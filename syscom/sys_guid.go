/*
create guid
*/
package syscom

import (
	// "crypto/rand"
	// "encoding/base64"
	// "io"
	"sync/atomic"
	"time"

	uuid "github.com/satori/go.uuid"
)

//系统计数器
var counter uint64

//偏移常量
const Offset = 1654990515

func GetUintID() int64 {
	atomic.AddUint64(&counter, 1)

	return time.Now().UnixNano() + int64(counter) + Offset
}

/**
 * 获取一个Guid值
 */
// func GetGUID_48() string {
// 	b := make([]byte, 48)
// 	if _, err := io.ReadFull(rand.Reader, b); err != nil {
// 		return ""
// 	}
// 	return GetMd5String(base64.URLEncoding.EncodeToString(b))
// }

func GetGUID() string {

	id, _ := uuid.NewV4()
	return id.String()
}
