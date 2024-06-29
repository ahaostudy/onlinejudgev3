namespace go judgesvc

enum Language {
	CPP
	C
	Python3
	Java
	Go
}

enum JudgeStatus {
	Accepted
	Finished
	Running
	CompileError
	RuntimeError
	WrongAnswer
	TimeLimitExceeded
	MemoryLimitExceeded
	OutputLimitExceeded
}

struct JudgeLimit {
    1: i32 maxCPUTime;
    2: i32 maxRealTime;
    3: i64 maxMemory;
    4: i64 maxStack;
    5: i32 maxProcessNumber;
    6: i64 maxOutputSize;
    7: i32 memoryLimitCheckOnly;
}

struct JudgeRequest {
    1: optional binary code;
    2: optional string codeFileId;
    3: required binary input;
    4: required Language language;
    5: optional JudgeLimit limit;
}

struct JudgeResponse {
    1: i64 time;
    2: i64 memory;
    3: JudgeStatus status;
    4: string output;
    5: string error;
}

struct UploadCodeRequest {
    1: required binary code;
    2: required Language language;
}

struct UploadCodeResponse {
    1: required string fileId;
}

struct DeleteCodeRequest {
    1: required string fileId;
}

struct DeleteCodeResponse {}

service JudgeService {
    JudgeResponse Judge(1: JudgeRequest req)
    UploadCodeResponse UploadCode(1: UploadCodeRequest req)
    DeleteCodeResponse DeleteCode(1: DeleteCodeRequest req)
}