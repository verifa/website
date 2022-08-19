export async function load({ fetch }) {
	const res = await fetch('/sitemap.xml');
	return {};
}
