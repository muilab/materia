import Diff from "../Diff"

export default function mountMountables() {
	let mountables = [...document.getElementsByClassName("mountable")]

	let fadeIndex = function(i) {
		return function() {
			Diff.mutations.queue(() => mountables[i].classList.add("mounted"))
		}
	}

	let count = 0

	for(let i = 0; i < mountables.length; i++) {
		if(mountables[i].classList.contains("mounted")) {
			continue
		}

		// Special case: Paragraphs in blockquotes should never be mounted.
		if(mountables[i].parentElement.tagName === "BLOCKQUOTE") {
			mountables[i].classList.remove("mountable")
			continue
		}

		window.setTimeout(fadeIndex(i), count * 50)
		count++
	}
}