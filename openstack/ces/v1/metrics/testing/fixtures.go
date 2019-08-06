package testing

const ListResponse = `
{
	"metrics": [{
		"namespace": "SYS.ECS",
		"dimensions": [{
			"name": "instance_id",
			"value": "d9112af5-6913-4f3b-bd0a-3f96711e004d"
		}],
		"metric_name": "cpu_util",
		"unit": "%"
	}],
	"meta_data": {
		"count": 1,
		"marker": "SYS.ECS.cpu_util.instance_id:d9112af5-6913-4f3b-bd0a-3f96711e004d",
		"total": 2
	}
}
`
const EndPageResponse = `
{
  "metrics": [],
  "meta_data": {
    "count": 0,
    "marker": "",
    "total": 2
  }
}
`