-- Drop tables if they exist
DROP TABLE IF EXISTS Reaction;
DROP TABLE IF EXISTS Comment;
DROP TABLE IF EXISTS PostCategory;
DROP TABLE IF EXISTS Post;
DROP TABLE IF EXISTS UserSession;
DROP TABLE IF EXISTS UserRole;
DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Category;

-- Create tables
CREATE TABLE Category (
    Category_ID INTEGER NOT NULL PRIMARY KEY,
    Name VARCHAR NOT NULL,
    Description VARCHAR NOT NULL
);

CREATE TABLE User (
    User_ID INTEGER NOT NULL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    passwords TEXT NOT NULL, 
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE UserRole (
    UserRoleID INTEGER NOT NULL PRIMARY KEY,
    User_ID INTEGER NOT NULL,
    RoleName VARCHAR NOT NULL,
    CanPost VARCHAR NOT NULL,
    CanLike VARCHAR NOT NULL,
    CanComment VARCHAR NOT NULL,
    canDislike VARCHAR NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES User(User_ID)
);

CREATE TABLE UserSession (
    UserSessionID VARCHAR NOT NULL PRIMARY KEY,
    User_ID INTEGER NOT NULL,
    Token CHAR NOT NULL,
    CreationTime DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ExpireTime DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES User(User_ID)
);

CREATE TABLE Post (
    PostID INTEGER NOT NULL PRIMARY KEY,
    User_ID INTEGER NOT NULL,
    Title VARCHAR NOT NULL,
    Messages VARCHAR NOT NULL,
    TimeofPost DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    Like_count INTEGER NOT NULL,
    DisLike_Count INTEGER NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES User(User_ID)
);

CREATE TABLE PostCategory (
    PostCategory_ID INTEGER NOT NULL PRIMARY KEY,
    Category_ID INTEGER NOT NULL,
    PostID INTEGER NOT NULL,
    FOREIGN KEY (Category_ID) REFERENCES Category(Category_ID),
    FOREIGN KEY (PostID) REFERENCES Post(PostID)
);

CREATE TABLE Comment (
    Comment_ID INTEGER NOT NULL PRIMARY KEY,
    User_ID INTEGER NOT NULL,
    PostID INTEGER NOT NULL,
    TimeofComment DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    message VARCHAR NOT NULL,
    Like_Count INTEGER NOT NULL,
    DisLike_Count INTEGER NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES User(User_ID),
    FOREIGN KEY (PostID) REFERENCES Post(PostID)
);

CREATE TABLE Reaction (
    ReactionID INTEGER NOT NULL PRIMARY KEY,
    User_ID INTEGER NOT NULL,
    PostID INTEGER,
    Comment_ID INTEGER,
    Type VARCHAR NOT NULL,
    TimeofReaction DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES User(User_ID),
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (Comment_ID) REFERENCES Comment(Comment_ID)
);

-- Insert data
INSERT INTO Category (Category_ID, Name, Description)
VALUES 
(1, 'General Discussion', 'A place for general topics and conversations'),
(2, 'Programming', 'Discussions about programming languages and software development'),
(3, 'Announcements', 'Official announcements and updates'),
(4, 'Off-Topic', 'Casual conversations not related to the main topics');

INSERT INTO User (User_ID, username, email, passwords)
VALUES 
(1, 'alice', 'alice@example.com', 'password123'),
(2, 'bob', 'bob@example.com', 'password123'),
(3, 'charlie', 'charlie@example.com', 'password123');

INSERT INTO UserRole (UserRoleID, User_ID, RoleName, CanPost, CanLike, CanComment, canDislike)
VALUES 
(1, 1, 'Admin', 'TRUE', 'TRUE', 'TRUE', 'TRUE'),  -- Alice as Admin
(2, 2, 'Moderator', 'TRUE', 'TRUE', 'TRUE', 'TRUE'),  -- Bob as Moderator
(3, 3, 'User', 'TRUE', 'TRUE', 'TRUE', 'TRUE');  -- Charlie as User

INSERT INTO UserSession (UserSessionID, User_ID, Token, CreationTime, ExpireTime)
VALUES 
('session1', 1, 'token123', '2024-08-21 10:00:00', '2024-08-21 12:00:00'),
('session2', 2, 'token456', '2024-08-21 10:30:00', '2024-08-21 12:30:00'),
('session3', 3, 'token789', '2024-08-21 11:00:00', '2024-08-21 13:00:00');

INSERT INTO Post (PostID, User_ID, Title, Messages, Like_count, DisLike_Count)
VALUES 
(1, 1, 'Welcome to the Forum', 'This is the first post on the forum!', 10, 0),
(2, 2, 'Programming Tips', 'Share your best programming tips here.', 8, 2),
(3, 3, 'Forum Rules', 'Please read the rules before posting.', 15, 1);

INSERT INTO PostCategory (PostCategory_ID, Category_ID, PostID)
VALUES 
(1, 1, 1),  -- Linking "Welcome to the Forum" to "General Discussion"
(2, 2, 2),  -- Linking "Programming Tips" to "Programming"
(3, 3, 3);  -- Linking "Forum Rules" to "Announcements"

INSERT INTO Comment (Comment_ID, User_ID, PostID, message, Like_Count, DisLike_Count)
VALUES 
(1, 2, 1, 'Great to be here!', 5, 0),  -- Bob commenting on Alice's post
(2, 3, 2, 'Thanks for sharing!', 3, 1),  -- Charlie commenting on Bob's post
(3, 1, 3, 'Everyone should read this.', 10, 0);  -- Alice commenting on Charlie's post

INSERT INTO Reaction (ReactionID, User_ID, PostID, Comment_ID, Type)
VALUES 
(1, 2, 1, NULL, 'Like'),  -- Bob likes Alice's post
(2, 3, 2, NULL, 'Dislike'),  -- Charlie dislikes Bob's post
(3, 1, NULL, 1, 'Like'),  -- Alice likes Bob's comment
(4, 2, NULL, 2, 'Like'),  -- Bob likes Charlie's comment
(5, 3, NULL, 3, 'Like');  -- Charlie likes Alice's comment



-- View data from Category table
SELECT * FROM Category;

-- View data from User table
SELECT * FROM User;

-- View data from UserRole table
SELECT * FROM UserRole;

-- View data from UserSession table
SELECT * FROM UserSession;

-- View data from Post table
SELECT * FROM Post;

-- View data from PostCategory table
SELECT * FROM PostCategory;

-- View data from Comment table
SELECT * FROM Comment;

-- View data from Reaction table
SELECT * FROM Reaction;