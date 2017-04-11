# amperApi
amper music api 

@app.route('/api/grafana/render', methods=['POST'])
def query_metric_values():
	result = []
	url = onfig.QUERY_PROXY_ADDR + "/graph/history"
	start = int(request.form.get("from"))
	end = int(request.form.get("until"))
	step = int(request.form.get("step"))
	if (end - start) <= 86400:
		step = 60
	elif (end - start) <= 604800:
		step = 300
	elif (end - start) <= 2592000:
		step = 600
	elif (end - start) <= 7776000:
		step = 1800
	else :
		step = 3600
	for target  in request.form.get("target"):
		if ".select metric" in target:
			targets  = string.split(target,"#")
			host = targets[0]
			targets = targets[1:]
			result.append(getMetricValues(url,start,end,step,host,targets))
	
	return data

def getMetricValues(url,start,until,format,step,host,targets,result):
	endpoint_counters = []
	ret = []
	metric = strings.join(targets,".")
	if "{" in host:
		host = string.replace(host,"{","",-1)
		host = string.replace(host,"}","",-1)
		hosts = string.split(host,",")
		for h in hosts:
			item = dict()
			item["endpoint"] = h
			item["counter"] = metric
			endpoint_counters.append(item)
	else:
		item = dict()
		item["endpoint"] = h
		item["counter"] = metric
		endpoint_counters.append(item)
	
	if len(endpoint_counters) > 0:
		args = {
			"start":start,
			"end":end,
			"cf":"AVERAGE",
			"endpoint_counters":endpoint_counters,
			"step":step,
		}
		header = {"Content-Type":"application/x-www-form-urlencoded"}
		resp = requests.post(url,data=args,headers=header)
		if resp.status_code == 200:
			return resp.content
	return ret


def getNextCounterSegment(metric,counter):
	if len(metric) > 0 :
		metric = metric + "."
	
	counter = string.replace(counter,metric,"",1)
	segment = string.split(counter,".")[0]
	return segment

def checkSegmentExpandable(segment,counter):
	segments = string.split(counter,".")
	if segment == segments[len(segments)-1]:
		expandable = False
	else:
		expandable = True
	return expandable


@app.route('/api/grafana/metrics/find', methods=['GET'])
def query_hosts_metric():
	querylist = string.split(request.url,"?")
	query = querylist[1]
	query = string.replace(query,".%","",-1)
	query = string.replace(query,".undefined","",-1)
	query = string.replace(query,".select metric","",-1)
	
	if "#" in query:
		query = string.replace(query,"#.*","",-1)
		arrQuery = string.split(query,"#")
		host = arrQuery[0]
		arrMetric = arrQuery[1]
		maxQuery = 1000
		metrics = string.join(arrMetric,".")
		reg = re.compile("(^{|})")
		host = reg.sub(host,"")
		host = string.replace(host,",","\",\"",-1)
		endpoints = "[\""+host+"\"]"
		r = random.uniform(10, 20)
		formdata = {
			"endpoints":endpoints,
			"q":metrics,
			"limit":maxQuery,
			"_r":r,
		}
		result = []
		header = {"Content-Type":"application/x-www-form-urlencoded"}
		url = config.QUERY_PROXY_ADDR+"/api/counters"
		resp = requests.post(url,data=formdata,headers=header)
		if resp.status_code == 200:
			segmentPool = dict()
			for data in json.loads(resp.json)["data"]:
				counter = string(data[0])
				segment = getNextCounterSegment(metric,counter)
				expandable = checkSegmentExpandable(segment,counter)
				if False == segmentPool[segment]:
					item = dict()
					item["text"] = segment
					item["expandable"] = expandable
					result.appenditem
		return json.dumps(result)
	else:
		if len(query) == 1:
			query = ".+"
		r = random.uniform(10, 20)
		url =config.QUERY_PROXY_ADDR+"/api/endpoints"+"?q"+query+"&tags&limit=1000&_r="+string(r)+"&regex_qeury=1"
		resp = requests.get(url)
		result = []
		if resp.status_code == 200:
			for node in json.loads(resp.json)["data"]:
				item = dict()
				item["text"] = node
				item["expandable"]=true
				result.append(item)
			
		return json.dumps(result)

