-- Account table
CREATE TABLE Account (
    account_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    display_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    bio TEXT,
    avatar_url VARCHAR(255),
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    role VARCHAR(50) NOT NULL DEFAULT 'user', -- 'user', 'artist', 'admin'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_username ON Account (username);
CREATE INDEX idx_email ON Account (email);

-- Song table
CREATE TABLE Song (
    song_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration TIME NOT NULL,
    file_url VARCHAR(255) NOT NULL,
    account_id INT NOT NULL,
    album_id INT,
    genre VARCHAR(100),
    tags TEXT[],
    upload_date DATE,
    play_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    repost_count INT DEFAULT 0,
    is_public BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (account_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_title ON Song (title);
CREATE INDEX idx_account_id ON Song (account_id);
CREATE INDEX idx_genre ON Song (genre);

-- Album table
CREATE TABLE Album (
    album_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    account_id INT NOT NULL,
    release_date DATE,
    is_public BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (account_id) REFERENCES Account(account_id)
);

-- Playlist table
CREATE TABLE Playlist (
    playlist_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    account_id INT NOT NULL,
    is_public BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_account_id_playlist ON Playlist (account_id);

-- Playlist_Songs table
CREATE TABLE Playlist_Songs (
    playlist_id INT NOT NULL,
    song_id INT NOT NULL,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (playlist_id, song_id),
    FOREIGN KEY (playlist_id) REFERENCES Playlist(playlist_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_playlist_id ON Playlist_Songs (playlist_id);
CREATE INDEX idx_song_id ON Playlist_Songs (song_id);

-- Follow table
CREATE TABLE Follow (
    follower_id INT NOT NULL,
    following_id INT NOT NULL,
    followed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, following_id),
    FOREIGN KEY (follower_id) REFERENCES Account(account_id),
    FOREIGN KEY (following_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_follower_id ON Follow (follower_id);
CREATE INDEX idx_following_id ON Follow (following_id);

-- Comment table
CREATE TABLE Comment (
    comment_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    song_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_comment ON Comment (account_id);
CREATE INDEX idx_song_id_comment ON Comment (song_id);

-- Like table
CREATE TABLE "Like" (
    account_id INT NOT NULL,
    song_id INT NOT NULL,
    liked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, song_id),
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_like ON "Like" (account_id);
CREATE INDEX idx_song_id_like ON "Like" (song_id);

-- Listen_History table
CREATE TABLE Listen_History (
    account_id INT NOT NULL,
    song_id INT NOT NULL,
    listened_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, song_id),
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_history ON Listen_History (account_id);
CREATE INDEX idx_song_id_history ON Listen_History (song_id);

-- Subscription table
CREATE TABLE Subscription (
    subscription_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    plan_type VARCHAR(50) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    FOREIGN KEY (account_id) REFERENCES Account(account_id)
);

CREATE INDEX idx_account_id_subscription ON Subscription (account_id);

-- Report table
CREATE TABLE Report (
    report_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    song_id INT NOT NULL,
    reason TEXT NOT NULL,
    reported_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_report ON Report (account_id);
CREATE INDEX idx_song_id_report ON Report (song_id);

-- Repost table
CREATE TABLE Repost (
    repost_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    song_id INT NOT NULL,
    reposted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (song_id) REFERENCES Song(song_id)
);

CREATE INDEX idx_account_id_repost ON Repost (account_id);
CREATE INDEX idx_song_id_repost ON Repost (song_id);
