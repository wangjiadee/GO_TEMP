/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2020-09-16 09:46:13
 * @LastEditTime: 2020-09-16 09:52:23
 * @FilePath: /go/src/github.com/Ralph.org/Element/String_Strut/string_addr.go
 * @Effect: DO
 */
package main 

import (
	"strings"
	"fmt"
)




func main(){
	str := "Fucking you! Fucking everythings"

	// Cut according to string
	ret := strings.Split(str,"uc")
	for _,s := range ret {
		fmt.Println(s)
	}

	//  Judge string end mark
	flg := strings.HasSuffix("Fuck u,tonny","tom")
	fmt.Println(flg) 
	//  Judge string start mark 
	flg1 := strings.HasPrefix("Fuck u,tonny","Fuck")
	fmt.Println(flg1) 
	
}