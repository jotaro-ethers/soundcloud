-- Bảng Account
CREATE TABLE Account (
    account_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    bio TEXT,
    avatar_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_username ON Account (username);
CREATE INDEX idx_email ON Account (email);

-- Bảng Song
CREATE TABLE Song (
    song_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration TIME NOT NULL,
    file_url VARCHAR(255) NOT NULL,
    account_id INT,
    genre VARCHAR(100),
    upload_date DATE,
    play_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    FOREIGN KEY (account_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_title ON Song (title);
CREATE INDEX idx_account_id ON Song (account_id);
CREATE INDEX idx_genre ON Song (genre);

-- Bảng Playlist
CREATE TABLE Playlist (
    playlist_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    account_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_account_id_playlist ON Playlist (account_id);

-- Bảng Playlist_Songs
CREATE TABLE Playlist_Songs (
    playlist_id INT,
    song_id INT,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (playlist_id, song_id),
    FOREIGN KEY (playlist_id) REFERENCES Playlist(playlist_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_playlist_id ON Playlist_Songs (playlist_id);
CREATE INDEX idx_song_id ON Playlist_Songs (song_id);

-- Bảng Follow
CREATE TABLE Follow (
    follower_id INT,
    following_id INT,
    followed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, following_id),
    FOREIGN KEY (follower_id) REFERENCES Account(account_id),
    FOREIGN KEY (following_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_follower_id ON Follow (follower_id);
CREATE INDEX idx_following_id ON Follow (following_id);

-- Bảng Comment
CREATE TABLE Comment (
    comment_id SERIAL PRIMARY KEY,
    account_id INT,
    song_id INT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_comment ON Comment (account_id);
CREATE INDEX idx_song_id_comment ON Comment (song_id);

-- Bảng Like
CREATE TABLE "Like" (
    account_id INT,
    song_id INT,
    liked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, song_id),
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_like ON "Like" (account_id);
CREATE INDEX idx_song_id_like ON "Like" (song_id);

-- Bảng Listen_History
CREATE TABLE Listen_History (
    account_id INT,
    song_id INT,
    listened_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, song_id),
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_history ON Listen_History (account_id);
CREATE INDEX idx_song_id_history ON Listen_History (song_id);

-- Bảng Subscription
CREATE TABLE Subscription (
    subscription_id SERIAL PRIMARY KEY,
    account_id INT,
    plan_type VARCHAR(50),
    start_date DATE,
    end_date DATE,
    FOREIGN KEY (account_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_account_id_subscription ON Subscription (account_id);

-- Bảng Report
CREATE TABLE Report (
    report_id SERIAL PRIMARY KEY,
    account_id INT,
    song_id INT,
    reason TEXT,
    reported_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_report ON Report (account_id);
CREATE INDEX idx_song_id_report ON Report (song_id);
