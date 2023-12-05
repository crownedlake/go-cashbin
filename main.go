package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
    enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
    if err != nil {
        fmt.Println("Error in Casbin Enforcer creation:", err)
        return
    }

    // Define some test requests
    requests := [][]string{
        {"alice", "data1", "read"},
        {"bob", "data1", "write"},
        {"alice", "data2", "write"},
        {"bob", "data2", "read"},
    }

    // Check the requests
    for _, req := range requests {
        result, err := enforcer.Enforce(req[0], req[1], req[2])
        if err != nil {
            fmt.Println("Error in enforcement:", err)
        } else if result {
            fmt.Printf("%s can access %s with %s permission\n", req[0], req[1], req[2])
        } else {
            fmt.Printf("%s cannot access %s with %s permission\n", req[0], req[1], req[2])
        }
    }
}
