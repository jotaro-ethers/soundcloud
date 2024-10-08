// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addListenHistoryStmt, err = db.PrepareContext(ctx, addListenHistory); err != nil {
		return nil, fmt.Errorf("error preparing query AddListenHistory: %w", err)
	}
	if q.addSongToPlaylistStmt, err = db.PrepareContext(ctx, addSongToPlaylist); err != nil {
		return nil, fmt.Errorf("error preparing query AddSongToPlaylist: %w", err)
	}
	if q.cancelSubscriptionStmt, err = db.PrepareContext(ctx, cancelSubscription); err != nil {
		return nil, fmt.Errorf("error preparing query CancelSubscription: %w", err)
	}
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createCommentStmt, err = db.PrepareContext(ctx, createComment); err != nil {
		return nil, fmt.Errorf("error preparing query CreateComment: %w", err)
	}
	if q.createPlaylistStmt, err = db.PrepareContext(ctx, createPlaylist); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePlaylist: %w", err)
	}
	if q.createReportStmt, err = db.PrepareContext(ctx, createReport); err != nil {
		return nil, fmt.Errorf("error preparing query CreateReport: %w", err)
	}
	if q.createSongStmt, err = db.PrepareContext(ctx, createSong); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSong: %w", err)
	}
	if q.createSubscriptionStmt, err = db.PrepareContext(ctx, createSubscription); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSubscription: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteCommentStmt, err = db.PrepareContext(ctx, deleteComment); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteComment: %w", err)
	}
	if q.deletePlaylistStmt, err = db.PrepareContext(ctx, deletePlaylist); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePlaylist: %w", err)
	}
	if q.deleteReportStmt, err = db.PrepareContext(ctx, deleteReport); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteReport: %w", err)
	}
	if q.deleteRepostStmt, err = db.PrepareContext(ctx, deleteRepost); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRepost: %w", err)
	}
	if q.deleteSongStmt, err = db.PrepareContext(ctx, deleteSong); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSong: %w", err)
	}
	if q.followUserStmt, err = db.PrepareContext(ctx, followUser); err != nil {
		return nil, fmt.Errorf("error preparing query FollowUser: %w", err)
	}
	if q.getAccountByIdStmt, err = db.PrepareContext(ctx, getAccountById); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountById: %w", err)
	}
	if q.getAllAccountsStmt, err = db.PrepareContext(ctx, getAllAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllAccounts: %w", err)
	}
	if q.getAllPlaylistsStmt, err = db.PrepareContext(ctx, getAllPlaylists); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllPlaylists: %w", err)
	}
	if q.getAllSongsStmt, err = db.PrepareContext(ctx, getAllSongs); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllSongs: %w", err)
	}
	if q.getCommentsBySongIdStmt, err = db.PrepareContext(ctx, getCommentsBySongId); err != nil {
		return nil, fmt.Errorf("error preparing query GetCommentsBySongId: %w", err)
	}
	if q.getFollowersStmt, err = db.PrepareContext(ctx, getFollowers); err != nil {
		return nil, fmt.Errorf("error preparing query GetFollowers: %w", err)
	}
	if q.getFollowingStmt, err = db.PrepareContext(ctx, getFollowing); err != nil {
		return nil, fmt.Errorf("error preparing query GetFollowing: %w", err)
	}
	if q.getLikesBySongIdStmt, err = db.PrepareContext(ctx, getLikesBySongId); err != nil {
		return nil, fmt.Errorf("error preparing query GetLikesBySongId: %w", err)
	}
	if q.getListenHistoryByAccountIdStmt, err = db.PrepareContext(ctx, getListenHistoryByAccountId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListenHistoryByAccountId: %w", err)
	}
	if q.getPlaylistByIdStmt, err = db.PrepareContext(ctx, getPlaylistById); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlaylistById: %w", err)
	}
	if q.getReportsByAccountIdStmt, err = db.PrepareContext(ctx, getReportsByAccountId); err != nil {
		return nil, fmt.Errorf("error preparing query GetReportsByAccountId: %w", err)
	}
	if q.getReportsBySongIdStmt, err = db.PrepareContext(ctx, getReportsBySongId); err != nil {
		return nil, fmt.Errorf("error preparing query GetReportsBySongId: %w", err)
	}
	if q.getRepostsByAccountIdStmt, err = db.PrepareContext(ctx, getRepostsByAccountId); err != nil {
		return nil, fmt.Errorf("error preparing query GetRepostsByAccountId: %w", err)
	}
	if q.getRepostsBySongIdStmt, err = db.PrepareContext(ctx, getRepostsBySongId); err != nil {
		return nil, fmt.Errorf("error preparing query GetRepostsBySongId: %w", err)
	}
	if q.getSongByIdStmt, err = db.PrepareContext(ctx, getSongById); err != nil {
		return nil, fmt.Errorf("error preparing query GetSongById: %w", err)
	}
	if q.getSongsByPlaylistIdStmt, err = db.PrepareContext(ctx, getSongsByPlaylistId); err != nil {
		return nil, fmt.Errorf("error preparing query GetSongsByPlaylistId: %w", err)
	}
	if q.getSubscriptionsByAccountIdStmt, err = db.PrepareContext(ctx, getSubscriptionsByAccountId); err != nil {
		return nil, fmt.Errorf("error preparing query GetSubscriptionsByAccountId: %w", err)
	}
	if q.likeSongStmt, err = db.PrepareContext(ctx, likeSong); err != nil {
		return nil, fmt.Errorf("error preparing query LikeSong: %w", err)
	}
	if q.removeSongFromPlaylistStmt, err = db.PrepareContext(ctx, removeSongFromPlaylist); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveSongFromPlaylist: %w", err)
	}
	if q.repostSongStmt, err = db.PrepareContext(ctx, repostSong); err != nil {
		return nil, fmt.Errorf("error preparing query RepostSong: %w", err)
	}
	if q.unfollowUserStmt, err = db.PrepareContext(ctx, unfollowUser); err != nil {
		return nil, fmt.Errorf("error preparing query UnfollowUser: %w", err)
	}
	if q.unlikeSongStmt, err = db.PrepareContext(ctx, unlikeSong); err != nil {
		return nil, fmt.Errorf("error preparing query UnlikeSong: %w", err)
	}
	if q.updateAccountStmt, err = db.PrepareContext(ctx, updateAccount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccount: %w", err)
	}
	if q.updatePlaylistStmt, err = db.PrepareContext(ctx, updatePlaylist); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePlaylist: %w", err)
	}
	if q.updateSongStmt, err = db.PrepareContext(ctx, updateSong); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSong: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addListenHistoryStmt != nil {
		if cerr := q.addListenHistoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addListenHistoryStmt: %w", cerr)
		}
	}
	if q.addSongToPlaylistStmt != nil {
		if cerr := q.addSongToPlaylistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addSongToPlaylistStmt: %w", cerr)
		}
	}
	if q.cancelSubscriptionStmt != nil {
		if cerr := q.cancelSubscriptionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cancelSubscriptionStmt: %w", cerr)
		}
	}
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createCommentStmt != nil {
		if cerr := q.createCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCommentStmt: %w", cerr)
		}
	}
	if q.createPlaylistStmt != nil {
		if cerr := q.createPlaylistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPlaylistStmt: %w", cerr)
		}
	}
	if q.createReportStmt != nil {
		if cerr := q.createReportStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createReportStmt: %w", cerr)
		}
	}
	if q.createSongStmt != nil {
		if cerr := q.createSongStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSongStmt: %w", cerr)
		}
	}
	if q.createSubscriptionStmt != nil {
		if cerr := q.createSubscriptionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSubscriptionStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteCommentStmt != nil {
		if cerr := q.deleteCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCommentStmt: %w", cerr)
		}
	}
	if q.deletePlaylistStmt != nil {
		if cerr := q.deletePlaylistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePlaylistStmt: %w", cerr)
		}
	}
	if q.deleteReportStmt != nil {
		if cerr := q.deleteReportStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteReportStmt: %w", cerr)
		}
	}
	if q.deleteRepostStmt != nil {
		if cerr := q.deleteRepostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRepostStmt: %w", cerr)
		}
	}
	if q.deleteSongStmt != nil {
		if cerr := q.deleteSongStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSongStmt: %w", cerr)
		}
	}
	if q.followUserStmt != nil {
		if cerr := q.followUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing followUserStmt: %w", cerr)
		}
	}
	if q.getAccountByIdStmt != nil {
		if cerr := q.getAccountByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountByIdStmt: %w", cerr)
		}
	}
	if q.getAllAccountsStmt != nil {
		if cerr := q.getAllAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllAccountsStmt: %w", cerr)
		}
	}
	if q.getAllPlaylistsStmt != nil {
		if cerr := q.getAllPlaylistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllPlaylistsStmt: %w", cerr)
		}
	}
	if q.getAllSongsStmt != nil {
		if cerr := q.getAllSongsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllSongsStmt: %w", cerr)
		}
	}
	if q.getCommentsBySongIdStmt != nil {
		if cerr := q.getCommentsBySongIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCommentsBySongIdStmt: %w", cerr)
		}
	}
	if q.getFollowersStmt != nil {
		if cerr := q.getFollowersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFollowersStmt: %w", cerr)
		}
	}
	if q.getFollowingStmt != nil {
		if cerr := q.getFollowingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFollowingStmt: %w", cerr)
		}
	}
	if q.getLikesBySongIdStmt != nil {
		if cerr := q.getLikesBySongIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLikesBySongIdStmt: %w", cerr)
		}
	}
	if q.getListenHistoryByAccountIdStmt != nil {
		if cerr := q.getListenHistoryByAccountIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListenHistoryByAccountIdStmt: %w", cerr)
		}
	}
	if q.getPlaylistByIdStmt != nil {
		if cerr := q.getPlaylistByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlaylistByIdStmt: %w", cerr)
		}
	}
	if q.getReportsByAccountIdStmt != nil {
		if cerr := q.getReportsByAccountIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getReportsByAccountIdStmt: %w", cerr)
		}
	}
	if q.getReportsBySongIdStmt != nil {
		if cerr := q.getReportsBySongIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getReportsBySongIdStmt: %w", cerr)
		}
	}
	if q.getRepostsByAccountIdStmt != nil {
		if cerr := q.getRepostsByAccountIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRepostsByAccountIdStmt: %w", cerr)
		}
	}
	if q.getRepostsBySongIdStmt != nil {
		if cerr := q.getRepostsBySongIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRepostsBySongIdStmt: %w", cerr)
		}
	}
	if q.getSongByIdStmt != nil {
		if cerr := q.getSongByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSongByIdStmt: %w", cerr)
		}
	}
	if q.getSongsByPlaylistIdStmt != nil {
		if cerr := q.getSongsByPlaylistIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSongsByPlaylistIdStmt: %w", cerr)
		}
	}
	if q.getSubscriptionsByAccountIdStmt != nil {
		if cerr := q.getSubscriptionsByAccountIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSubscriptionsByAccountIdStmt: %w", cerr)
		}
	}
	if q.likeSongStmt != nil {
		if cerr := q.likeSongStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing likeSongStmt: %w", cerr)
		}
	}
	if q.removeSongFromPlaylistStmt != nil {
		if cerr := q.removeSongFromPlaylistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeSongFromPlaylistStmt: %w", cerr)
		}
	}
	if q.repostSongStmt != nil {
		if cerr := q.repostSongStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing repostSongStmt: %w", cerr)
		}
	}
	if q.unfollowUserStmt != nil {
		if cerr := q.unfollowUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unfollowUserStmt: %w", cerr)
		}
	}
	if q.unlikeSongStmt != nil {
		if cerr := q.unlikeSongStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unlikeSongStmt: %w", cerr)
		}
	}
	if q.updateAccountStmt != nil {
		if cerr := q.updateAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountStmt: %w", cerr)
		}
	}
	if q.updatePlaylistStmt != nil {
		if cerr := q.updatePlaylistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePlaylistStmt: %w", cerr)
		}
	}
	if q.updateSongStmt != nil {
		if cerr := q.updateSongStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSongStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                              DBTX
	tx                              *sql.Tx
	addListenHistoryStmt            *sql.Stmt
	addSongToPlaylistStmt           *sql.Stmt
	cancelSubscriptionStmt          *sql.Stmt
	createAccountStmt               *sql.Stmt
	createCommentStmt               *sql.Stmt
	createPlaylistStmt              *sql.Stmt
	createReportStmt                *sql.Stmt
	createSongStmt                  *sql.Stmt
	createSubscriptionStmt          *sql.Stmt
	deleteAccountStmt               *sql.Stmt
	deleteCommentStmt               *sql.Stmt
	deletePlaylistStmt              *sql.Stmt
	deleteReportStmt                *sql.Stmt
	deleteRepostStmt                *sql.Stmt
	deleteSongStmt                  *sql.Stmt
	followUserStmt                  *sql.Stmt
	getAccountByIdStmt              *sql.Stmt
	getAllAccountsStmt              *sql.Stmt
	getAllPlaylistsStmt             *sql.Stmt
	getAllSongsStmt                 *sql.Stmt
	getCommentsBySongIdStmt         *sql.Stmt
	getFollowersStmt                *sql.Stmt
	getFollowingStmt                *sql.Stmt
	getLikesBySongIdStmt            *sql.Stmt
	getListenHistoryByAccountIdStmt *sql.Stmt
	getPlaylistByIdStmt             *sql.Stmt
	getReportsByAccountIdStmt       *sql.Stmt
	getReportsBySongIdStmt          *sql.Stmt
	getRepostsByAccountIdStmt       *sql.Stmt
	getRepostsBySongIdStmt          *sql.Stmt
	getSongByIdStmt                 *sql.Stmt
	getSongsByPlaylistIdStmt        *sql.Stmt
	getSubscriptionsByAccountIdStmt *sql.Stmt
	likeSongStmt                    *sql.Stmt
	removeSongFromPlaylistStmt      *sql.Stmt
	repostSongStmt                  *sql.Stmt
	unfollowUserStmt                *sql.Stmt
	unlikeSongStmt                  *sql.Stmt
	updateAccountStmt               *sql.Stmt
	updatePlaylistStmt              *sql.Stmt
	updateSongStmt                  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                              tx,
		tx:                              tx,
		addListenHistoryStmt:            q.addListenHistoryStmt,
		addSongToPlaylistStmt:           q.addSongToPlaylistStmt,
		cancelSubscriptionStmt:          q.cancelSubscriptionStmt,
		createAccountStmt:               q.createAccountStmt,
		createCommentStmt:               q.createCommentStmt,
		createPlaylistStmt:              q.createPlaylistStmt,
		createReportStmt:                q.createReportStmt,
		createSongStmt:                  q.createSongStmt,
		createSubscriptionStmt:          q.createSubscriptionStmt,
		deleteAccountStmt:               q.deleteAccountStmt,
		deleteCommentStmt:               q.deleteCommentStmt,
		deletePlaylistStmt:              q.deletePlaylistStmt,
		deleteReportStmt:                q.deleteReportStmt,
		deleteRepostStmt:                q.deleteRepostStmt,
		deleteSongStmt:                  q.deleteSongStmt,
		followUserStmt:                  q.followUserStmt,
		getAccountByIdStmt:              q.getAccountByIdStmt,
		getAllAccountsStmt:              q.getAllAccountsStmt,
		getAllPlaylistsStmt:             q.getAllPlaylistsStmt,
		getAllSongsStmt:                 q.getAllSongsStmt,
		getCommentsBySongIdStmt:         q.getCommentsBySongIdStmt,
		getFollowersStmt:                q.getFollowersStmt,
		getFollowingStmt:                q.getFollowingStmt,
		getLikesBySongIdStmt:            q.getLikesBySongIdStmt,
		getListenHistoryByAccountIdStmt: q.getListenHistoryByAccountIdStmt,
		getPlaylistByIdStmt:             q.getPlaylistByIdStmt,
		getReportsByAccountIdStmt:       q.getReportsByAccountIdStmt,
		getReportsBySongIdStmt:          q.getReportsBySongIdStmt,
		getRepostsByAccountIdStmt:       q.getRepostsByAccountIdStmt,
		getRepostsBySongIdStmt:          q.getRepostsBySongIdStmt,
		getSongByIdStmt:                 q.getSongByIdStmt,
		getSongsByPlaylistIdStmt:        q.getSongsByPlaylistIdStmt,
		getSubscriptionsByAccountIdStmt: q.getSubscriptionsByAccountIdStmt,
		likeSongStmt:                    q.likeSongStmt,
		removeSongFromPlaylistStmt:      q.removeSongFromPlaylistStmt,
		repostSongStmt:                  q.repostSongStmt,
		unfollowUserStmt:                q.unfollowUserStmt,
		unlikeSongStmt:                  q.unlikeSongStmt,
		updateAccountStmt:               q.updateAccountStmt,
		updatePlaylistStmt:              q.updatePlaylistStmt,
		updateSongStmt:                  q.updateSongStmt,
	}
}
