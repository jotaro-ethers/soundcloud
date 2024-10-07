// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	AddListenHistory(ctx context.Context, arg AddListenHistoryParams) error
	AddSongToPlaylist(ctx context.Context, arg AddSongToPlaylistParams) error
	CancelSubscription(ctx context.Context, subscriptionID int32) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) error
	CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) error
	CreateReport(ctx context.Context, arg CreateReportParams) error
	CreateSong(ctx context.Context, arg CreateSongParams) error
	CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) error
	DeleteAccount(ctx context.Context, accountID int32) error
	DeleteComment(ctx context.Context, commentID int32) error
	DeletePlaylist(ctx context.Context, playlistID int32) error
	DeleteReport(ctx context.Context, reportID int32) error
	DeleteRepost(ctx context.Context, repostID int32) error
	DeleteSong(ctx context.Context, songID int32) error
	FollowUser(ctx context.Context, arg FollowUserParams) error
	GetAccountById(ctx context.Context, accountID int32) (Account, error)
	GetAllAccounts(ctx context.Context, arg GetAllAccountsParams) ([]Account, error)
	GetAllPlaylists(ctx context.Context) ([]Playlist, error)
	GetAllSongs(ctx context.Context) ([]Song, error)
	GetCommentsBySongId(ctx context.Context, songID int32) ([]Comment, error)
	GetFollowers(ctx context.Context, followingID int32) ([]Follow, error)
	GetFollowing(ctx context.Context, followerID int32) ([]Follow, error)
	GetLikesBySongId(ctx context.Context, songID int32) ([]Like, error)
	GetListenHistoryByAccountId(ctx context.Context, accountID int32) ([]ListenHistory, error)
	GetPlaylistById(ctx context.Context, playlistID int32) (Playlist, error)
	GetReportsByAccountId(ctx context.Context, accountID int32) ([]Report, error)
	GetReportsBySongId(ctx context.Context, songID int32) ([]Report, error)
	GetRepostsByAccountId(ctx context.Context, accountID int32) ([]Repost, error)
	GetRepostsBySongId(ctx context.Context, songID int32) ([]Repost, error)
	GetSongById(ctx context.Context, songID int32) (Song, error)
	GetSongsByPlaylistId(ctx context.Context, playlistID int32) ([]Song, error)
	GetSubscriptionsByAccountId(ctx context.Context, accountID int32) ([]Subscription, error)
	LikeSong(ctx context.Context, arg LikeSongParams) error
	RemoveSongFromPlaylist(ctx context.Context, arg RemoveSongFromPlaylistParams) error
	RepostSong(ctx context.Context, arg RepostSongParams) error
	UnfollowUser(ctx context.Context, arg UnfollowUserParams) error
	UnlikeSong(ctx context.Context, arg UnlikeSongParams) error
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	UpdatePlaylist(ctx context.Context, arg UpdatePlaylistParams) error
	UpdateSong(ctx context.Context, arg UpdateSongParams) error
}

var _ Querier = (*Queries)(nil)
