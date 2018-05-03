// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import "github.com/mattermost/mattermost-server/model"

func (a *App) GetScheme(id string) (*model.Scheme, *model.AppError) {
	if result := <-a.Srv.Store.Scheme().Get(id); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Scheme), nil
	}
}

func (a *App) GetSchemesPage(scope string, page int, perPage int) ([]*model.Scheme, *model.AppError) {
	return a.GetSchemes(scope, page*perPage, perPage)
}

func (a *App) GetSchemes(scope string, offset int, limit int) ([]*model.Scheme, *model.AppError) {
	if result := <-a.Srv.Store.Scheme().GetAllPage(scope, offset, limit); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Scheme), nil
	}
}

func (a *App) CreateScheme(scheme *model.Scheme) (*model.Scheme, *model.AppError) {
	if err := a.IsPhase2MigrationCompleted(); err != nil {
		return nil, err
	}

	// Clear any user-provided values for trusted properties.
	scheme.DefaultTeamAdminRole = ""
	scheme.DefaultTeamUserRole = ""
	scheme.DefaultChannelAdminRole = ""
	scheme.DefaultChannelUserRole = ""
	scheme.CreateAt = 0
	scheme.UpdateAt = 0
	scheme.DeleteAt = 0

	if result := <-a.Srv.Store.Scheme().Save(scheme); result.Err != nil {
		return nil, result.Err
	} else {
		return scheme, nil
	}
}

func (a *App) PatchScheme(scheme *model.Scheme, patch *model.SchemePatch) (*model.Scheme, *model.AppError) {
	if err := a.IsPhase2MigrationCompleted(); err != nil {
		return nil, err
	}

	scheme.Patch(patch)
	scheme, err := a.UpdateScheme(scheme)
	if err != nil {
		return nil, err
	}

	return scheme, err
}

func (a *App) UpdateScheme(scheme *model.Scheme) (*model.Scheme, *model.AppError) {
	if err := a.IsPhase2MigrationCompleted(); err != nil {
		return nil, err
	}

	if result := <-a.Srv.Store.Scheme().Save(scheme); result.Err != nil {
		return nil, result.Err
	} else {
		return scheme, nil
	}
}

func (a *App) DeleteScheme(schemeId string) (*model.Scheme, *model.AppError) {
	if err := a.IsPhase2MigrationCompleted(); err != nil {
		return nil, err
	}

	if result := <-a.Srv.Store.Scheme().Delete(schemeId); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Scheme), nil
	}
}

func (a *App) GetTeamsForSchemePage(scheme *model.Scheme, page int, perPage int) ([]*model.Team, *model.AppError) {
	return a.GetTeamsForScheme(scheme, page*perPage, perPage)
}

func (a *App) GetTeamsForScheme(scheme *model.Scheme, offset int, limit int) ([]*model.Team, *model.AppError) {
	if result := <-a.Srv.Store.Team().GetTeamsByScheme(scheme.Id, offset, limit); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Team), nil
	}
}

func (a *App) GetChannelsForSchemePage(scheme *model.Scheme, page int, perPage int) (model.ChannelList, *model.AppError) {
	return a.GetChannelsForScheme(scheme, page*perPage, perPage)
}

func (a *App) GetChannelsForScheme(scheme *model.Scheme, offset int, limit int) (model.ChannelList, *model.AppError) {
	if result := <-a.Srv.Store.Channel().GetChannelsByScheme(scheme.Id, offset, limit); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(model.ChannelList), nil
	}
}

func (a *App) IsPhase2MigrationCompleted() *model.AppError {
	// TODO: Actually check the Phase 2 migration has completed before permitting these actions.

	return nil
}
