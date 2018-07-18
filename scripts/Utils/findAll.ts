export default function findAll(className: string) {
	return document.getElementsByClassName(className) as HTMLCollectionOf<HTMLElement>
}