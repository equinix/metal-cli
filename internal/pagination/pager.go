package pager

import (
	"context"

	metal "github.com/equinix-labs/metal-go/metal/v1"
)

func GetAllProjects(s metal.ProjectsApiService, include []string, exclude []string) ([]metal.Project, error) {
	var projects []metal.Project

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)
	for {
		projectPage, _, err := s.FindProjects(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}
		projects = append(projects, projectPage.GetProjects()...)
		if projectPage.Meta.GetLastPage() > projectPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return projects, nil
	}
}

func GetProjectDevices(s metal.ApiFindProjectDevicesRequest) ([]metal.Device, error) {
	var devices []metal.Device

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)
	for {
		devicePage, _, err := s.Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}
		devices = append(devices, devicePage.Devices...)
		if devicePage.Meta.GetLastPage() > devicePage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return devices, nil
	}
}
