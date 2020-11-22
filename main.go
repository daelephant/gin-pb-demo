/**
 * @Author: yin
 * @Description:main
 * @Version: 1.0.0
 * @Time : 2020-11-22 16:17
 */
package main

import (
	"fmt"
	"gin-pb-demo/module"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//输出
	r.GET("/protobuf", func(c *gin.Context) {
		data := &module.User{
			Name: "dada",
			Age:  22,
		}
		c.ProtoBuf(http.StatusOK, data) //存入protobuf

		fmt.Println(data)
		fmt.Println(data.Age)
		fmt.Printf("data:%+v \n", data.Age)
		c.JSON(http.StatusOK, data)
		return
	})

	r.GET("/getProto", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:8080/protobuf")
		if err != nil {
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				user := &module.User{}
				proto.UnmarshalMerge(body, user)
				fmt.Println(*user)
				//c.ProtoBuf(http.StatusOK, user)
				//return
				c.JSON(http.StatusOK, user)
				return
			}

		}

	})

	r.Run()

}
