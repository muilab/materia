import Materia from "../Materia"

// New
export function newObject(cmf: Materia, button: HTMLButtonElement) {
	let dataType = button.dataset.type

	cmf.network.post(`/api/new/${dataType}`)
	.then(response => response.json())
	.then(obj => cmf.app.load(`/${dataType}/${obj.id}/edit`))
	.catch(err => cmf.status.showError(err))
}