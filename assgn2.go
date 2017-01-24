package main

import (
        "fmt"
        "github.com/softlayer/softlayer-go/services"
        "github.com/softlayer/softlayer-go/session"
        "github.com/softlayer/softlayer-go/sl"
)

func main() {
        sess := session.New()

        service := services.GetAccountService(sess)

        vms, _ := service.Mask("id,hostname,domain,billingItem[orderItem[order[userRecord[username]]]],operatingSystem[passwords]").GetVirtualGuests()

        for _, vm := range vms {
                fmt.Printf("ID: [%d],\t Hostname: %s.%s,\t Username: %s,\t VMLoginUser: %s,\t VMLoginPassword: %s\n", *vm.Id, *vm.Hostname, *vm.Domain, sl.Grab(vm, "BillingItem.OrderItem.Order.UserRecord.Username"), *vm.OperatingSystem.Passwords[0].Username, *vm.OperatingSystem.Passwords[0].Password)
        }
}
