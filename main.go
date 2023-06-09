package main

import (
  discovery "k8s.io/client-go/discovery"
  restclient "k8s.io/client-go/rest"
)
func main() {
  fakeconfig, _ := discovery.NewDiscoveryClientForConfig(&restclient.Config{})
  // silence not used error
  _ = fakeconfig
}
