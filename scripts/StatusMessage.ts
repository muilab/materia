import { delay } from "./Utils"

const defaultInfoDelay = 2000
const defaultErrorDelay = 4000

export default class StatusMessage {
	container: HTMLElement
	text: HTMLElement

	constructor(container: HTMLElement, text: HTMLElement) {
		this.container = container
		this.text = text
	}

	show(message: string, duration: number) {
		let messageId = String(Date.now())

		this.text.textContent = message

		this.container.classList.remove("fade-out")
		this.container.dataset.messageId = messageId

		// Negative duration means we're displaying it forever until the user manually closes it
		if(duration === -1) {
			return
		}

		delay(duration).then(() => {
			if(this.container.dataset.messageId !== messageId) {
				return
			}

			this.close()
		})
	}

	clearStyle() {
		this.container.classList.remove("info-message")
		this.container.classList.remove("error-message")
	}

	showError(message: string | Error, duration?: number) {
		this.clearStyle()
		this.show(message.toString(), duration || defaultInfoDelay)
		this.container.classList.add("error-message")
	}

	showInfo(message: string, duration?: number) {
		this.clearStyle()
		this.show(message, duration || defaultErrorDelay)
		this.container.classList.add("info-message")
	}

	close() {
		this.container.classList.add("fade-out")
	}
}