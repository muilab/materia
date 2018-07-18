export default class NetworkManager {
	isLoading: boolean
	loadingStateChangeCallbacks: Array<Function>

	constructor() {
		this.loadingStateChangeCallbacks = new Array<Function>()
	}

	onLoadingStateChange(callback: Function) {
		this.loadingStateChangeCallbacks.push(callback)
	}

	post(url: string, body?: any) {
		if(this.loading) {
			return Promise.resolve(null)
		}
	
		if(body !== undefined && typeof body !== "string") {
			body = JSON.stringify(body)
		}
	
		this.loading = true
	
		return fetch(url, {
			method: "POST",
			body,
			credentials: "same-origin"
		})
		.then(response => {
			this.loading = false
	
			if(response.status === 200) {
				return Promise.resolve(response)
			}
	
			return response.text().then(err => {
				throw err
			})
		})
		.catch(err => {
			this.loading = false
			throw err
		})
	}

	get loading() {
		return this.isLoading
	}

	set loading(state: boolean) {
		if(this.isLoading === state) {
			return
		}

		this.isLoading = state

		for(let callback of this.loadingStateChangeCallbacks) {
			callback()
		}
	}
}