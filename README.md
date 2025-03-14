# `kluster1`: Kubernetes TDD, Simplified.

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue)](https://golang.org/)

**Stop managing Kubernetes infrastructure, start testing your applications.**

`kluster1` is a Go module that empowers Test-Driven Development (TDD) for Kubernetes applications. It allows you to integrate local Kubernetes cluster lifecycle management directly into your Go test code, eliminating the need for external Kubernetes tools and ensuring consistent, fast, and reliable testing environments.
### Problem:

Developing and testing Kubernetes applications locally using traditional tools like Kind or Minikube is often:

* **Slow and cumbersome:**  Setting up and managing external cluster tools adds overhead and time to the development workflow.
* **Inconsistent:** Different developer setups lead to "works on my machine" issues and unreliable test results.
* **Focus-shifting:** Developers spend time managing infrastructure instead of focusing on writing tests and application code.

### Solution: `kluster1`

`kluster1` solves these problems by:

* **Embedding Kubernetes clusters in your tests:**  Manage cluster creation, configuration, and destruction directly within your Go test functions.
* **Providing consistent environments:** Ensures every developer gets the same, reliable Kubernetes environment for testing.
* **Speeding up TDD cycles:**  Reduces setup time and automates Kubernetes lifecycle management, leading to faster iteration and feedback.
* **Go Native:**  Feels natural for Go developers as a standard Go module.
* **Zero External Dependencies:** No need to install or configure external Kubernetes cluster tools.

### Key Features:

* **Programmatic Kubernetes Cluster Lifecycle:** Create, configure, and destroy local Kubernetes clusters within your Go tests.
* **Lightweight and Fast:** Optimised for rapid test execution and minimal overhead.
* **Easy Integration:**  Simple API to integrate into existing Go testing frameworks.

## Quick Start
1. Add the `kluster1` module to your Go project
```bash
go get github.com/ogarciacar/kluster1
```

2. Write your first Kubernetes App test
```go
// Example Go test using kluster1
package main

import (
  "testing"
  
  "github.com/ogarciacar/kluster1/kluster1"
)

func TestKubernetesApp(t *testing.T) {

  // 1. Creates a new Kubernetes cluster with a specified release version.
  k1, err := kluster1.NewCluster(kluster1.K8sRelease_v1_30_10)
  
  // 2. Check the cluster was successfully created
  if err != nil {
    t.Fatalf("should no be error %s", err)
  }

  // 3. Ensures the cluster is cleaned up after the test completes.
  defer k1.Destroy()
 
  // ... Your Kubernetes test logic here, using these objects:
  // clientset := k1.GetClientset()
  // kubeconfigPath := k1.GetKubeconfigPath()
  // ingressPort := k1.GetIngressPort()

  // ... example with k1.GetClientset()
  clientset := k1.GetClientset()
  serverVersion, err := clientset.ServerVersion()
  
  if err != nil {
    t.Fatalf("should no be error %s", err)
  }
  
  if kluster1.K8sRelease_v1_30_10.String() == serverVersion.String() {
    t.Fatalf("should be the same server version: expected %s got, %s", kluster1.K8sRelease_v1_30_10.String(), serverVersion.String())
  }
}
```

3. Run the test
```Shell
go test ./... -v
```

## More Examples?
For more comprehensive examples, check out:

* [Creating a NGIX Pod and make http calls](./kluster1/nginx_pod_test.go)
* [Set up a cluster to be shared by multiple tests in the same package](./kluster1/shared_cluster_test.go)