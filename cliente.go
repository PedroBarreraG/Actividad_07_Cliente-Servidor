package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var activo int64
var i int64
var idAEnviar int64

func Proceso(id int64) {
	for {
		if activo == 0 {
			return
		}

		fmt.Printf("id %d: %d\n", id, i)
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func cliente() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "C"
	c.Write([]byte(msg))

	for {
		if err != nil {
			fmt.Println(err)
			continue
		}
		b := make([]byte, 100)
		bs, err := c.Read(b)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			mensaje := string(b[:bs])
			if mensaje == "N" {
				fmt.Println("No hay procesos disponibles")
				break
			}

			idI := strings.Split(mensaje, ",")
			newid, _ := strconv.ParseInt(idI[0], 10, 64)
			newi, _ := strconv.ParseInt(idI[1], 10, 64)
			activo = 1
			i = newi
			idAEnviar = newid
			go Proceso(newid)
			break
		}
	}
	c.Close()
}

func clienteEND() {
	if activo == 1 {
		c, err := net.Dial("tcp", ":9999")
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := strconv.FormatInt(int64(idAEnviar), 10)
		msg += ","
		msg += strconv.FormatInt(int64(i), 10)
		activo = 0

		c.Write([]byte(msg))
		c.Close()
	}
}

func main() {
	activo = 0

	go cliente()

	var input string
	fmt.Scanln(&input)

	go clienteEND()

	fmt.Scanln(&input)
	fmt.Println("Cliente Terminado")
}
