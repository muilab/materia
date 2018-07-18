import StatusBar from "elements/status-bar/status-bar"
import delay from "./delay"
import requestIdleCallback from "./requestIdleCallback"

const newVersionCheckDelay = location.hostname.includes("beta.") ? 3000 : 60000

let etags = new Map<string, string>()
let hasNewVersion = false

async function checkNewVersion(url: string, status: StatusBar) {
	if(hasNewVersion) {
		return
	}

	try {
		let headers = {}

		if(etags.has(url)) {
			headers["If-None-Match"] = etags.get(url)
		}

		let response = await fetch(url, {
			headers,
			credentials: "omit"
		})

		// Not modified response
		if(response.status === 304) {
			return
		}

		if(!response.ok) {
			console.warn("Error fetching", url, response.status)
			return
		}

		let newETag = response.headers.get("ETag")
		let oldETag = etags.get(url)

		if(newETag) {
			etags.set(url, newETag)
		}

		if(oldETag && newETag && oldETag !== newETag) {
			status.showInfo("A new version of the website is available. Please refresh the page.", -1)

			// Do not check for new versions again.
			hasNewVersion = true
			return
		}
	} catch(err) {
		console.warn("Error fetching", url + "\n", err)
	} finally {
		checkNewVersionDelayed(url, status)
	}
}

export default function checkNewVersionDelayed(url: string, status: StatusBar) {
	return delay(newVersionCheckDelay).then(() => {
		requestIdleCallback(() => checkNewVersion(url, status))
	})
}
