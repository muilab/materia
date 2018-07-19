import Materia from "../Materia"
import findAPIEndpoint from "../Utils/findAPIEndpoint"

// New
export function newObject(cmf: Materia, button: HTMLButtonElement) {
	let dataType = button.dataset.type

	cmf.network.post(`/api/new/${dataType}`)
	.then(response => response.json())
	.then(obj => cmf.app.load(`/${dataType}/${obj.id}/edit`))
	.catch(err => cmf.status.showError(err))
}

// Delete
export function deleteObject(cmf: Materia, button: HTMLButtonElement) {
	let confirmType = button.dataset.confirmType
	let returnPath = button.dataset.returnPath

	if(!confirm(`Are you sure you want to delete this ${confirmType}?`)) {
		return
	}

	let endpoint = findAPIEndpoint(button)

	cmf.network.post(endpoint + "/delete")
	.then(() => cmf.app.load(returnPath))
	.catch(err => cmf.status.showError(err))
}