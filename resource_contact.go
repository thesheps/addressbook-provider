package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func contact() *schema.Resource {
        return &schema.Resource{
                Create: contactCreate,
                Read:   contactRead,
                Update: contactUpdate,
                Delete: contactDelete,

                Schema: map[string]*schema.Schema{
                        "forename": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "surname": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func contactCreate(d *schema.ResourceData, m interface{}) error {
        forename := d.Get("forename").(string)
        surname := d.Get("surname").(string)

        d.SetId(forename + " " + surname)

	return contactRead(d, m)
}

func contactRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func contactUpdate(d *schema.ResourceData, m interface{}) error {
	return contactRead(d, m)
}

func contactDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}