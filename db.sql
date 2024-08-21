
CREATE TABLE Category (
                Category_ID INTEGER NOT NULL,
                Name VARCHAR NOT NULL,
                Description VARCHAR NOT NULL,
                PRIMARY KEY (Category_ID)
);


CREATE TABLE User (
                User_ID INTEGER NOT NULL,
                username VARCHAR NOT NULL,
                email VARCHAR NOT NULL,
                password VARCHAR NOT NULL,
                createdAt DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
                PRIMARY KEY (User_ID)
);


CREATE TABLE UserRole (
                UserRoleID INTEGER NOT NULL,
                User_ID INTEGER NOT NULL,
                RoleName VARCHAR NOT NULL,
                CanPost VARCHAR NOT NULL,
                CanLike VARCHAR NOT NULL,
                CanComment VARCHAR NOT NULL,
                canDislike VARCHAR NOT NULL,
                PRIMARY KEY (UserRoleID)
);


CREATE TABLE UserSession (
                UserSessionID VARCHAR NOT NULL,
                User_ID INTEGER NOT NULL,
                Token CHAR NOT NULL,
                CreationTime DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
                ExpireTime DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
                PRIMARY KEY (UserSessionID)
);


CREATE TABLE Post (
                PostID INTEGER NOT NULL,
                User_ID INTEGER NOT NULL,
                Title VARCHAR NOT NULL,
                Messages VARCHAR NOT NULL,
                TimeofPost DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
                Like_count INTEGER NOT NULL,
                DisLike_Count INTEGER NOT NULL,
                PRIMARY KEY (PostID)
);


CREATE TABLE PostCategory (
                PostCategory_ID INTEGER NOT NULL,
                Category_ID INTEGER NOT NULL,
                PostID INTEGER NOT NULL,
                PRIMARY KEY (PostCategory_ID)
);


CREATE TABLE Comment (
                Comment_ID INTEGER NOT NULL,
                User_ID INTEGER NOT NULL,
                PostID INTEGER NOT NULL,
                TimeofComment DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
                message VARCHAR NOT NULL,
                Like_Count INTEGER NOT NULL,
                DisLike_Count INTEGER NOT NULL,
                PRIMARY KEY (Comment_ID)
);


CREATE TABLE Reaction (
                ReactionID INTEGER NOT NULL,
                User_ID INTEGER NOT NULL,
                PostID INTEGER,
                Comment_ID INTEGER,
                Type VARCHAR NOT NULL,
                TimeofReaction DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
                PRIMARY KEY (ReactionID)
);


ALTER TABLE PostCategory ADD CONSTRAINT category_postcategory_fk
FOREIGN KEY (Category_ID)
REFERENCES Category (Category_ID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE Post ADD CONSTRAINT user_post_fk
FOREIGN KEY (User_ID)
REFERENCES User (User_ID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE Comment ADD CONSTRAINT user_comment_fk
FOREIGN KEY (User_ID)
REFERENCES User (User_ID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE Reaction ADD CONSTRAINT user_reaction_fk
FOREIGN KEY (User_ID)
REFERENCES User (User_ID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE UserSession ADD CONSTRAINT user_usersession_fk
FOREIGN KEY (User_ID)
REFERENCES User (User_ID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE UserRole ADD CONSTRAINT user_userrole_fk
FOREIGN KEY (User_ID)
REFERENCES User (User_ID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE Comment ADD CONSTRAINT post_comment_fk
FOREIGN KEY (PostID)
REFERENCES Post (PostID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE PostCategory ADD CONSTRAINT post_postcategory_fk
FOREIGN KEY (PostID)
REFERENCES Post (PostID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE Reaction ADD CONSTRAINT post_reaction_fk
FOREIGN KEY (PostID)
REFERENCES Post (PostID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;

ALTER TABLE Reaction ADD CONSTRAINT comment_reaction_fk
FOREIGN KEY (Comment_ID)
REFERENCES Comment (Comment_ID)
ON DELETE NO ACTION
ON UPDATE NO ACTION;