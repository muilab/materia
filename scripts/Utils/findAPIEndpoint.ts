export default function findAPIEndpoint(element: HTMLElement) {
	if(element.dataset.api !== undefined) {
		return element.dataset.api
	}

	let apiObject: HTMLElement
	let parent = element

	while(parent = parent.parentElement) {
		if(parent.dataset.api !== undefined) {
			apiObject = parent
			break
		}
	}

	if(!apiObject) {
		throw "API object not found"
	}

	return apiObject.dataset.api
}