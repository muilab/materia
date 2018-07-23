import Materia from "../Materia"

const themes = {
	"Default": {},
	"Blue": {
		"link-color-h": "200",
		"link-color-l": "48%"
	},
	"Pink": {
		"link-color-h": "330",
		"link-color-l": "48%"
	},
	"Green": {
		"link-color-h": "110",
		"link-color-l": "34%"
	}
}

// Changes the active theme
export function changeTheme(cmf: Materia, button: HTMLButtonElement) {
	let themeName = button.dataset.theme
	let root = document.documentElement
	let theme = themes[themeName]

	for(let property in theme) {
		if(!theme.hasOwnProperty(property)) {
			continue
		}

		if(themes.Default[property] === undefined) {
			themes.Default[property] = root.style.getPropertyValue(`--${property}`)
		}

		root.style.setProperty(`--${property}`, theme[property])
	}
}