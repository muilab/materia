import Diff from "scripts/Diff"
import { delay } from "scripts/Utils"

const defaultInfoDelay = 2000
const defaultErrorDelay = 4000

export default class StatusBar extends HTMLElement {
	text: HTMLElement

	connectedCallback() {
		this.text = document.getElementById("status-bar-text")
		this.showInfo("Ready!")
	}

	show(message: string, duration: number) {
		let messageId = String(Date.now())

		this.text.textContent = message
		this.classList.remove("fade-out")
		this.dataset.messageId = messageId

		// Negative duration means we're displaying it forever until the user manually closes it
		if(duration === -1) {
			return
		}

		delay(duration).then(() => {
			if(this.dataset.messageId !== messageId) {
				return
			}

			this.close()
		})
	}

	showInfo(message: string, duration?: number) {
		this.clearStyle()
		this.show(message, duration || defaultErrorDelay)
		this.classList.add("info-message")
	}

	showError(message: string | Error, duration?: number) {
		this.clearStyle()
		this.show(message.toString(), duration || defaultInfoDelay)
		this.classList.add("error-message")
	}

	clearStyle() {
		this.classList.remove("info-message")
		this.classList.remove("error-message")
	}

	close() {
		this.classList.add("fade-out")
	}
}