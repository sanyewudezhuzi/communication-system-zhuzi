package utils

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Transfer struct {
	Conn  net.Conn   // 连接
	Infos [8096]byte // 传输时使用的缓冲
}

// 读取数据
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Infos[:4])
	if err != nil {
		if err != io.EOF {
			fmt.Println("conn.Read(infos[:4]) fail, err =", err)
		}
		return
	}

	// 根据infos[:4]转成一共uint32类型
	pkgLen := binary.BigEndian.Uint32(this.Infos[:4])
	n, err := this.Conn.Read(this.Infos[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(infos[:pkgLen]) fail, err =", err)
		return
	}

	err = json.Unmarshal(this.Infos[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err =", err)
		return
	}
	return

}

// 发送请求
func (this *Transfer) WritePkg(data []byte) (err error) {

	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Infos[:4], pkgLen)

	// 发送data长度
	n, err := this.Conn.Write(this.Infos[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(infos[:4]) fail, err =", err)
		return
	}

	// 发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail, err =", err)
	}

	return
}
