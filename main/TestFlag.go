package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "everyOne", "The Greeting Object")
}
func main() {

}
