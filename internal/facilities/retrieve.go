package facilities

import (
	"context"
	"fmt"
	"strings"

	"github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {

	retrieveFacilitesCmd := &cobra.Command{
		Use:     `get`,
		Aliases: []string{"list"},
		Short:   "Retrieves a list of facilities.",
		Long:    "Retrieves a list of facilities available to the current user.",
		Example: `  # Lists facilities for current user:
  metal facilities get`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				facilityList *metalv1.FacilityList
				err          error
			)
			cmd.SilenceUsage = true

			inc := []metalv1.FindFacilitiesIncludeParameterInner{}
			exc := []metalv1.FindFacilitiesIncludeParameterInner{}
			facilityList, _, err = c.Service.FindFacilities(context.Background()).Include(inc).Exclude(exc).Execute()
			if err != nil {
				return fmt.Errorf("Could not list Facilities: %w", err)
			}

			facilities := facilityList.GetFacilities()
			data := make([][]string, len(facilities))

			for i, facility := range facilities {
				var metro string
				if facility.Metro != nil {
					metro = facility.Metro.GetCode()
				}

				facilityFeatures := facility.GetFeatures()
				var stringFeatures []string
				for _, feature := range facilityFeatures {
					stringFeatures = append(stringFeatures, string(feature))
				}
				data[i] = []string{facility.GetName(), facility.GetCode(), metro, strings.Join(([]string)(stringFeatures), ",")}
			}
			header := []string{"Name", "Code", "Metro", "Features"}
			return c.Out.Output(facilities, header, &data)
		},
	}

	return retrieveFacilitesCmd
}
