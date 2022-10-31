
// base url for all urls to go off of
const baseURL = "http://localhost:8080";

const urls = {
	"CurrentData": "/api/data/day",
};

// get the url given its name
export function UrlGet(name) {
	return baseURL+urls[name];
}


// headers sent when doing a POST operation
const headers = {
	'Content-Type': 'application/json',
}

// returns the default options used when sending a post request
// NOTE: body is sent as is, make sure JSONify if required first
export function GetPostOptions(body) {
	return {
		method: 'POST',
		headers: headers,
		body: body,
	};
}

export default UrlGet;
