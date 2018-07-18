import Application from "./Application"
import NetworkManager from "./NetworkManager"
import SVGIcon from "elements/svg-icon/svg-icon"
import StatusBar from "elements/status-bar/status-bar"
import findAll from "./Utils/findAll"
import mountMountables from "./Utils/mountMountables"
import requestIdleCallback from "./Utils/requestIdleCallback"
import checkNewVersionDelayed from "./Utils/checkNewVersionDelayed"
import * as actions from "./Actions"
import Diff from "./Diff"

export default class Materia {
	app: Application
	title: string
	network: NetworkManager
	status: StatusBar

	constructor(app: Application) {
		this.app = app
		this.title = "mui Materia"
	}

	init() {
		// App init
		this.app.init()

		// Event listeners
		document.addEventListener("readystatechange", this.onReadyStateChange.bind(this))
		document.addEventListener("DOMContentLoaded", this.onContentLoaded.bind(this))

		// If we finished loading the DOM (either "interactive" or "complete" state),
		// immediately trigger the event listener functions.
		if(document.readyState !== "loading") {
			this.app.emit("DOMContentLoaded")
			this.run()
		}

		// Idle
		requestIdleCallback(this.onIdle.bind(this))
	}

	onReadyStateChange() {
		if(document.readyState !== "interactive") {
			return
		}

		this.run()
	}

	onContentLoaded() {
		mountMountables()
		this.assignActions()
		this.applyPageTitle()
	}

	run() {
		this.app.content = document.getElementById("content")
		this.app.loading = document.getElementById("loading")
		this.status = document.getElementsByTagName("status-bar")[0] as StatusBar

		this.registerWebComponents()

		// Network manager
		this.network = new NetworkManager()
		this.network.onLoadingStateChange(loading => {
			if(loading) {
				this.app.loading.classList.remove("fade-out")
			} else {
				this.app.loading.classList.add("fade-out")
			}
		})

		// Fade out loading animation
		this.app.loading.classList.add("fade-out")
	}

	onIdle() {
		// Register event listeners
		document.addEventListener("keydown", this.onKeyDown.bind(this), false)
		window.addEventListener("popstate", this.onPopState.bind(this))
		window.addEventListener("error", this.onError.bind(this))

		// Service worker
		// this.serviceWorkerManager = new ServiceWorkerManager(this, "/service-worker")
		// this.serviceWorkerManager.register()

		// Periodically check etags of scripts and styles to let the user know about page updates
		checkNewVersionDelayed("/scripts", this.status)
		checkNewVersionDelayed("/styles", this.status)
	}

	onKeyDown(e: KeyboardEvent) {
		// ...
	}

	// This is called every time an uncaught JavaScript error is thrown
	onError(evt: ErrorEvent) {
		let report = {
			message: evt.message,
			stack: evt.error.stack,
			fileName: evt.filename,
			lineNumber: evt.lineno,
			columnNumber: evt.colno,
			userAgent: navigator.userAgent,
		}

		console.log("Error report:", report)

		// this.network.post("/api/new/clienterrorreport", report)
		// .then(() => console.log("Successfully reported the error to the website staff."))
		// .catch(() => console.warn("Failed reporting the error to the website staff."))
	}

	registerWebComponents() {
		if(!("customElements" in window)) {
			console.warn("Web components not supported in your current browser")
			return
		}

		// Custom element names must have a dash in their name
		const elements = new Map<string, Function>([
			["svg-icon", SVGIcon],
			["status-bar", StatusBar]
		])

		// Register all custom elements
		for(const [tag, definition] of elements.entries()) {
			window.customElements.define(tag, definition)
		}
	}

	reloadContent() {
		let headers = new Headers()
		let path = location.pathname

		return fetch("/_" + path, {
			credentials: "same-origin",
			headers
		})
		.then(response => {
			if(location.pathname !== path) {
				return Promise.reject("old request")
			}

			return Promise.resolve(response)
		})
		.then(response => response.text())
		.then(html => Diff.innerHTML(this.app.content, html))
		.then(() => this.app.emit("DOMContentLoaded"))
	}

	applyPageTitle() {
		let header = document.querySelector("h1")

		if(!header) {
			document.title = this.title
		} else {
			document.title = header.textContent
		}
	}

	assignActions() {
		for(let element of findAll("action")) {
			let actionTrigger = element.dataset.trigger
			let actionName = element.dataset.action
	
			// Filter out invalid definitions
			if(!actionTrigger || !actionName) {
				continue
			}
	
			let oldAction = element["action assigned"]
	
			if(oldAction) {
				// If the action assigned is the exact same, skip this element
				if(oldAction.trigger === actionTrigger && oldAction.action === actionName) {
					continue
				}
	
				// Otherwise remove the existing event listener and continue to assign
				element.removeEventListener(oldAction.trigger, oldAction.handler)
			}
	
			// This prevents default actions on links
			if(actionTrigger === "click" && element.tagName === "A") {
				element.onclick = null
			}
	
			// Warn us about undefined actions
			if(!(actionName in actions)) {
				this.status.showError(`Action '${actionName}' has not been defined`)
				continue
			}
	
			// Register the actual action handler
			let actionHandler = e => {
				actions[actionName](this, element, e)
	
				e.stopPropagation()
				e.preventDefault()
			}
	
			element.addEventListener(actionTrigger, actionHandler)
	
			// Use "action assigned" flag instead of removing the class.
			// This will make sure that DOM diffs which restore the class name
			// will not assign the action multiple times to the same element.
			element["action assigned"] = {
				trigger: actionTrigger,
				action: actionName,
				handler: actionHandler
			}
		}
	}

	onPopState(e: PopStateEvent) {
		this.app.load(location.pathname, {
			addToHistory: false
		})
	}
}