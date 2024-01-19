package virtualcircuit

import (
	"context"
	"fmt"
	"strconv"
	"time"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

type vcParamOuter interface {
	GetId() string
	GetName() string
	GetSpeed() int64
	GetCreatedAt() time.Time
}

var (
	_ vcParamOuter = (*metal.VlanVirtualCircuit)(nil)
	_ vcParamOuter = (*metal.VrfVirtualCircuit)(nil)
)

func getParams(p vcParamOuter) (id, name, speed, time string) {
	id = p.GetId()
	name = p.GetName()
	sp := p.GetSpeed()
	speed = strconv.FormatInt(int64(sp), 10)
	time = p.GetCreatedAt().String()
	return id, name, speed, time
}

func (c *Client) Retrieve() *cobra.Command {
	var vcID string
	retrieveVirtualCircuitCmd := &cobra.Command{
		Use:     "get -i <id>",
		Aliases: []string{"list"},
		Short:   "Retrieves virtual circuit for a specific circuit Id.",
		Long:    "Retrieves virtual circuit for a specific circuit Id.",
		Example: `  # Retrieve virtual circuit for a specific circuit::

  # Retrieve the details of a specific virtual-circuit:
  metal vc get -i e9a969b3-8911-4667-9d99-57cd3dd4ef6f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			inc := []string{}
			header := []string{"ID", "Name", "Speed", "Created"}
			vcUID, _, err := c.Service.GetVirtualCircuit(context.Background(), vcID).Include(c.Servicer.Includes(inc)).Execute()
			if err != nil {
				return fmt.Errorf("could not listvirtual circuits : %w", err)
			}
			data := make([][]string, 1)

			id, name, speed, time := getParams(vcUID.VlanVirtualCircuit)
			data[0] = []string{id, name, speed, time}
			return c.Out.Output(vcUID, header, &data)
		},
	}
	retrieveVirtualCircuitCmd.Flags().StringVarP(&vcID, "id", "i", "", "Specify UUID of the virtual-circuit")
	_ = retrieveVirtualCircuitCmd.MarkFlagRequired("id")
	return retrieveVirtualCircuitCmd
}
