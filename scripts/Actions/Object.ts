import Materia from "../Materia"

// New
export function newObject(cmf: Materia, button: HTMLButtonElement) {
	let dataType = button.dataset.type

	// cmf.post(`/api/new/${dataType}`)
	// .then(response => response.json())
	// .then(obj => mat.app.load(`/${dataType}/${obj.id}/edit`))
	// .catch(err => mat.statusMessage.showError(err))
}