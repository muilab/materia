export default function requestIdleCallback(func: Function) {
	if("requestIdleCallback" in window) {
		window["requestIdleCallback"](func)
	} else {
		func()
	}
}