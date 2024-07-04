namespace go usersvc

include "./base.thrift"

struct User {
    1: i64 id
    2: string nickname
    3: string username
    4: string password
    5: string email
    6: string avatar
    7: string signature
    8: i32 role
}

struct RegisterReq {
    1: string email
    2: string captcha
    3: string password
}

struct RegisterResp {
    1: i64 userId
}

struct LoginReq {
    1: optional string username
    2: optional string email
    3: optional string password
    4: optional string captcha
}

struct LoginResp {
    1: i64 userId
}

struct CreateUserReq {
    1: string nickname
    2: string username
    3: string password
    4: string email
    5: string avatar
    6: string signature
    7: i32 role
    8: i64 loggedInId
}

struct CreateUserResp {
    1: i64 userId
}

struct UpdateUserReq {
    1: i64 id
    2: i64 loggedInId
    3: optional string nickname
    4: optional string username
    5: optional string password
    6: optional string email
    7: optional string avatar
    8: optional string signature
    9: optional i32 role
}

struct GenCaptchaReq {
    1: string email
}

struct GenCaptchaResp {
    1: string captcha
}

struct GetUserReq {
    1: optional i64 id
    2: optional string username
}

struct GetUserResp {
    1: User user
}

struct GetUserListReq {
    1: list<i64> userIdList
}

struct GetUserListResp {
    1: list<User> userList
}

struct UploadAvatarReq {
    1: i64 userId
    2: binary body
    3: string ext
}

struct UploadAvatarResp {
    1: string avatar
}

struct DownloadAvatarReq {
    1: string avatar
}

struct DownloadAvatarResp {
    1: binary body
}

struct DeleteAvatarReq {
    1: i64 userid
}

service UserService {
    RegisterResp Register(1: RegisterReq req)
    LoginResp Login(1: LoginReq req)
    CreateUserResp CreateUser(1: CreateUserReq req)
    base.Empty UpdateUser(1: UpdateUserReq req)
    GenCaptchaResp GenCaptcha(1: GenCaptchaReq req)
    GetUserResp GetUser(1: GetUserReq req)
    GetUserListResp GetUserList(1: GetUserListReq req)
    UploadAvatarResp UploadAvatar(1: UploadAvatarReq req)
    DownloadAvatarResp DownloadAvatar(1: DownloadAvatarReq req)
    base.Empty DeleteAvatar(1: DeleteAvatarReq req)
}