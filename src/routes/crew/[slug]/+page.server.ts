import { crewByID } from "$lib/crew/crew";
import { blogTypes, getPostsGlob } from "$lib/posts/posts";
import { error } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
	const member = crewByID(params.slug);
	if (!member) {
		error(404, "could not find crew with ID " + params.slug);
	}
	return {
		posts: getPostsGlob({
			types: blogTypes,
			author: member.id
		}),
		member: member,
	}
}