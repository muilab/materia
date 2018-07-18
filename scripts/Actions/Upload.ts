import Materia from "../Materia";
import bytesHumanReadable from "../Utils/bytesHumanReadable"
import uploadWithProgress from "../Utils/uploadWithProgress"

// Select file
export function selectFile(cmf: Materia, button: HTMLButtonElement) {
	let fileType = button.dataset.type
	let endpoint = button.dataset.endpoint

	// Click on virtual file input element
	let input = document.createElement("input")
	input.setAttribute("type", "file")

	input.onchange = () => {
		let file = input.files[0]

		if(!file) {
			return
		}

		// Check mime type for images
		if(fileType === "image" && !file.type.startsWith("image/")) {
			cmf.status.showError(file.name + " is not an image file!")
			return
		}

		uploadFile(file, fileType, endpoint, cmf)
	}

	input.click()
}

// Upload file
function uploadFile(file: File, fileType: string, endpoint: string, cmf: Materia) {
	let reader = new FileReader()

	reader.onloadend = async () => {
		let fileSize = reader.result.byteLength

		if(fileSize === 0) {
			cmf.status.showError("File is empty")
			return
		}

		cmf.status.showInfo(`Preparing to upload ${fileType} (${bytesHumanReadable(fileSize)})`, -1)

		try {
			await uploadWithProgress(endpoint, {
				method: "POST",
				credentials: "include",
				headers: {
					"Content-Type": "application/octet-stream"
				},
				body: reader.result
			}, e => {
				let progress = e.loaded / (e.lengthComputable ? e.total : fileSize) * 100
				cmf.status.showInfo(`Uploading ${fileType}...${progress.toFixed(1)}%`, -1)
			})

			cmf.status.showInfo(`Successfully uploaded your new ${fileType}.`)
		} catch(err) {
			cmf.status.showError(`Failed uploading your new ${fileType}.`)
			console.error(err)
		}
	}

	cmf.status.showInfo(`Reading ${fileType} from disk...`, -1)
	reader.readAsArrayBuffer(file)
}
