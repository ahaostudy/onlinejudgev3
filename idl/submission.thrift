namespace go submit

include "./base.thrift"

struct JudgeResult {
    1: i64 time
    2: i64 memory
    3: i64 status
    4: string message
    5: string output
    6: string error
}

struct Submit {
    1: i64 id
    2: i64 userId
    3: i64 problemId
    4: string code
    5: i64 langId
    6: i64 status
    7: i64 time
    8: i64 memory
    9: i64 noteId
    10: i64 createdAt
}

struct DebugRequest {
    1: binary code
    2: binary input
    3: i64 langId
}

struct DebugResponse {
    1: JudgeResult result
}

struct SubmitRequest {
    1: i64 problemId
    2: binary code
    3: i64 langId
    4: i64 userId
    5: required i64 contestId
}

struct SubmitResponse {
    1: i64 submitId
}

struct GetSubmitResultRequest {
    1: i64 submitId
}

struct GetSubmitResultResponse {
    1: JudgeResult result
}

struct GetSubmitListRequest {
    1: i64 userId
    2: i64 problemId
}

struct GetSubmitListResponse {
    1: list<Submit> submitList
}

struct GetSubmitRequest {
    1: i64 id
}

struct GetSubmitResponse {
    1: Submit submit
}

struct SubmitStatus {
    1: i64 count
    2: i64 acceptedCount
}

struct GetSubmitStatusRequest {
}

struct GetSubmitStatusResponse {
    1: map<i64, SubmitStatus> submitStatus
}

struct IsAcceptedRequest {
    1: i64 userId
    2: i64 problemId
}

struct IsAcceptedResponse {
    1: bool isAccepted
}

struct GetAcceptedStatusRequest {
    1: i64 userId
}

struct GetAcceptedStatusResponse {
    1: map<i64, bool> acceptedStatus
}

struct GetLatestSubmitsRequest {
    1: i64 userId
    2: i64 count
}

struct GetLatestSubmitsResponse {
    1: list<Submit> submitList
}

struct DeleteSubmitRequest {
    1: i64 id
    2: i64 userId
}

struct GetSubmitCalendarRequest {
    1: i64 userId
}

struct GetSubmitCalendarResponse {
    1: map<string, i64> submitCalendar
}

struct GetSubmitStatisticsRequest {
    1: i64 userId
}

struct GetSubmitStatisticsResponse {
    1: i64 sloveCount
    2: i64 submitCount
    3: i64 easyCount
    4: i64 middleCount
    5: i64 hardCount
    6: map<i64, i64> langCounts
}

service SubmitService {
    DebugResponse Debug(1: DebugRequest request)
    SubmitResponse Submit(1: SubmitRequest request)
    GetSubmitResultResponse GetSubmitResult(1: GetSubmitResultRequest request)

    GetSubmitResponse GetSubmit(1: GetSubmitRequest request)
    GetSubmitListResponse GetSubmitList(1: GetSubmitListRequest request)

    GetSubmitStatusResponse GetSubmitStatus(1: GetSubmitStatusRequest request)
    IsAcceptedResponse IsAccepted(1: IsAcceptedRequest request)
    GetAcceptedStatusResponse GetAcceptedStatus(1: GetAcceptedStatusRequest request)
    GetLatestSubmitsResponse GetLatestSubmits(1: GetLatestSubmitsRequest request)
    base.Empty DeleteSubmit(1: DeleteSubmitRequest request)
    GetSubmitCalendarResponse GetSubmitCalendar(1: GetSubmitCalendarRequest request)
    GetSubmitStatisticsResponse GetSubmitStatistics(1: GetSubmitStatisticsRequest request)
}