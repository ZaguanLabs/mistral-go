#!/usr/bin/env zsh
# Run all local tests, excluding the legacy tests that contact the live API.
emulate -R zsh
setopt ERR_EXIT NO_UNSET PIPE_FAIL
cd -- "${0:A:h:h}"

# Exact names keep mock variants and new regression tests enabled.
local -a live_tests=(
 TestFIM TestFIMWithStop TestFIMInvalidModel
 TestChat TestChatCodestral TestChatFunctionCall TestChatFunctionCall2
 TestChatJsonMode TestChatStream TestChatStreamFunctionCall TestChatStreamJsonMode
 TestUploadFile TestListFiles TestListFilesWithFilters TestListFilesNilParams
 TestRetrieveFile TestDeleteFile TestDownloadFile TestGetSignedURL
 TestGetSignedURLWithExpiry TestUploadFileWithEmptyContent TestUploadFileWithLargeFilename
 TestCreateFineTuningJob TestCreateFineTuningJobWithAllParams
 TestListFineTuningJobs TestListFineTuningJobsWithFilters TestListFineTuningJobsNilParams
 TestGetFineTuningJob TestCancelFineTuningJob TestStartFineTuningJob
 TestCreateBatchJob TestCreateBatchJobWithAllParams TestListBatchJobs
 TestListBatchJobsWithFilters TestListBatchJobsNilParams TestGetBatchJob TestCancelBatchJob
 TestEmbeddings TestListModels
)
exec go test -race ./... -count=1 -skip "^(${(j:|:)live_tests})$" "$@"
