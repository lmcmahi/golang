/*******************************************************************************
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package server

import (
	"net/rpc"
	"log"
)

// dummy..

type Service struct {

}

type UserInfos struct {
	Name string
	Credential string
}

type Reply struct {
	Result interface{}
}
type BuyInput struct {
	UserInfos
	Number int

}

type SellInput struct {
	UserInfos
	Number int
}

type serviceOps interface {
	Buy(buyInput BuyInput, reply *Reply) error
	Sell( sellInput SellInput, reply *Reply) error
}

func (service *Service) Buy(buyInput BuyInput, reply *Reply) error {
	log.Println(buyInput.Number,buyInput.UserInfos)
	reply.Result = "none"
	return nil
}

func (service *Service) Sell( sellInput SellInput, reply *Reply) error {
	return nil
}
// the service is registed....
func RegisterService() {
	service := new (Service)
	// now register the operations...
	rpc.Register(service)
	rpc.HandleHTTP()
}