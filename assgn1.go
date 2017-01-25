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

	vms, _ := service.Mask("id,firstName,lastName,email,companyName,address1,city,users[username],users[id],users[firstName],users[lastName],users[email]").GetObject()
	fmt.Printf("ID: %d,\t FirstName: %s,\t LastName: %s,\t Email: %s,\t CompanyName: %s,\t Address: %s,\t City: %s\n", *vms.Id, *vms.FirstName, *vms.LastName, *vms.Email, *vms.CompanyName, *vms.Address1, *vms.City)

	fmt.Printf("\nPortal Users for above user:\n")

	for _, value := range vms.Users {
		fmt.Printf("ID: %d,\t Username: %s,\t FirstName: %s,\t LastName: %s,\t Email: %s,\n", *value.Id, *value.Username, sl.Grab(value, "FirstName"), sl.Grab(value, "LastName"), sl.Grab(value, "Email"))
	}

        vmds, _ := service.Mask("id,hostname,domain,billingItem[orderItem[order[userRecord[username]]]],operatingSystem[passwords]").GetVirtualGuests()
	fmt.Printf("\nList of VMs provisioned for user:\n")
        for _, vmd := range vmds {
                fmt.Printf("ID: [%d],\t Hostname: %s.%s,\t Username: %s,\t VMLoginUser: %s,\t VMLoginPassword: %s\n", *vmd.Id, *vmd.Hostname, *vmd.Domain, sl.Grab(vmd, "BillingItem.OrderItem.Order.UserRecord.Username"), *vmd.OperatingSystem.Passwords[0].Username, *vmd.OperatingSystem.Passwords[0].Password)
        }

}
