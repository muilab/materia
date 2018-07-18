import Materia from "../Materia"
import findAPIEndpoint from "../Utils/findAPIEndpoint"

// Save new data from an input field
export async function save(cmf: Materia, input: HTMLElement) {
	let body = {}
	let isContentEditable = input.isContentEditable
	let value = isContentEditable ? input.textContent : (input as HTMLInputElement).value

	if(value === undefined) {
		return
	}

	// Trim value
	value = value.trim()

	// Add field value to our request body
	if((input as HTMLInputElement).type === "number" || input.dataset.type === "number") {
		if(input.getAttribute("step") === "1" || input.dataset.step === "1") {
			body[input.dataset.field] = parseInt(value)
		} else {
			body[input.dataset.field] = parseFloat(value)
		}
	} else {
		body[input.dataset.field] = value
	}

	// Disable editing on the element
	if(isContentEditable) {
		input.contentEditable = "false"
	} else {
		(input as HTMLInputElement).disabled = true
	}

	// Find API endpoint
	let apiEndpoint = findAPIEndpoint(input)

	try {
		// Apply the change
		await cmf.network.post(apiEndpoint, body)

		// Reload content
		cmf.reloadContent()
	} catch(err) {
		cmf.reloadContent()
		cmf.status.showError(err)
	} finally {
		// Enable editing on the element
		if(isContentEditable) {
			input.contentEditable = "true"
		} else {
			(input as HTMLInputElement).disabled = false
		}
	}
}