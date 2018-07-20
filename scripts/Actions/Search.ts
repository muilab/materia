import Materia from "../Materia"
import Diff from "../Diff";

export async function searchMaterials(cmf: Materia, input: HTMLInputElement) {
	if(input.value === "") {
		return
	}

	let response = await fetch("/_/library/search/" + input.value)
	let html = await response.text()
	Diff.innerHTML(cmf.app.content, html)
}