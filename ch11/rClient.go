package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var u1 = User{"admin", "admin"}
var u2 = User{"tsoukalos", "pass"}
var u3 = User{"", "pass"}

func deleteEndpoint(server string, user User) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusBadRequest
	}

	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest(http.MethodDelete, server+deleteEndPoint, u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	if resp == nil {
		return http.StatusBadRequest
	}

	data, err := io.ReadAll(resp.Body)
	fmt.Print("/delete returned: ", string(data))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return resp.StatusCode
}

func getEndpoint(server string, user User) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in unmarshalling: ", err)
		return http.StatusBadRequest
	}

	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest(http.MethodGet, server+getEndPoint, u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	if resp == nil {
		return resp.StatusCode
	}

	data, err := io.ReadAll(resp.Body)
	fmt.Print("/get returned: ", string(data))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return resp.StatusCode
}

func addEndpoint(server string, user User) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in unmarshalling: ", err)
		return http.StatusBadRequest
	}

	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest(http.MethodPost, server+addEndPoint, u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func timeEndpoint(server string) (int, string) {
	req, err := http.NewRequest(http.MethodPost, server+timeEndPoint, nil)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusBadRequest, ""
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode, ""
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(data)
}

func slashEndpoint(server, URL string) (int, string) {
	req, err := http.NewRequest(http.MethodPost, server+URL, nil)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusBadRequest, ""
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if resp == nil {
		return resp.StatusCode, ""
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(data)
}

const addEndPoint = "/add"
const getEndPoint = "/get"
const deleteEndPoint = "/delete"
const timeEndPoint = "/time"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Need: Server URL")
		return
	}
	server := os.Args[1]

	fmt.Println("/add")
	httpCode := addEndpoint(server, u1)
	if httpCode != http.StatusOK {
		fmt.Println("u1 Return code:", httpCode)
	} else {
		fmt.Println("u1 Data added:", u1, httpCode)
	}

	httpCode = addEndpoint(server, u2)
	if httpCode != http.StatusOK {
		fmt.Println("u2 Return code:", httpCode)
	} else {
		fmt.Println("u2 Data added:", u2, httpCode)
	}

	httpCode = addEndpoint(server, u3)
	if httpCode != http.StatusOK {
		fmt.Println("u3 Return code:", httpCode)
	} else {
		fmt.Println("u3 Data added:", u3, httpCode)
	}

	fmt.Println("/get")
	httpCode = getEndpoint(server, u1)
	fmt.Println("/get u1 return code:", httpCode)
	httpCode = getEndpoint(server, u2)
	fmt.Println("/get u2 return code:", httpCode)
	httpCode = getEndpoint(server, u3)
	fmt.Println("/get u3 return code:", httpCode)

	fmt.Println("/delete")
	httpCode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", httpCode)
	httpCode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", httpCode)
	httpCode = deleteEndpoint(server, u2)
	fmt.Println("/delete u2 return code:", httpCode)
	httpCode = deleteEndpoint(server, u3)
	fmt.Println("/delete u3 return code:", httpCode)

	fmt.Println("/time")
	httpCode, myTime := timeEndpoint(server)
	fmt.Print("/time returned: ", httpCode, " ", myTime)
	time.Sleep(time.Second)
	httpCode, myTime = timeEndpoint(server)
	fmt.Print("/time returned: ", httpCode, " ", myTime)

	fmt.Println("/")
	URL := "/"
	httpCode, response := slashEndpoint(server, URL)
	fmt.Print("/ returned: ", httpCode, " with response: ", response)

	fmt.Println("/what")
	URL = "/what"
	httpCode, response = slashEndpoint(server, URL)
	fmt.Print(URL, " returned: ", httpCode, " with response: ", response)
}
