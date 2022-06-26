package tools

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// 读取数据
func readPkg(conn net.Conn) (mes message.Message, err error) {
	infos := make([]byte, 1024*4)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(infos[:4])
	if err != nil {
		if err != io.EOF {
			fmt.Println("conn.Read(infos[:4]) fail, err =", err)
		}
		return
	}

	// 根据infos[:4]转成一共uint32类型
	pkgLen := binary.BigEndian.Uint32(infos[:4])
	n, err := conn.Read(infos[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(infos[:pkgLen]) fail, err =", err)
		return
	}

	err = json.Unmarshal(infos[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err =", err)
		return
	}
	return mes, nil

}

// 发送请求
func writePkg(conn net.Conn, data []byte) (err error) {

	pkgLen := uint32(len(data))
	var infos [4]byte
	binary.BigEndian.PutUint32(infos[:4], pkgLen)

	// 发送data长度
	n, err := conn.Write(infos[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(infos[:4]) fail, err =", err)
		return
	}

	// 发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail, err =", err)
		return
	}

	return
}
