cd w12/4_jwt_sessions

revise main.go - sm := NewSessionsDB(db)

session_common.go

session_db.go

user.go - change password

p_1.png - p_8.png





revise main.go - sm := NewSessionsJWT("golangcourseSessionSecret")

session_jwt.go

compare with session_db.go

revise main.go - sm := NewSessionsJWTVer("golangcourseSessionSecret", db)

session_jwt_ver.go

demostrate session expiration with VER and without VER
