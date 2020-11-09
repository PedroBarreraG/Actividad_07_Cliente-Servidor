package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var activos []int64
var i []int64

func Proceso(id int64) {
	for {
		if activos[id] == 0 {
			return
		}
		fmt.Printf("id %d: %d\n", id, i[id])
		i[id] = i[id] + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	b := make([]byte, 100)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		mensaje := string(b[:bs])
		if mensaje == "C" {
			idAEnviar := idReturn()
			if idAEnviar != -1 {
				respuesta := strconv.FormatInt(int64(idAEnviar), 10)
				respuesta += ","
				respuesta += strconv.FormatInt(int64(i[idAEnviar]), 10)
				activos[idAEnviar] = 0
				c.Write([]byte(respuesta))
			} else {
				respuesta := "N"
				c.Write([]byte(respuesta))
			}
		} else {
			idI := strings.Split(mensaje, ",")
			newid, _ := strconv.ParseInt(idI[0], 10, 64)
			newi, _ := strconv.ParseInt(idI[1], 10, 64)
			activos[newid] = 1
			i[newid] = newi - 1
			go Proceso(newid)
		}
	}
}

func idReturn() int64 {
	if activos[0] == 1 {
		return 0
	}
	if activos[1] == 1 {
		return 1
	}
	if activos[2] == 1 {
		return 2
	}
	if activos[3] == 1 {
		return 3
	}
	if activos[4] == 1 {
		return 4
	}
	return -1
}

func main() {
	activos = append(activos, 1)
	activos = append(activos, 1)
	activos = append(activos, 1)
	activos = append(activos, 1)
	activos = append(activos, 1)

	i = append(i, 0)
	i = append(i, 0)
	i = append(i, 0)
	i = append(i, 0)
	i = append(i, 0)

	go Proceso(0)
	go Proceso(1)
	go Proceso(2)
	go Proceso(3)
	go Proceso(4)

	go servidor()

	var input string
	fmt.Scanln(&input)
}
