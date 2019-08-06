package testing

const GetJobResultResponse = `
{
    "status": "SUCCESS",
    "entities": {
        "sub_jobs_total": 1,
        "sub_jobs": [
            {
                "status": "SUCCESS",
                "entities": {
                    "server_id": "bae51750-0089-41a1-9b18-5c777978ff6d"
                },
                "job_id": "2c9eb2c5544cbf6101544f0635672b60",
                "job_type": "createSingleServer",
                "begin_time": "2016-04-25T20:04:47.591Z",
                "end_time": "2016-04-25T20:08:21.328Z",
                "error_code": null,
                "fail_reason": null
            }
        ]
    },
    "job_id": "2c9eb2c5544cbf6101544f0602af2b4f",
    "job_type": "createServer",
    "begin_time": "2016-04-25T20:04:34.604Z",
    "end_time": "2016-04-25T20:08:41.593Z",
    "error_code": null,
    "fail_reason": null
}
`