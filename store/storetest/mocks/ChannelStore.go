// Code generated by mockery v1.0.0

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// ChannelStore is an autogenerated mock type for the ChannelStore type
type ChannelStore struct {
	mock.Mock
}

// AnalyticsDeletedTypeCount provides a mock function with given fields: teamId, channelType
func (_m *ChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) store.StoreChannel {
	ret := _m.Called(teamId, channelType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, channelType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// AnalyticsTypeCount provides a mock function with given fields: teamId, channelType
func (_m *ChannelStore) AnalyticsTypeCount(teamId string, channelType string) store.StoreChannel {
	ret := _m.Called(teamId, channelType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, channelType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// AutocompleteInTeam provides a mock function with given fields: teamId, term
func (_m *ChannelStore) AutocompleteInTeam(teamId string, term string) store.StoreChannel {
	ret := _m.Called(teamId, term)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// ClearCaches provides a mock function with given fields:
func (_m *ChannelStore) ClearCaches() {
	_m.Called()
}

// CreateDirectChannel provides a mock function with given fields: userId, otherUserId
func (_m *ChannelStore) CreateDirectChannel(userId string, otherUserId string) store.StoreChannel {
	ret := _m.Called(userId, otherUserId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, otherUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: channelId, time
func (_m *ChannelStore) Delete(channelId string, time int64) store.StoreChannel {
	ret := _m.Called(channelId, time)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64) store.StoreChannel); ok {
		r0 = rf(channelId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id, allowFromCache
func (_m *ChannelStore) Get(id string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(id, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(id, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: teamId
func (_m *ChannelStore) GetAll(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllChannelMembersForUser provides a mock function with given fields: userId, allowFromCache
func (_m *ChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(userId, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(userId, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllChannelMembersNotifyPropsForChannel provides a mock function with given fields: channelId, allowFromCache
func (_m *ChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(channelId, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(channelId, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByName provides a mock function with given fields: team_id, name, allowFromCache
func (_m *ChannelStore) GetByName(team_id string, name string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(team_id, name, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, bool) store.StoreChannel); ok {
		r0 = rf(team_id, name, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByNameIncludeDeleted provides a mock function with given fields: team_id, name, allowFromCache
func (_m *ChannelStore) GetByNameIncludeDeleted(team_id string, name string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(team_id, name, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, bool) store.StoreChannel); ok {
		r0 = rf(team_id, name, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByNames provides a mock function with given fields: team_id, names, allowFromCache
func (_m *ChannelStore) GetByNames(team_id string, names []string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(team_id, names, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, []string, bool) store.StoreChannel); ok {
		r0 = rf(team_id, names, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetChannelCounts provides a mock function with given fields: teamId, userId
func (_m *ChannelStore) GetChannelCounts(teamId string, userId string) store.StoreChannel {
	ret := _m.Called(teamId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetChannelUnread provides a mock function with given fields: channelId, userId
func (_m *ChannelStore) GetChannelUnread(channelId string, userId string) store.StoreChannel {
	ret := _m.Called(channelId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(channelId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetChannels provides a mock function with given fields: teamId, userId
func (_m *ChannelStore) GetChannels(teamId string, userId string) store.StoreChannel {
	ret := _m.Called(teamId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetDeleted provides a mock function with given fields: team_id, offset, limit
func (_m *ChannelStore) GetDeleted(team_id string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(team_id, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(team_id, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetDeletedByName provides a mock function with given fields: team_id, name
func (_m *ChannelStore) GetDeletedByName(team_id string, name string) store.StoreChannel {
	ret := _m.Called(team_id, name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(team_id, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetForPost provides a mock function with given fields: postId
func (_m *ChannelStore) GetForPost(postId string) store.StoreChannel {
	ret := _m.Called(postId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetFromMaster provides a mock function with given fields: id
func (_m *ChannelStore) GetFromMaster(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMember provides a mock function with given fields: channelId, userId
func (_m *ChannelStore) GetMember(channelId string, userId string) store.StoreChannel {
	ret := _m.Called(channelId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(channelId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMemberCount provides a mock function with given fields: channelId, allowFromCache
func (_m *ChannelStore) GetMemberCount(channelId string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(channelId, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(channelId, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMemberCountFromCache provides a mock function with given fields: channelId
func (_m *ChannelStore) GetMemberCountFromCache(channelId string) int64 {
	ret := _m.Called(channelId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(channelId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetMemberForPost provides a mock function with given fields: postId, userId
func (_m *ChannelStore) GetMemberForPost(postId string, userId string) store.StoreChannel {
	ret := _m.Called(postId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(postId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMembers provides a mock function with given fields: channelId, offset, limit
func (_m *ChannelStore) GetMembers(channelId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(channelId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(channelId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMembersByIds provides a mock function with given fields: channelId, userIds
func (_m *ChannelStore) GetMembersByIds(channelId string, userIds []string) store.StoreChannel {
	ret := _m.Called(channelId, userIds)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, []string) store.StoreChannel); ok {
		r0 = rf(channelId, userIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMembersForUser provides a mock function with given fields: teamId, userId
func (_m *ChannelStore) GetMembersForUser(teamId string, userId string) store.StoreChannel {
	ret := _m.Called(teamId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMoreChannels provides a mock function with given fields: teamId, userId, offset, limit
func (_m *ChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(teamId, userId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int, int) store.StoreChannel); ok {
		r0 = rf(teamId, userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPinnedPosts provides a mock function with given fields: channelId
func (_m *ChannelStore) GetPinnedPosts(channelId string) store.StoreChannel {
	ret := _m.Called(channelId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPublicChannelsByIdsForTeam provides a mock function with given fields: teamId, channelIds
func (_m *ChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) store.StoreChannel {
	ret := _m.Called(teamId, channelIds)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, []string) store.StoreChannel); ok {
		r0 = rf(teamId, channelIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPublicChannelsForTeam provides a mock function with given fields: teamId, offset, limit
func (_m *ChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(teamId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(teamId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTeamChannels provides a mock function with given fields: teamId
func (_m *ChannelStore) GetTeamChannels(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// IncrementMentionCount provides a mock function with given fields: channelId, userId
func (_m *ChannelStore) IncrementMentionCount(channelId string, userId string) store.StoreChannel {
	ret := _m.Called(channelId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(channelId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// InvalidateAllChannelMembersForUser provides a mock function with given fields: userId
func (_m *ChannelStore) InvalidateAllChannelMembersForUser(userId string) {
	_m.Called(userId)
}

// InvalidateCacheForChannelMembersNotifyProps provides a mock function with given fields: channelId
func (_m *ChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {
	_m.Called(channelId)
}

// InvalidateChannel provides a mock function with given fields: id
func (_m *ChannelStore) InvalidateChannel(id string) {
	_m.Called(id)
}

// InvalidateChannelByName provides a mock function with given fields: teamId, name
func (_m *ChannelStore) InvalidateChannelByName(teamId string, name string) {
	_m.Called(teamId, name)
}

// InvalidateMemberCount provides a mock function with given fields: channelId
func (_m *ChannelStore) InvalidateMemberCount(channelId string) {
	_m.Called(channelId)
}

// IsUserInChannelUseCache provides a mock function with given fields: userId, channelId
func (_m *ChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {
	ret := _m.Called(userId, channelId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(userId, channelId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PermanentDelete provides a mock function with given fields: channelId
func (_m *ChannelStore) PermanentDelete(channelId string) store.StoreChannel {
	ret := _m.Called(channelId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteByTeam provides a mock function with given fields: teamId
func (_m *ChannelStore) PermanentDeleteByTeam(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteMembersByChannel provides a mock function with given fields: channelId
func (_m *ChannelStore) PermanentDeleteMembersByChannel(channelId string) store.StoreChannel {
	ret := _m.Called(channelId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteMembersByUser provides a mock function with given fields: userId
func (_m *ChannelStore) PermanentDeleteMembersByUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// RemoveMember provides a mock function with given fields: channelId, userId
func (_m *ChannelStore) RemoveMember(channelId string, userId string) store.StoreChannel {
	ret := _m.Called(channelId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(channelId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Restore provides a mock function with given fields: channelId, time
func (_m *ChannelStore) Restore(channelId string, time int64) store.StoreChannel {
	ret := _m.Called(channelId, time)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64) store.StoreChannel); ok {
		r0 = rf(channelId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: channel, maxChannelsPerTeam
func (_m *ChannelStore) Save(channel *model.Channel, maxChannelsPerTeam int64) store.StoreChannel {
	ret := _m.Called(channel, maxChannelsPerTeam)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Channel, int64) store.StoreChannel); ok {
		r0 = rf(channel, maxChannelsPerTeam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveDirectChannel provides a mock function with given fields: channel, member1, member2
func (_m *ChannelStore) SaveDirectChannel(channel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) store.StoreChannel {
	ret := _m.Called(channel, member1, member2)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Channel, *model.ChannelMember, *model.ChannelMember) store.StoreChannel); ok {
		r0 = rf(channel, member1, member2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveMember provides a mock function with given fields: member
func (_m *ChannelStore) SaveMember(member *model.ChannelMember) store.StoreChannel {
	ret := _m.Called(member)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.ChannelMember) store.StoreChannel); ok {
		r0 = rf(member)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SearchInTeam provides a mock function with given fields: teamId, term
func (_m *ChannelStore) SearchInTeam(teamId string, term string) store.StoreChannel {
	ret := _m.Called(teamId, term)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SearchMore provides a mock function with given fields: userId, teamId, term
func (_m *ChannelStore) SearchMore(userId string, teamId string, term string) store.StoreChannel {
	ret := _m.Called(userId, teamId, term)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, string) store.StoreChannel); ok {
		r0 = rf(userId, teamId, term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SetDeleteAt provides a mock function with given fields: channelId, deleteAt, updateAt
func (_m *ChannelStore) SetDeleteAt(channelId string, deleteAt int64, updateAt int64) store.StoreChannel {
	ret := _m.Called(channelId, deleteAt, updateAt)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64, int64) store.StoreChannel); ok {
		r0 = rf(channelId, deleteAt, updateAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: channel
func (_m *ChannelStore) Update(channel *model.Channel) store.StoreChannel {
	ret := _m.Called(channel)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Channel) store.StoreChannel); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateLastViewedAt provides a mock function with given fields: channelIds, userId
func (_m *ChannelStore) UpdateLastViewedAt(channelIds []string, userId string) store.StoreChannel {
	ret := _m.Called(channelIds, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func([]string, string) store.StoreChannel); ok {
		r0 = rf(channelIds, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateMember provides a mock function with given fields: member
func (_m *ChannelStore) UpdateMember(member *model.ChannelMember) store.StoreChannel {
	ret := _m.Called(member)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.ChannelMember) store.StoreChannel); ok {
		r0 = rf(member)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
