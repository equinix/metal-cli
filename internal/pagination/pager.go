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

func GetDeviceEvents(s metal.EventsApiService, deviceId string, include []string, exclude []string) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		eventsPage, _, err := s.FindDeviceEvents(context.Background(), deviceId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}

func GetProjectEvents(s metal.EventsApiService, projectId string, include []string, exclude []string) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		eventsPage, _, err := s.FindProjectEvents(context.Background(), projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}

func GetOrganizationEvents(s metal.EventsApiService, orgId string, include []string, exclude []string) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		eventsPage, _, err := s.FindOrganizationEvents(context.Background(), orgId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}

func GetAllEvents(s metal.EventsApiService, include []string, exclude []string) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		eventsPage, _, err := s.FindEvents(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}
