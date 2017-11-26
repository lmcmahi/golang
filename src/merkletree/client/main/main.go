/*******************************************************************************
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"net/rpc"
	"log"
	"merkletree/server"
)

func main(){
	client,err := rpc.DialHTTP("tcp","localhost:8976")
	if err != nil {
		log.Fatal("Error connecting..",err)
	}
	userInfo := server.UserInfos{Name:"madi",Credential:"coco"}
	args := &server.BuyInput{}
	args.UserInfos = userInfo
	args.Number = 23
	reply := server.Reply{}
	err = client.Call("Service.Buy",args,&reply)
	if err != nil {
		log.Fatal("Error making the call..",err)
	}
	log.Println("Returned response..",reply)
}