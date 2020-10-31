package provider

import (
	"fmt"
	"context"
	"strings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sethvargo/go-diceware/diceware"
)

func resourcePassword() *schema.Resource {
	return &schema.Resource{
		Description: "The resource `passphrase_password` generates a passphrase using diceware algorithm.",
		Create: createPassphraseFunc,
		Read: readNil,
		Delete: schema.RemoveFromState,
		MigrateState: resourceRandomPassphraseMigrateState,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"word_count": {
				Description: "Number of word in the result.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     12,
				ForceNew:    true,
			},
			"separator": {
				Description: "Separator of each word in the result",
				Type: schema.TypeString,
				Optional: true,
				Default: "-",
				ForceNew: true,
			},
			"result": {
				Description: "The generated random string.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: importStringFunc,
		},
	}
}

func resourceRandomPassphraseMigrateState(
	v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	switch v {
	default:
		return is, fmt.Errorf("Unexpected schema version: %d", v)
	}
}

func createPassphraseFunc(d *schema.ResourceData, meta interface{}) error {
	wordCount := d.Get("word_count").(int)
	separator := d.Get("separator").(string)
	
	words, err := diceware.Generate(wordCount)
	if err != nil  {
		return err
	}

	d.Set("result", strings.Join(words, separator))
	d.SetId("none")
	return nil
}

func readNil(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func importStringFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	val := d.Id()
	d.SetId("none")
	d.Set("result", val)
	return []*schema.ResourceData{d}, nil
}
