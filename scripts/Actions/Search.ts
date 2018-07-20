import Materia from "../Materia"
import Diff from "../Diff";

export async function searchMaterials(cmf: Materia, input: HTMLInputElement) {
	let url = "/_/library/search/" + input.value

	if(input.value === "") {
		url = "/_/"
	}

	let response = await fetch(url)
	let html = await response.text()
	await Diff.innerHTML(cmf.app.content, html)
	cmf.app.emit("DOMContentLoaded")
}