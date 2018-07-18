import { findAll } from "./findAll";

export default function assignActions() {
	for(let element of findAll("action")) {
		let actionTrigger = element.dataset.trigger
		let actionName = element.dataset.action

		// Filter out invalid definitions
		if(!actionTrigger || !actionName) {
			continue
		}

		let oldAction = element["action assigned"]

		if(oldAction) {
			if(oldAction.trigger === actionTrigger && oldAction.action === actionName) {
				continue
			}

			element.removeEventListener(oldAction.trigger, oldAction.handler)
		}

		// This prevents default actions on links
		if(actionTrigger === "click" && element.tagName === "A") {
			element.onclick = null
		}

		// Warn us about undefined actions
		if(!(actionName in actions)) {
			this.statusMessage.showError(`Action '${actionName}' has not been defined`)
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